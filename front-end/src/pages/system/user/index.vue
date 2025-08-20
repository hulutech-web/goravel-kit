<script lang="ts" setup>
import {ref} from 'vue'
import {message} from 'ant-design-vue'
import type {VxeGridListeners} from 'vxe-table'
import useUser from "./index.ts"
import Api from "@/api";

const router = useRouter();
const {gridOptions} = useUser()
const xGrid = ref()
const serveApiUrl = import.meta.env.VITE_API_URL;

// 编辑用户
const editCourse = (row: any) => {
  const grid = xGrid.value
  grid.setEditRow(row)
}




const gridEvent: VxeGridListeners = {
  proxyQuery() {
    //设置column为cover的列，将其数据设置为数组，数组形式如下：[{name:'xx',url:'xx'}...]
  },
}


const fileRef = ref(null)
const selectFile = () => {
  fileRef.value.showModal()
  //通过xGrid获取到row的列，并且把cover数据改为
}

const toChapter = (row) => {
  router.push({path: `/course/${row.id}/chapter`})
}
</script>

<template>
  <div class="demo-page-wrapper">
    <vxe-grid
        ref="xGrid"
        v-bind="gridOptions"
        v-on="gridEvent"
    >

      <template #avatar="{ row }">
        <a-space>
          <vxe-image :src="`${serveApiUrl}/uploads/${row.avatar}`" width="64" height="64"></vxe-image>
          <FileSelect ref="fileRef" v-model="row.avatar">
          </FileSelect>
        </a-space>
      </template>
      <template #action="{ row }">
        <a-space>
          <a-popconfirm
              title="确认删除吗？"
              ok-text="是"
              cancel-text="否"
              @confirm="delCourse(row)"
          >
            <a-button size="small" type="primary" danger>删除</a-button>
          </a-popconfirm>
          <a-button type="primary" size="small" @click="toChapter(row)">章节</a-button>
        </a-space>
      </template>
    </vxe-grid>
  </div>
</template>

<style scoped lang="less">
.demo-page-wrapper {
  padding: 20px;
}
</style>