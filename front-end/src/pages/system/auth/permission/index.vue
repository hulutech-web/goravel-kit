<template>
  <div>
    <vxe-toolbar ref="toolbarRef" :buttons="toolbarButtons" @button-click="buttonClickEvent"></vxe-toolbar>
    <vxe-table
        border
        ref="tableRef"
        :edit-config="{mode: 'row', trigger: 'click'}"
        :column-config="columnConfig"
        :tree-config="treeConfig"
        :data="persForms.permissions">
      <vxe-column field="name" title="名称" :edit-render="{ name: 'AInput' }">
      </vxe-column>
      <vxe-column field="code" title="编码(与菜单【权限校验】一致)" :edit-render="{ name: 'AInput' }"></vxe-column>
      <vxe-column field="type" title="类型" :edit-render="{ name: 'AInput' }"></vxe-column>
      <vxe-column field="menu_id" title="菜单" :edit-render="{ name: 'AInput' }"></vxe-column>
    </vxe-table>
  </div>
</template>

<script setup lang="ts">
import Api from "@/api";
import {VxeToolbarEvents, VxeToolbarPropTypes} from "vxe-pc-ui/types/components/toolbar";

const treeConfig = reactive({
  transform: true,
  rowField: 'id',
  parentField: 'parentId'
})

const persForms = ref({
  permissions: []
})
const tableRef = ref(null)
const perForm = ref({
  id: null,
  name: "",
  code: "",
  type: 1,
  menu_id: null,
})

const toolbarButtons = ref<VxeToolbarPropTypes.Buttons>([
  // {name: '新增', code: 'add', status: 'primary'},
  // {name: '删除', code: 'del', status: 'error'},
  {name: '保存', code: 'save', status: 'success'}
])
const buttonClickEvent: VxeToolbarEvents.ButtonClick = async (params) => {
  console.log(params.$event.target)
  switch (params.code) {
    case "add":
      persForms.value.permissions.push(perForm.value)
      break;
    case "del":
      persForms.value.permissions = persForms.value.permissions.filter(i => i.id !== perForm.value.id)
      break;
    case "save":
      //获取编辑过的数据
      await Api.permissionController.update({id: id.value}, chapterForms.value)
      tableRef.value.commitProxy('query')
      break;
  }
}
const columnConfig = reactive({
  resizable: true
})
const initPers = async () => {
  // @ts-ignore
  persForms.value.permissions = await Api.permissionController.list()
  console.log( persForms.value.permissions)
}
initPers()
</script>

<style scoped>

</style>