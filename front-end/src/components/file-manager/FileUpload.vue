<template>
  <div>
    <slot name="default">

    </slot>
    <a-modal v-model:visible="innerVisible" title="上传文件" @cancel="handleClose" style="width:500px;"
             :maskClosable="false" :center="true" :footer="false">
      <a-form :model="uploadData"></a-form>
      <a-form-item label="文件夹">
        <a-select @change="changeCate" v-model:value="uploadData.cate_id">
          <a-select-option value="">请选择分类</a-select-option>
          <a-select-option v-for="(c,ind) in myFolders" :key="ind" :value="c.id">{{ c.name }}</a-select-option>
        </a-select>
      </a-form-item>
      <a-form-item label="文件">

        <a-upload-dragger
            v-model:fileList="fileList"
            name="file"
            :multiple="true"
            :headers="headers"
            @change="handleChange"
            :before-upload="beforeUpload"
        >
          <p class="ant-upload-drag-icon">
            <inbox-outlined></inbox-outlined>
          </p>
          <p class="ant-upload-text">点击或者拖拽上传内容</p>
          <p class="ant-upload-hint">
            支持多种格式，单文件多文件都支持
          </p>

        </a-upload-dragger>
      </a-form-item>
      <a-form-item>
        <a-button
            type="primary"
            :disabled="fileList.length === 0"
            :loading="uploading"
            style="margin-top: 16px"
            @click="handleUpload"
        >
          {{ uploading ? '上传中...' : '开始上传' }}
        </a-button>
      </a-form-item>
    </a-modal>
  </div>

</template>

<script setup>
import {ref} from 'vue';
import {request} from "@/utils/request";

import {InboxOutlined} from '@ant-design/icons-vue';
import {message} from "ant-design-vue";

const fileCateID = ref(null)
const handleChange = (info) => {
  const status = info.file.status;
  if (status !== 'uploading') {
    // console.log(info.file, info.fileList);
  }
  if (status === 'done') {
    message.success(`${info.file.name} file uploaded successfully.`);
  } else if (status === 'error') {
    message.error(`上传失败：${info.file.response.data}`);
  }
};
const uploadData = ref({
  cate_id: null,
  files: []
});
const headers = {
  Authorization: 'Bearer ' + localStorage.getItem("token")
}

const props = defineProps({
  modelValue: Boolean,
  cate_id: Number,
  folders: Array,
});

const fileList = ref([]);

const myFolders = ref([]);

const uploadRef = ref();
const emit = defineEmits(['update:modelValue', 'success']);
const innerVisible = ref(props.modelValue);

watch(() => props.modelValue, (val) => {
  innerVisible.value = val;
})
const my_id = ref(null)
watch(() => props.cate_id, (val) => {
  my_id.value = val;
})
watch(() => props.folders, (val) => {
  myFolders.value = val
})
watchEffect((fileList)=>{
  uploadData.value.files = fileList
})
const submitUpload = () => {
  if(uploadData.value.cate_id==null){
    message.error("请选择上传的文件夹")
    return;
  }
  if(fileList.value.length===0){
    message.error("请选择上传的文件")
    return;
  }
  console.log(fileList.value)
  if (my_id.value) {

  } else {
    message.error("请选择上传的文件夹")
  }

}


const beforeUpload = file => {
  fileList.value = [...(fileList.value || []), file];
  return false;
};

const changeCate = (e) => {
  my_id.value = e
}
const handleClose = () => {
  emit('update:modelValue', false);
};

const uploading = ref(false);
const handleUpload = async () => {
  if (!my_id.value) {
    message.error("请选择上传的文件夹");
    return;
  }

  if (fileList.value.length === 0) {
    message.error("请选择上传的文件");
    return;
  }

  uploading.value = true;

  try {
    // Create an array of upload promises
    const uploadPromises = fileList.value.map(fileItem => {
      const formData = new FormData();
      formData.append('file', fileItem.originFileObj); // Use originFileObj for the actual file
      formData.append('cate_id', my_id.value); // Add category ID to form data

      return request(`/api/admin/upload`, {
        method: 'post',
        data: formData,
        headers: {
          'Content-Type': 'multipart/form-data',
          'Authorization': 'Bearer ' + localStorage.getItem("token")
        }
      });
    });

    // Execute all uploads in parallel
    const results = await Promise.all(uploadPromises);
  } catch (error) {
    console.error('上传错误:', error);
    message.error('上传过程中发生错误');
  } finally {
    uploading.value = false;
  }
};
</script>
