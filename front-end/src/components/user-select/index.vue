<template>
  <div>
    <div>
      <a-button type="primary" shape="round" @click="openModal = true">
        <template #icon>
          <UserOutlined/>
        </template>
      </a-button>
      <span class="font-bold text-gray-500 text-small">
        【会员编号：{{ props.id }}
        姓名：{{ props.username }}】
      </span>
    </div>
    <div class="mt-1">
      <a-modal v-model:visible="openModal" width="600px" title="选择会员" centered>
        <a-input-search v-model:value="username" enter-button @search="fillData" placeholder="请输入名字或手机号"/>
        <div class="mt-3">
          <a-table :dataSource="dataSource" :columns="columns" size="small" bordered>
            <template #bodyCell="{ column, record }">
              <template v-if="column.key == 'action'">
                <a-button type="primary" @click="selectOne(record)">
                  <CheckOutlined/>
                  选择
                </a-button>
              </template>
            </template>
          </a-table>
        </div>
      </a-modal>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Api from "@/api"

const dataSource = ref([]);
const props = defineProps({
  id: {
    type: Number,
    default: 0,
  },
  username: {
    type: String,
    default: "",
  },
});

const columns = [
  {
    dataIndex: "id",
    title: "会员ID",
    key: "id",
  },
  {
    dataIndex: "username",
    title: "会员姓名",
    key: "username",
  },
  {
    dataIndex: "phone",
    title: "电话",
    key: "phone",
  },
  {
    dataIndex: "action",
    title: "操作",
    key: "action",
    width: 100,
  },
];

const username = ref("");


const fillData = async () => {
  const {data} = await Api.userController.index({username: username.value});
  dataSource.value = data;
};
const openModal = ref(false);

/*监听openModal的执行，当为true时才开启网络加载数据*/
watch(openModal,(val)=>{
  if(val){
    fillData()
  }
})
const emit = defineEmits(["update:id", "update:username"]);
const selectOne = (row) => {
  console.log(row)
  //提交数据到父组件
  emit("update:id", row.id);
  emit("update:username", row.username);
  openModal.value = false;
};
</script>

<style scoped></style>