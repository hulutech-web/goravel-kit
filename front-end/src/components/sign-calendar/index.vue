<template>
  <div>
    <!-- 团课课程预约 -->
    <FullCalendar :options="calendarOptions" ref="fullCalendarRef" />
  </div>
</template>

<script setup lang="ts">
import dayjs from "dayjs";
import FullCalendar from "@fullcalendar/vue3";
import useOpts from "./options";
import Api from "@/api"
import { toReactive } from "@vueuse/core";
const instance = getCurrentInstance();
const fullCalendarRef = ref();
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({}),
  }
});

const emit = defineEmits(["update:modelValue"]);

interface TagEvent {
  id: string,
  num: number,
  title: string,
  groupcourse_id: number
  groupcourse_name: string
  start: string,
  end: string
  coverimage: string
}
const tagArr = ref<TagEvent[]>([])



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
const heartbeatStyle = {
  animation: 'heartbeat .5s infinite alternate'
};
let allEventElements = [];
// #region 功能函数
const evClickCbk = (info) => {
  const eventEl = info.el;
  const events = getEvents();
  allEventElements.map(item=>{
    item.style.removeProperty('animation');
  })
  Object.assign(eventEl.style, heartbeatStyle);
  //提交数据到父组件
  emit("update:modelValue", info.event);
};

const dtClickCbk = (info) => {
  console.log("子组件：空白处被点击", info);
};
const mountedCbk = (info) => {
  allEventElements.push(info.el)
  const eventEl = info.el;
  // console.log("子组件：任意时间触发时", info)
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
  console.log("子组件：选择事件选择", event);
};
const evtAddCbk = (info) => {
  // console.log("子组件：添加事件被触发", info);
};
const evtCbk = (info) => {
  console.log("子组件：事件被触发", info);
};
const evtContentCbk = (arg, createElement, containerEl) => {
  console.log("子组件：自定义样式");
};
const evtMouseEnterCbk = (info) => {

};
const evtMouseLeaveCbk = (info) => {

};
const eventDropCbk = (info) => {
  console.log("子组件：拖拽事件时");
};
const user_id=ref()
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
    eventDropCbk,
    user_id.value,
);
//获取所有事件
const getEvents = () => {
  return fullCalendarRef.value.getApi().view.calendar.getEvents();
};
// 通过ID获取事件
const getEventById = (id) => {
  return fullCalendarRef.value.getApi().view.calendar.getEventById(id);
};
// 添加事件
const addEvent = (event) => {
  console.log(2, "添加事件到页面--未提交后台", event);
  const calendarApi = fullCalendarRef.value.getApi();
  calendarApi.addEvent(event);
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
});
// #endregion
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
</style>