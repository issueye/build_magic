<script lang="ts" setup>
import { onBeforeMount } from "vue";
import router from "@/router";
import { Auth } from "@/store/store";

// 初始化判断本地是否有 token 然后跳转到对应页面
onBeforeMount(() => {
  router.push("/project");
  Auth.logged = true;
});
</script>

<template></template> 
