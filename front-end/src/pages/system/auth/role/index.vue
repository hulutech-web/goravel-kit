<script setup lang="ts">
import useRole from "@/composible/useRole";
import Api from "@/api";
import {VxeGridListeners} from "vxe-table";

const {gridOptions} = useRole()
// TABLE
const xGrid = ref()

interface RowVO {
}

const gridEvent: VxeGridListeners<RowVO> = {
  proxyQuery() {
    console.log('数据代理查询事件')
    const grid = xGrid.value
    // 获取表格中的数据
    const data = grid.getTableData().fullData
  },
  proxyDelete() {
    console.log('数据代理删除事件')
  },
  proxySave() {
    console.log('数据代理保存事件')
  }
}
/*
* 设置权限 */
const open = ref(false)
const role_id = ref(null)
const setPer = async (row) => {
  open.value = true;
  role_id.value = row.id
  await loadPers(row.id)
}
const onClose = () => {
  open.value = false;
}
const tableRef = ref()

const formIDs = ref([])
const permissionOpt = ref([])
const currentID = ref(null)
const loadPers = async (id) => {
  currentID.value = id;
  let d = await Api.roleController.permissions({id: id})
  formIDs.value = d.map(item => item.id)
}

const defaultPers = ref([])
const loadAllPers = async () => {
  permissionOpt.value = await Api.permissionController.option()
}
const onSubmit = async () => {
  await Api.roleController.syncPermissions({id:currentID.value},{formIDs:formIDs.value})
}

loadAllPers()
</script>

<template>
  <div>
    <vxe-grid ref='xGrid' v-bind="gridOptions" v-on="gridEvent">
      <template #action="{ row }">
        <div>
          <a-button type="primary">删除</a-button>
          <a-button type="primary">编辑</a-button>
          <a-button type="primary" @click="setPer(row)">权限</a-button>
        </div>
      </template>
    </vxe-grid>
    <a-drawer title="权限" size="large" v-model:visible="open" @close="onClose">
      <template #extra>
        <a-button style="margin-right: 8px" @click="onClose">取消</a-button>
        <a-button type="primary" @click="onSubmit">确认</a-button>
      </template>
      <div>
        <a-checkbox-group v-model:value="formIDs">
          <a-row :gutter="[16,16]">
            <a-col :span="12" v-for="(op,inde) in permissionOpt" :key="inde">
              <a-checkbox :value="op.value">
                {{ op.label }}
                <span class="text-sm text-blue-500">{{ op.code }}</span>
              </a-checkbox>
            </a-col>
          </a-row>
        </a-checkbox-group>
      </div>
    </a-drawer>
  </div>
</template>

<style scoped lang="less">
.expand-wrapper {
  padding: 20px;
}
</style>