<template>
  <div>
      <a-space>
        <a-button-group>
            <a-button type="primary" @click="handleCreateFolder" v-auth:netdisk>创建文件夹</a-button>
        </a-button-group>
      </a-space>

      <a-row>
        <a-col :span="24">
          <FileList :folders="folders"  @refresh="fetchData" ref="fileListRef" />
        </a-col>
      </a-row>
      <FoldCreate v-model="folderCreateVisible" @success="fetchData" @resetVisible="resetVisible" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import FileList from '@/components/file-manager/FileList.vue';
import FoldCreate from '@/components/file-manager/FoldCreate.vue';
import Api from "@/api"
const folders = ref([]);
const getFolders = async ()=>{
  folders.value = await Api.fileCateController.list()
}
const fileListRef=ref(null)
const uploadVisible = ref(false);
const folderCreateVisible = ref(false);

// 获取文件和文件夹数据
const fetchData = async () => {
  await getFolders();
  fileListRef.value.searchBy();
};



// 打开创建文件夹弹窗
const handleCreateFolder = () => {
  folderCreateVisible.value = true;
};

const resetVisible = () => {
  uploadVisible.value = false;
  folderCreateVisible.value = false;
};

// 初始化时加载数据
onMounted(() => {
  fetchData();
});

const open = ref<boolean>(false);

const showModal = () => {
  open.value = true;
};

const handleOk = (e: MouseEvent) => {
  console.log(e);
  open.value = false;
};
</script>
