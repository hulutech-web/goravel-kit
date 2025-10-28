<template>
  <div style="min-width: 400px;display: flex">
    <a-select v-model:value="user_id" :filter-option="false"
              :not-found-content="state.fetching ? undefined : null"
              :disabled="disabled"
              allow-clear
              style="width: 100%"
              @change="change" show-search @search="handleSearch"
              :options="selEmployees">
      <template v-if="state.fetching" #notFoundContent>
        <a-spin size="small"/>
      </template>
    </a-select>
    <slot name="default">

    </slot>
  </div>
</template>

<script setup lang="ts">
import {debounce} from 'lodash-es';
import Api from "@/api"

const emit = defineEmits(["update:modelValue"]);
/*定义组件名字*/
const props = defineProps({
  modelValue: {
    type: Number,
    default: null,
  },
  disabled: {
    type: Boolean,
    default: false
  }
});
const user_id = ref(props.modelValue);
const selEmployees = ref([]);


let lastFetchId = 0;
const state = reactive({
  data: [],
  value: [],
  fetching: false,
});
const handleSearch = debounce(async (value) => {
  // console.log('fetching user', value);
  state.fetching = true;
  let data = await Api.userController.option({username: value});
  if (data) {
    selEmployees.value = data;
  }
  // console.log(selEmployees.value)
  state.fetching = false;
},);


watch(
    () => props.modelValue,
    (newValue) => {
      if (newValue) {
        user_id.value = newValue;
        handleSearch();
      }
    },
    {
      deep: true,
      immediate: true,
    }
);
const change = (val) => {
  emit("update:modelValue", val);
};
</script>

<style scoped></style>
