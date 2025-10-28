<template>
  <div>
    <a-row>
      <a-col :span="6">
        <a-table size="small" :data-source="myFolders" :columns="columns" style="width: 100%" bordered>
          <template #bodyCell="{ column, record }">
            <template v-if="column.dataIndex === 'name'">
              <a-button type="link" @click="loadFolderData(record.id)">{{ record.name }}</a-button>
            </template>
          </template>
        </a-table>
      </a-col>
      <a-col :span="18">
        <a-space>
          <a-input placeholder="请输入内容" v-model="keyword" style="width:400px;">
          </a-input>
          <a-button type="primary" danger @click="searchBy">搜索</a-button>
          <a-button type="primary" plain @click="keyword=''">清空</a-button>
        </a-space>
        <div class="border border-solid border-stone-200">
          <a-radio-group v-model:value="file" class="min-h-[500px] mt-2" style="width: 100%">
            <a-row :gutter="[32,64]">
              <a-col :span="3" v-for="(f,ind) in files" :key="ind">
                <div class="w-[80px] h-[80px] text-center" style="border: .5px dotted #e5e7eb;">
                  <a-image
                      v-if="f.type==='image'"
                      ref="imageRef"
                      :height="80"
                      style="height:80px;object-fit: cover;"
                      :src="`${baseApiUrl}/uploads/${f.path}`"
                      :preview-src-list="[`${baseApiUrl}/uploads/${f.path}`]"
                  />
                  <file-filled @click="showFile(f.path)" :height="'80px'" :width="'80px'"
                               style="width: 80px; height: 80px;font-size:80px;color:#e6a23c" v-else/>
                  <div class="text-center font-light text-xs text-blue-400">
                    <div class="flex">
                      <a-input
                          class="flex-1"
                          ref="inputRef"
                          v-model:value="f.name"
                          type="text"
                          size="small"
                          disabled
                          :style="{ width: '100%',color:'grey' }"
                      />
                    </div>
                    <div>
                      <label class="text-lg cursor-pointer" :for="`f-${f.id}`">选择</label>
                      <a-radio :id="`f-${f.id}`" :value="f"></a-radio>
                    </div>
                  </div>
                </div>
              </a-col>
            </a-row>
          </a-radio-group>

        </div>
        <a-pagination
            class="mt-3"
            v-model:current-page="pageForm.page"
            v-model:page-size="pageForm.page_size"
            :page-sizes="[24,48,96]"
            :background="true"
            layout="total, sizes, prev, pager, next, jumper"
            :total="pageForm.total"
            @change="handleSizeChange"
        />
      </a-col>

    </a-row>
  </div>
</template>

<script setup>
import {ref} from "vue";
import {Modal} from 'ant-design-vue';
import Api from "@/api"

const columns = [
  {
    title: '名称',
    dataIndex: 'name',
    key: 'name',
  },
]
const baseApiUrl = import.meta.env.VITE_API_URL?import.meta.env.VITE_API_URL:'';


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
  loadFolderData()
}


const handleSizeChange = async (e) => {
  let response = await Api.fileController.index({currentPage: e, pageSize: 24})
  files.value = response.data
  pageForm.value.total = response.total
  pageForm.value.page = response.meta?.current_page
};

const handleCurrentChange = (val) => {
  pageForm.value.page = val;
  loadFolderData()
};
// 上传文件显示
const uploadVisible = ref(false)
//分类id
const cate_id = ref(null)
// 上传文件
const handleUpload = (id) => {
  uploadVisible.value = true
  cate_id.value = id;
}

const emit = defineEmits(['open-folder', 'delete-folder', 'download-file', 'delete-file']);

const file = ref({});

const loadFolderData = async (id) => {
  if (id) {
    const response = await Api.fileCateController.files({id: id})
    files.value = response.data
    pageForm.value.total = response.total
    pageForm.value.page = response.meta?.current_page
  } else {
    let response = await Api.fileController.index({pageSize: 24})
    files.value = response.data
    pageForm.value.total = response.total
    pageForm.value.page = response.meta?.current_page
  }
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
const getFile = () => {
  return file.value
}
defineExpose({
  getFile,
  searchBy
})

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
  let url = `/uploads/${path}`
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

</script>
<style>
.a-dropdown-link {
  cursor: pointer;
  color: #409EFF;
}

.a-icon-arrow-down {
  font-size: 12px;
}
</style>
