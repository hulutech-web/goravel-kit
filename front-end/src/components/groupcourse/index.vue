<template>
  <div>
    <div>
      <a-button type="primary" @click="openModal = true">
        <template #icon>
          <ImportOutlined/>
        </template>
      </a-button>
      <span class="font-bold text-gray-500 text-sm">
        编号：【{{ props.course_id }}】
        课名：【{{ props.course_name }}】
        场馆：【{{ props.venue_name }}】
      </span>
    </div>
    <div class="mt-1">
      <a-modal v-model:visible="openModal" title="选择课程" centered width="800px">
        <a-input-search
            v-model:value="title"
            enter-button
            @search="fillData"
            class="mb-2"
            placeholder="请输入名字"
        />
        <a-table :dataSource="dataSource" :columns="columns" bordered size="small">
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'id'">
              {{record.id}}
            </template>
            <template v-if="column.key === 'title'">
              {{record.title}}
            </template>
            <template v-if="column.key === 'action'">
              <a-button type="primary" @click="selectOne(record)">
                <CheckOutlined/>
                选择
              </a-button>
            </template>
            <template v-if="column.key === 'venue_name'">
              {{record.venue.name}}
            </template>
          </template>
        </a-table>
      </a-modal>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Api from "@/api"

const dataSource = ref([]);
const props = defineProps({
  course_id: {
    type: Number,
    default: 0,
  },
  course_name: {
    type: String,
    default: "",
  },
  venue_name:{
    type:String,
    default:""
  }
});
const columns = [
  {
    dataIndex: "id",
    title: "课程ID",
    key: "id",
  },
  {
    dataIndex: "title",
    title: "课程名称",
    key: "title",
  },
  {
    dataIndex: "venue_name",
    title: "场馆",
    key: "venue_name",
  },
  {
    dataIndex: "action",
    title: "操作",
    key: "action",
  },
];
const title = ref("");
onMounted(() => {
  fillData();
});
const fillData = async () => {
  const {data} = await Api.courseController.index({title: title.value});
  if(data&&data.length>0){
    dataSource.value = data.map(item=>{
      return {
        ...item,
        venue_name:item.venue.name,
      }
    })
  }
  console.log("dataSource",dataSource.value)
};
const openModal = ref(false);
const emit = defineEmits(["update:course_id", "update:course_name","update:venue_name"]);
const selectOne = (row) => {
  //提交数据到父组件
  emit("update:course_id", row.id);
  emit("update:course_name", row.title);
  emit("update:venue_name", row.venue.name);
  openModal.value = false;
};
</script>

<style scoped></style>
