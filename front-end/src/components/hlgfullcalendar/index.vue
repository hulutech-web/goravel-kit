<template>
  <div>
    <FullCalendar :options="calendarOptions" ref="fullCalendarRef" />
    <!-- 这里是修改过程 -->
    <a-drawer
        :zIndex="100"
        width="800"
        placement="right"
        v-model:visible="openDrawer"
        @close="close"
        title="课程排期"
    >
      <a-form ref="formRef" :model="scheduleState">
        <a-form-item
            label="开始时间"
            name="start_time"
            :rules="[{ required: true,message:'请选择时间' }]"
        >
          <a-input v-model:value="scheduleState.start_time" disabled></a-input>
        </a-form-item>
        <a-form-item
            label="结束时间"
            name="end_time"
            :rules="[{ required: true }]"
        >
          <a-input v-model:value="scheduleState.end_time" disabled></a-input>
        </a-form-item>
        <a-form-item
            label="排期颜色"
            name="color"
            :rules="[{ required: true ,message:'请选择颜色'}]"
        >
          <picker v-model="scheduleState.color"/>
        </a-form-item>
        <a-form-item
            label="选择课程"
            name="course_id"
            :rules="[{ required: true,message:'请选择课程' }]"
        >
          <Groupcourse
              v-model:course_id="scheduleState.course_id"
              v-model:course_name="scheduleState.course_name"
              v-model:venue_name="scheduleState.venue_name"
          />
        </a-form-item>
        <a-form-item
            label="选择教练"
            name="coach_id"
            :rules="[{ required: true ,message:'请选择教练'}]"
        >
          <EmployeeSel v-model="scheduleState.coach_id"/>
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button type="primary" @click="onUpdate">提交</a-button>
            <slot :currentEvent="currentEvent">
              <a-button type="primary">关闭</a-button>
            </slot>
          </a-space>
        </a-form-item>
      </a-form>
    </a-drawer>
  </div>
</template>

<script setup lang="ts">
import dayjs from "dayjs";
import Groupcourse from "@/components/groupcourse/index.vue";
import FullCalendar from "@fullcalendar/vue3";
import useOpts from "./options";
// import useRulesStore from "@/store/useRulesStore.ts";
import { toReactive } from "@vueuse/core";
import Api from "@/api"
import EmployeeSel from "@/components/employee-sel/index.vue";

// const rulesStore = useRulesStore();
const instance = getCurrentInstance();
const { proxy } = instance;
const fullCalendarRef = ref();
const activeKey = ref(["1"]);
const props = defineProps({
  getSelData: {
    type: Function,
    default: () => {},
  },
  getDropData: {
    type: Function,
    default: () => {},
  },
});
// 国际化用
// UTC格式化成本地时间
const formatLocalTime = (
    timeStr: string,
    formatter = "YYYY-MM-DD HH:mm:ss"
) => {
  if (!(timeStr && dayjs(timeStr))) return "";
  // 使用 Day.js 的 parse 函数将 FullCalendar 时间解析为 Day.js 对象
  const parsedTime = dayjs(timeStr);
  //显示中国时间减去8小时
  const localTime = parsedTime.subtract(8, "hours");

  // 将 Day.js 对象格式化为所需的格式，例如 'YYYY-MM-DD HH:mm:ss'
  const formattedTime = localTime.format("YYYY-MM-DD HH:mm:ss");

  return formattedTime;
};

