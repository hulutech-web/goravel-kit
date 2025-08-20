<template>
  <div>
    <slot name="default">
      <a-space>
        <a-button @click="showModal" size="small">
          <cloud-upload-outlined/>
          选择
        </a-button>
        <span class="text-sm text-sky-500" style="word-break: break-all;">{{ file.name }}</span>
      </a-space>
    </slot>
    <a-row>
      <a-col :span="24">
        <a-modal v-model:visible="visible" title="上传文件" @ok="getFile" @cancel="handleClose"
                 style="width:1200px;top: 20px"
                 :maskClosable="false" :center="true">
          <FList ref="fileRef" :folders="folders" @refresh="fetchData"/>
        </a-modal>
      </a-col>
    </a-row>
  </div>
</template>

<script setup lang="ts">
import {ref} from 'vue';
import FList from '@/components/file-select/FList.vue';
import Api from "@/api"


const props = defineProps({
  modelValue: String,
})
const uri = ref(props.modelValue)
const file = ref({
  id: null,
  uri: "",
  name: ""
})
const folders = ref([]);
const visible = ref(false);
const handleClose = () => {
}
const getFolders = async () => {
  folders.value = await Api.fileCateController.list()
}
const fileRef = ref(null)
const uploadVisible = ref(false);
const folderCreateVisible = ref(false);

// 获取文件和文件夹数据
const fetchData = async () => {
  await getFolders();
};
const emit = defineEmits(['update:modelValue'])
const getFile = () => {
  file.value = fileRef.value.getFile()
  uri.value = file.value.uri
  emit('update:modelValue', file.value.uri)
  visible.value = false
  return file.value
}


watch(visible, (val) => {
  if (val) {
    fetchData()
  }
})

const open = ref<boolean>(false);

const showModal = () => {
  visible.value = true;
  nextTick(() => {
    fileRef.value.searchBy();
  })
};
defineExpose({
  showModal,
});
const handleOk = (e: MouseEvent) => {
  console.log(e);
  open.value = false;
};
</script>
