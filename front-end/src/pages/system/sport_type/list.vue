<script lang="ts" setup>
import {ref} from 'vue'
import {message} from 'ant-design-vue'
import type {VxeGridListeners} from 'vxe-table'
import {destroy} from "@/api/sportTypeController";
import useSportType from "@/composible/useSportType"

const serveApiUrl = import.meta.env.VITE_API_URL;

const {gridOptions} = useSportType()
const xGrid = ref()
const router = useRouter()
// 编辑用户
const editCourse = (row: any) => {
  const grid = xGrid.value
  grid.setEditRow(row)
}

// 删除用户
const delCourse = async (row: any) => {
  try {
    await destroy({id: row.id})
    xGrid.value.commitProxy('query')
  } catch (error) {
    message.error('删除失败')
  }
}


const gridEvent: VxeGridListeners = {
  proxyQuery() {
    /*设置选项*/
  },
}

const memberAction = (row) => {
  router.push({path: `/team/${row.id}/manage`})
}
</script>

<template>
  <div class="demo-page-wrapper">
    <vxe-grid
        ref="xGrid"
        v-bind="gridOptions"
        v-on="gridEvent"
    >
      <template #icon="{ row }">
        <a-space>
          <vxe-image :src="`${serveApiUrl}/uploads/${row.icon}`" width="64" height="64"></vxe-image>
          <FileSelect ref="fileRef" v-model="row.icon">
          </FileSelect>
        </a-space>
      </template>
      <template #is_active="{ row }">
        <a-switch v-model:checked="row.is_active"
                  checked-children="开"
                  un-checked-children="关" />
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