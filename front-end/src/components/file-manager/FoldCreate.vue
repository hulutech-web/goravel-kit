<template>
  <div>
    <a-modal v-model:visible="innerOpen" title="创建文件夹" @close="handleClose" :maskClosable="false">
      <a-form :model="form" laba-width="100px" ref="formRef">
        <a-form-item label="文件夹名称" name="name" has-feedback  :rules="[{ required: true, message: '请输入名称' }]">
          <a-input v-model:value="form.name"></a-input>
        </a-form-item>
        <a-form-item label="文件类型" name="type" has-feedback  :rules="[{ required: true, message: '请选择类型' }]">
          <a-select v-model:value="form.type">
            <a-select-option value="">请选择</a-select-option>
            <a-select-option value="image">图片</a-select-option>
            <a-select-option value="video">视频</a-select-option>
            <a-select-option value="audio">音频</a-select-option>
            <a-select-option value="file">文件</a-select-option>
          </a-select>
        </a-form-item>
      </a-form>
      <template #footer>
        <a-button @click="handleClose">取消</a-button>
        <a-button type="primary" @click="handleCreate">创建</a-button>
      </template>
    </a-modal>
  </div>
</template>

<script setup>
import {ref} from 'vue';
import Api from "@/api"
import {message as $message} from "ant-design-vue";
const formRef=ref()
const createFolder = async () => {
  await Api.fileCateController.store(form.value)
}

const cates = ref([])
const loadCate = async () => {
  const {data} = await Api.fileCateController.list()
  cates.value = data;
}
const props = defineProps({
  modelValue: Boolean
});
const innerOpen = ref(props.modelValue);

watch(() => props.modelValue, (val) => {
  innerOpen.value = val;
  console.log("watch", innerOpen.value)
}, {immediate: true})
const emit = defineEmits(['update:modelValue', 'success']);

const form = ref({
  pid: null,
  name: '',
  type: ''
});

const handleCreate = async () => {

  formRef.value?.validateFields().then(async (res) => {
        console.log("res",res)
        await createFolder(form.value);
        $message.error({content: "成功", key: "success"});
        emit('success');
        emit('update:modelValue', false);
      })
      .catch((e) => {
        console.error(e);
      })
      .finally(() => {
      });

};

const handleClose = () => {
  emit('update:modelValue', false);
};
</script>
