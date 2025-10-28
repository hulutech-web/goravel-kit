<template>
  <div class="mt-3">

    <a-row :gutter="[32,2]">
      <a-col :span="4">
        <a-divider>文件夹</a-divider>
        <button v-for="(f,index) in myFolders" class="mb-3 cursor-pointer w-full py-2 text-start"
                :class="selected_id==f.id?'active':''"
                style="border:none;display:block;" :key="index" @click="loadFolderData(f.id)">
          <folder-filled style="color:#d48a1b"/>
          <span class="my-2">
              {{ f.name }}
            </span>
          <span class="text-sm text-blue-500">{{ f.count }}</span>
        </button>

      </a-col>
      <a-col :span="20">
        <a-space>
          <a-select v-model:value="fold_id" style="width:200px;">
            <a-select-option value="">请选择</a-select-option>
            <a-select-option v-for="(f,index) in myFolders"
                             :key="index" :value="f.id">
              {{ f.name }}
            </a-select-option>
          </a-select>
          <a-input placeholder="请输入内容" v-model:value="keyword" style="width:400px;">
          </a-input>
          <a-button type="primary" danger @click="searchBy">搜索</a-button>
          <a-button type="primary" plain @click="keyword=''">清空</a-button>
          <FileUpload v-model="uploadVisible" @success="fetchData" v-model:cate_id="cate_id" :folders="myFolders">
            <a-button @click="uploadVisible=true">上传文件</a-button>
          </FileUpload>
          <a-button type="primary" v-show="del_obj.del_ids.length>0" danger plain @click="delBatch">删除</a-button>
        </a-space>
        <div class="border border-solid border-stone-200">
          <a-spin tip="Loading..." :spinning="loading">
            <a-checkbox-group v-model:value="del_obj.del_ids" class="min-h-[700px] mt-2" style="width: 100%">
              <a-row :gutter="[32,64]">
                <a-col :span="3" v-for="(f,ind) in files" :key="ind">
                  <div class="w-[140px] h-[140px] text-center" style="border: .5px dotted #e5e7eb;">
                    <a-image
                        v-if="f.type==='image'"
                        ref="imageRef"
                        :height="140"
                        style="height:140px;object-fit: cover;"
                        :src="`${baseApiUrl}/uploads/${f.path}`"
                        :preview-src-list="[`${baseApiUrl}/uploads/${f.path}`]"
                    />
                    <file-filled @click="showFile(f.path)" :height="'160px'" :width="'160px'"
                                 style="width: 160px; height: 160px;font-size:160px;color:#e6a23c" v-else/>
                    <div class="text-center font-light text-xs text-blue-400">
                      <div class="flex">
                        <a-input
                            class="flex-1"
                            ref="inputRef"
                            v-model:value="f.name"
                            type="text"
                            size="small"
                            :style="{ width: '100%',color:'grey' }"
                        />
                        <a-button size="small" type="link" @click="changeName(f)">改名</a-button>
                      </div>
                      <div>
                        <label class="text-lg cursor-pointer" :for="`f-${f.id}`">选择</label>
                        <a-checkbox :id="`f-${f.id}`" :value="f.id"></a-checkbox>
                      </div>
                    </div>
                  </div>
                </a-col>
              </a-row>
            </a-checkbox-group>
          </a-spin>
        </div>
        <a-pagination
            class="mt-3"
            v-model:current="pageForm.page"
            v-model:page-size="pageForm.page_size"
            :page-sizes="[24,48,96]"
            :background="true"
            layout="total, sizes, prev, pager, next, jumper"
            :total="pageForm.total"
            @change="handleSizeChange"
            :show-total="total => `共 ${pageForm.total} 条`"
        />
      </a-col>

    </a-row>
  </div>
</template>

<script setup>
import {ref} from "vue";
import {Modal} from 'ant-design-vue';
import Api from "@/api";
import {request} from "@/utils/request";
import FileUpload from "@/components/file-manager/FileUpload.vue";
import XEUtils from "xe-utils";


const baseApiUrl = import.meta.env.VITE_API_URL ? import.meta.env.VITE_API_URL : '';

const props = defineProps({
  folders: Array,
});
const pageForm = ref({
  page: 1,
  total: 0,
  page_size: 24,
})
const files = ref([]);
const myFolders = ref([])
watch(() => props.folders, (val) => {
  myFolders.value = val
}, {immediate: true, deep: true})

const fold_id = ref("")
const keyword = ref("")
const searchBy = () => {
  pageForm.value.page = 1;
  loadFolderData()
}


const handleSizeChange = async (e) => {
  console.log(e)
  let response = await Api.fileController.index({
    cid: selected_id.value ? selected_id.value : fold_id.value,
    keyword: keyword.value,
    currentPage: e,
    pageSize: 24
  })
  files.value = response.data
  pageForm.value.total = response.total
  pageForm.value.page = response.meta?.current_page
};


// 上传文件显示
const uploadVisible = ref(false)
//分类id
const cate_id = ref(null)

watch(() => uploadVisible, (val) => {
  if (!val) {
    //继续向上传递
    console.log("==================", val)
  }
})

const emit = defineEmits(['open-folder', 'delete-folder', 'download-file', 'delete-file']);
const del_obj = ref({
  del_ids: []
})
const loading = ref(false);
const selected_id = ref(null)
const loadFolderData = async (id) => {
  if (id) {
    pageForm.value.page = 1;
  }
  console.log("loadFolderData", keyword.value, fold_id.value)

  loading.value = true
  selected_id.value = id
  console.log(keyword.value)
  let response = await request({
    url: `/api/admin/file?pageSize=24&currentPage=${pageForm.value.page
    }&${XEUtils.serialize({name: keyword.value, cid: selected_id.value ? selected_id.value : fold_id.value})}`,
    method: "GET",
  });
  files.value = response.data
  pageForm.value.total = response.total
  pageForm.value.page = response.meta?.current_page
  loading.value = false
}

const changeName = async (f) => {
  await Api.fileController.update({id: f.id}, f)
}

const handleOpenFolder = (id) => {
  fold_id.value = id
  loadFolderData(id)
};
const fetchData = (val) => {
  console.log(val)
}
const handleDeleteFolder = (id) => {
  emit('delete-folder', id);
};

const delBatch = async () => {
  Modal.confirm({
    content: '删除所有',
    async onOk() {
      await Api.fileController.delBatch(del_obj.value)
    },
    cancelText: 'Click to destroy all',
    onCancel() {
      Modal.destroyAll();
    },
  });
}
const showFile = async (path) => {
  let url = `${baseApiUrl}/uploads/${path}`
  let content = ``;
  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.text(); // 或 response.json()
    console.log(data)
    content = data;
  } catch (error) {
    console.error('Fetch error:', error);
  }

  Modal.info({
    title: '查看文件',
    content: content,
    okText: '关闭',
    cancelText: '取消',
    width: '60%',
    style: {
      top: '20px',
    },
  })
  // window.open(`http://localhost:3000/uploads/${path}`)
}

defineExpose({
  searchBy
})
</script>
<style>
.a-dropdown-link {
  cursor: pointer;
  color: #409EFF;
}

.a-icon-arrow-down {
  font-size: 12px;
}

.active {
  background: #409EFF;
  color: white;
}
</style>
