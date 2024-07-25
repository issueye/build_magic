<template>
  <div
    class="w-full h-[55px] flex justify-between items-center bg-white"
    style="border-bottom: 1px solid #e5e5e5"
  >
    <div>
      <span
        class="ml-5 text-[#262626] leading-8 text-base font-bold"
        style="vertical-align: middle"
        >{{ title }}</span
      >
      <el-divider direction="vertical" />
      <span class="text-[#BFBFBF] text-xs">{{ description }}</span>
    </div>
    <div class="mr-2.5">
      <slot name="actions"></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import { toRefs } from "vue";

const props = defineProps({
  title: {
    type: String,
    default: "",
  },
  description: {
    type: String,
    default: "",
  },
});

const { title, description } = toRefs(props);
</script>

