<script lang="ts" setup>
import {ref} from 'vue'
import {message} from 'ant-design-vue'
import type {VxeGridListeners} from 'vxe-table'
import useVenue from "./useVenue";

const serveApiUrl = import.meta.env.VITE_API_URL;
const router = useRouter();
const {gridOptions} = useVenue()
const xGrid = ref()

const setAccount = (row) => {
  router.push({path: `/system/tenant/${row.id}/account`})
}
// 编辑用户
const editCourse = (row: any) => {
  const grid = xGrid.value
  grid.setEditRow(row)
}

// 删除用户
const delCourse = async (row: any) => {
  try {
    /*    await destroy({id: row.id})*/
    xGrid.value.commitProxy('query')
  } catch (error) {
    message.error('删除失败')
  }
}


const setMap = (row) => {
  router.push({path: `/venue/${row.id}/map`})
}
const gridEvent: VxeGridListeners = {
  proxyQuery() {
    /*设置选项*/
    const column1 = xGrid.value.getColumnByField('course_type_id')
  },
}
</script>

<template>
  <div class="demo-page-wrapper">
    <vxe-grid
        ref="xGrid"
        v-bind="gridOptions"
        v-on="gridEvent"
    >
      <template #cover="{ row }">
        <a-space>
          <vxe-image :src="`${serveApiUrl}/uploads/${row.cover}`" width="64" height="64"></vxe-image>
          <FileSelect ref="fileRef" v-model="row.cover">
          </FileSelect>
        </a-space>
      </template>

      <template #action="{ row }">
        <a-space>
          <a-button type="primary" size="small" @click="editCourse(row)">编辑</a-button>
          <a-popconfirm
              title="确认删除吗？"
              ok-text="是"
              cancel-text="否"
              @confirm="delCourse(row)"
          >
            <a-button size="small" type="primary" danger>删除</a-button>
          </a-popconfirm>
          <a-button size="small" danger type="primary" @click="setMap(row)">地图</a-button>
          <a-button size="small" danger type="primary" @click="setAccount(row)">开设账号</a-button>
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