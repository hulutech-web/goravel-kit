<template>
  <div>
    <a-card>
      <p class="text-md font-bold">数据库表</p>
      <vxe-table
          border
          size="small"
          :data="tableData">
        <vxe-column field="table_schema" title="数据库名称"></vxe-column>
        <vxe-column field="table_name" title="表名称"></vxe-column>
        <vxe-column field="engine" width="80" title="引擎"></vxe-column>
        <vxe-column field="version" width="80" title="版本号"></vxe-column>
        <vxe-column field="row_format" title="行格式"></vxe-column>
        <vxe-column field="table_rows" title="表行数"></vxe-column>
        <vxe-column field="avg_row_length" title="行平均长"></vxe-column>
        <vxe-column field="data_length" title="数据文件长"></vxe-column>
        <vxe-column field="create_time" width="220" title="创建时间"></vxe-column>
        <vxe-column field="table_collation" title="规则"></vxe-column>
        <vxe-column field="table_comment" title="表注释"></vxe-column>
        <vxe-column field="action" title="管理">
          <template #default="{row}">
            <a-button type="primary" @click="manage(row)">管理</a-button>
          </template>
        </vxe-column>
      </vxe-table>
    </a-card>
  </div>
</template>

<script setup lang="ts">

import useTable from '@/composible/useTable'
import {useAuthStore} from "@/plugins";

const {useRole} = useAuthStore()
const router = useRouter()
const {getTable} = useTable()
const tableData = ref([])
const getData = async () => {
  tableData.value = await getTable()
}
getData()
const manage = useRole("admin", (row) => {
  router.push(`/system/crud/${row.table_name}/column`)
})

</script>

<style scoped>

</style>