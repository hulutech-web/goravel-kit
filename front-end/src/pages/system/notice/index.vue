<script lang="ts" setup>
import {ref} from 'vue'
import {message} from 'ant-design-vue'
import type {VxeGridListeners} from 'vxe-table'
import useNotice from "@/composible/useNotice";
import {useRouter} from 'vue-router';

const router = useRouter();
const {gridOptions} = useNotice()
const xGrid = ref()

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