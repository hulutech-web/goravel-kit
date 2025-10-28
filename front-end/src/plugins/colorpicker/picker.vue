<template>
  <div class="flex">
    <div class="cursor-pointer">
      <pick-colors v-model:value="value" show-alpha :colors="colors" :size="50" theme="dark" :show-alpha="true"
                   :add-color="false" @change="colorChange" :width="props.width" :height="props.height" />
      <p>
        <span class="font-extrabold" :style="{ 'color': value }">颜色值:{{ value }}</span>
      </p>
    </div>
    <div style="width:200px;
        border-left-style: dashed;
        border-left-width: 1px;
        border-left-color: #ccc;
        border-right-style: dashed;
        border-right-width: 1px;
        border-right-color: #ccc;
        border-bottom-style: dashed;
        border-bottom-width: 1px;
        border-bottom-color: #ccc;
        border-top-width: 20px;
        border-top-style: solid;
        box-sizing: border-box;padding:8px;margin-left:10px;" :style="{ 'border-top-color': value }">
      <a-skeleton active avatar :paragraph="{ rows: 2 }" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { watch } from 'vue';
import PickColors from 'vue-pick-colors'
const props = defineProps({
  width: {
    type: Number,
    default: 60
  },
  height: {
    type: Number,
    default: 60
  },
  modelValue: {
    type: String,
    default: ''
  }
})
const emit = defineEmits(['update:modelValue'])
const value = ref("")
watch(() => props.modelValue, (newValue) => {
  value.value = newValue
}, {
  immediate: true
})
const colors = ref([
  'rgba(255, 69, 0, 0.68)',
  'rgb(255, 120, 0)',
  '#ff4500',
  '#ffd700',
  '#00ced1',
  '#1e90ff',
  '#c71585',
  'hsv(51, 100, 98)',
  'hsva(120, 40, 94, 0.5)',
  'hsl(181, 100%, 37%)',
  'hsla(209, 100%, 56%, 0.73)',
])
const colorChange = (color: string) => {
  emit("update:modelValue", color)
}
</script>

<style scoped></style>