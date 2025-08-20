<template>
  <div>
    <div class="flex justify-between items-center">
      <slot name="default">
        <div class="w-128 grid grid-cols-4 gap-4">
          <div v-for="(s,ind) in my_salary_plans" :key="ind">
            <a-badge status="success" :text="`${s.plan_name},${s.hourly_rate}/节`"/>
          </div>
        </div>
      </slot>
      <div class="w-48">
        <a-space>
          <a-button type="primary" @click="openModal = true">
            <template #icon>
              <edit-outlined/>
            </template>
          </a-button>
          <a-button type="primary" @click="onSubmit">
            <template #icon>
              <check-outlined/>
            </template>
          </a-button>
        </a-space>
      </div>
    </div>
    <div class="mt-1">
      <a-modal v-model:visible="openModal" width="600px" title="选择方案" centered @ok="()=>openModal=false">
        <vxe-table
            :row-config="{keyField:'value'}"
            size="small"
            border
            ref="tableRef"
            :data="dataSource"
            :checkbox-config="{checkRowKeys: checkedKeys}"
            @checkbox-all="selectAllChangeEvent"
            @checkbox-change="selectChangeEvent">
          <vxe-column type="checkbox"></vxe-column>
          <vxe-column field="value" title="序号"></vxe-column>
          <vxe-column field="label" title="方案" :width="400"></vxe-column>
        </vxe-table>
      </a-modal>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Api from "@/api"

const tableRef = ref(null)
const selectAllChangeEvent = () => {
}


const props = defineProps({
  salary_plans: {
    type: Array as PropType<Plan[]>,
    default: () => []
  },
  id: {
    type: Number,
    default: 0
  },
});


const checkedKeys = ref([])

const my_salary_plans = ref(props.salary_plans);

const selectChangeEvent = (row) => {
  const $table = tableRef.value;
  const records = $table.getCheckboxRecords()
  myOptions.value = records
  // emit("update:salary_plans", records);
}
const myOptions = ref([])
const dataSource = ref([]);


interface Plan {
  id: number;
  plan_name: string;
}

const fillData = async () => {
  dataSource.value = await Api.salaryPlanController.option();
  /*处理回显*/
  checkedKeys.value = props.salary_plans.map(item => item.id)
};
const openModal = ref(false);

/*监听openModal的执行，当为true时才开启网络加载数据*/
watch(openModal, (val) => {
  if (val) {
    fillData()
  }
})

const emit = defineEmits(['freshTable'])
const onSubmit = async () => {
  await Api.salaryPlanController.syncUser({
    id: props.id,
    opts: myOptions.value
  })
  /*刷新父表格*/
  console.log("刷新父表格1")
  emit('freshTable')
}
</script>

<style scoped></style>