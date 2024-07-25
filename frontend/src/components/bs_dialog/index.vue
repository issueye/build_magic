<template>
  <el-dialog
    draggable
    v-model="props.visible"
    :title="props.title"
    :width="props.width"
    :before-close="onClose"
    :close-on-click-modal="false"
    @opened="onOpen"
  >
    <div class="p-[20px]">
      <slot name="body"> </slot>
    </div>
    <div class="flex p-[20px] justify-end" v-if="props.showFooter">
      <slot name="footer">
        <el-button type="primary" @click="submitForm">确定</el-button>
        <el-button @click="onClose">关闭</el-button>
      </slot>
    </div>
  </el-dialog>
</template>

<script setup>
const props = defineProps({
  title: {
    type: String,
    default: "",
  },
  visible: {
    type: Boolean,
    default: false,
  },
  showFooter: {
    type: Boolean,
    default: true,
  },
  width: 500,
});

const emits = defineEmits(["close", "open", "save"]);

const onClose = () => {
  emits("close");
};

const onOpen = () => {
  console.log("dialog", "open event");
  emits("open");
};

const submitForm = () => {
  emits("save");
};
</script>