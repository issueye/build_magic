<template>
    <div class="h-[calc(100% - 46px)] bg-white bs-main-body">
        <div :class="usePadding ? paddingClass : ''" :style="{height: `calc(100% - ${diffHeight}px)`}">
            <slot name="body"></slot>
        </div>
    </div>
</template>

<script setup lang="ts">
import { toRefs, ref } from 'vue';

const props = defineProps({
    usePadding: {
        type: Boolean,
        default: true,
    },
    diffHeight: {
        type: Number,
        default: 65,
    },
});

const paddingClass = ref('p-3')

const { usePadding, diffHeight } = toRefs(props);
</script>

<style lang="scss" scoped>
.bs-main-body {
    height: calc(100% - 46px);
    background-color: #fff;
    border-radius: 0 0 2px 2px;

    .body-box {
        padding: 10px;
        height: calc(100% - 59px);
    }

    .footer-box {
        padding: 10px;
        border-top: 1px solid #d9d9d9;
    }
}
</style>