//当前需要显示的事件
const currentEvent = ref({
  id: null,
  title: "",
  coach:{},
  venue:{},
  venue_id:null,
  description: "",
  backgroundColor: "",
  schedule: "",
  sign: "",
  users: [],
  start: "",
  end: "",
});
const openDrawer = ref(false);
// #region 功能函数
const evClickCbk = (info) => {
  // console.log("evClickCbk",info.event)
  if (!info.event.id) {
    console.log("evClickCbk==>临时事件", info);
    return;
  }
  openDrawer.value = true;
  currentEvent.value = info.event;
};
const dtClickCbk = (info) => {
  // console.log("子组件：空白处被点击", info);
};
const mountedCbk = (info) => {
  const eventEl = info.el;
  eventEl.style.maxWidth="300px";
  // 添加平滑的悬停动画效果
  eventEl.style.transition = "all 0.3s ease-out"; // 修正拼写错误并优化时间

  // 初始状态
  eventEl.style.transform = "scale(1)";
  eventEl.style.boxShadow = "0 2px 5px rgba(0, 0, 0, 0.1)";
  eventEl.style.margin="10px"

  // 鼠标悬停效果
  eventEl.addEventListener('mouseenter', () => {
    eventEl.style.boxShadow = "0 5px 15px rgba(0, 0, 0, 0.5)";
    eventEl.style.transform = "scale(0.98)";
    eventEl.style.zIndex = "100"; // 确保悬停元素在上层
  });

  // 鼠标离开效果
  eventEl.addEventListener('mouseleave', () => {
    eventEl.style.boxShadow = "0 2px 5px rgba(0, 0, 0, 0.1)";
    eventEl.style.transform = "scale(1)";
    eventEl.style.zIndex = "";
  });

};
const selCbk = (event) => {
  // console.log("子组件：选择事件选择", event);
  let tmpEvent = {
    start: formatLocalTime(event.start),
    end: formatLocalTime(event.end),
    title: event.title,
  };
  addEvent(tmpEvent);
  props.getSelData(event);
};
const evtAddCbk = (info) => {
  // console.log("子组件：添加事件被触发", info);
};
const evtCbk = async (info) => {
  console.log("子组件：事件被触发,startStr,endStr", info.event.startStr,info.event.endStr);
  await Api.scheduleController.change({id:info.event.id,start_time:info.event.startStr,end_time:info.event.endStr})
};
const evtContentCbk = (arg, createElement, containerEl) => {
  // console.log("子组件：自定义样式");
};
const evtMouseEnterCbk = (info) => {
  const eventEl = info.el;
};
const evtMouseLeaveCbk = (info) => {
  const eventEl = info.el;
};
const eventDropCbk = (info) => {
  console.log("子组件：拖拽事件时");
  let event = info.event;
  let newInfo = {
    id: event.id,
    start: event.start,
    end: event.end,
  };
  props.getDropData(newInfo);
};
const calendarOptions = useOpts(
    evClickCbk,
    dtClickCbk,
    mountedCbk,
    selCbk,
    evtAddCbk,
    evtCbk,
    evtContentCbk,
    evtMouseEnterCbk,
    evtMouseLeaveCbk,
    eventDropCbk
);
//获取所有事件
const getEvents = () => {
  return proxy.$refs["fullCalendar"].getApi().view.calendar.getEvents();
};
// 通过ID获取事件
const getEventById = (id) => {
  return proxy.$refs["fullCalendar"].getApi().view.calendar.getEventById(id);
};
// 添加事件
const addEvent = (event) => {
  console.log(2, "添加事件到页面--未提交后台", event);
  // const calendarApi = fullCalendarRef.value.getApi();
  // calendarApi.addEvent(event);
  // // 检查是否已存在相同ID的事件
  // const existingEvent = calendarApi.getEventById(event.id);
  // if (!existingEvent) {
  //   calendarApi.addEvent({
  //     ...event,
  //     id: event.id || `temp_${Date.now()}` // 确保有唯一ID
  //   });
  // } else {
  //   console.warn("Event already exists:", event.id);
  // }
};

// 通过id删除事件
const removeEventById = (event) => {
  let calendarApi = proxy.$refs["fullCalendar"].getApi();
  let calendarFunc = calendarApi.view.calendar;
  let getEvents = calendarFunc.getEvents(); //获取数据
  if (getEvents && getEvents.length > 0) {
    //循环删除数据（通过ID）
    getEvents.map((item) => {
      if (item.id === event.id) {
        calendarFunc.getEventById(item.id).remove();
      }
    });
  }
};

// 删除所有事件
const removeAllEvents = () => {
  let calendarApi = proxy.$refs["fullCalendar"].getApi();
  let calendarFunc = calendarApi.view.calendar;
  let getEvents = calendarFunc.getEvents(); //获取数据
  getEvents.map((item) => {
    item.remove();
  });
};

const refresh = () => {
  console.log("refresh");
  const calendarApi = fullCalendarRef.value.getApi();
  calendarApi.removeAllEvents();
  calendarApi.refetchEvents();
};

// 对外暴露方法refresh
defineExpose({
  refresh,
  openDrawer,
  eventDropCbk,
});
// #endregion
// ==折叠修改
let scheduleState = ref({
  id: null,
  start_time: "",
  end_time: "",
  color: "",
  course_id:null,
  user_id: null,
  course_name:"",
  venue:{},
  venue_name:""
});

watch(openDrawer, async (newVal) => {
  if (newVal) {
    //关闭时，清空折叠面板中的数据
   let data = await Api.scheduleController.show({id:currentEvent.value.id});
    scheduleState.value={
      ...data,
      course_id:data.course_id,
      course_name:data.course.title,
      venue_name:data.course.venue.name,
    }
    // console.log("scheduleState",scheduleState.value)
  } else {
    scheduleState.value = ref({
      id: null,
      start_time: "",
      end_time: "",
      color: "",
      course_id:null,
      user_id: null,
      course_name:"",
      venue:{},
      venue_name:""
    });
  }
});
const onUpdate = async () => {
  await Api.scheduleController.update({id:scheduleState.value.id},scheduleState.value);
  openDrawer.value = false;
  refresh();
};

const close = () => {
  //重置数据scheduleState
  scheduleState.value = ref({
    id: null,
    start_time: "",
    end_time: "",
    color: "",
    course_id:null,
    user_id: null,
    course_name:"",
    venue:{},
    venue_name:""
  });
  refresh();
};

</script>

<style>
@keyframes heartbeat {
  0% {
    /*心跳放大*/
    opacity: 0.9;
    transform: rotate(0.2deg);
  }

  100% {
    /*心跳缩小*/
    opacity: 1;
    transform: rotate(-0.2deg);
  }
}
.fc-event {
  margin-bottom: 2px !important; /* 添加事件间的小间隔 */
  box-sizing: border-box !important;
}

.fc-event-main {
  overflow: hidden !important;
}
</style>
