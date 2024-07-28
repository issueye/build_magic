<script setup lang="ts">
import { Icon } from "@iconify/vue";
import {
  WindowMinimise,
  WindowMaximise,
  WindowUnmaximise,
  Quit,
} from "@/wailsjs/runtime/runtime";
import emitter from "./utils/eventbus";

import Docs from "@/views/docs/index.vue";
import { ref } from "vue";
import { Auth } from "./store/store";

// 侧边栏导航
const menu = [
  { text: "项目", href: "/project", icon: "codicon:project" },
  { text: "模板", href: "/templates", icon: "carbon:template" },
];
// 侧边栏导航激活样式
const activeClass = "text-blue-600 dark:text-blue-400";
// 是否非 Mac 平台
const isNotMac = navigator.userAgent.toUpperCase().indexOf("MAC") < 0;
// 是否最大化
const isMaximised = ref(false);
</script>

<template>
  <main class="flex h-screen">
    <!-- 侧边栏导航，未登录时不显示 -->
    <nav
      v-show="Auth.logged"
      class="flex-none flex flex-col justify-between w-[48px] items-center text-center select-none z-20 bg-gray-50/10 dark:bg-slate-900/80"
      style="--wails-draggable: drag"
    >
      <div
        class="w-full mt-12 my-4 flex flex-col gap-6 text-2xl text-gray-500 dark:text-gray-200"
      >
        <div v-for="item in menu" :key="item.text">
          <el-tooltip effect="dark" :content="item.text" placement="right">
            <router-link :to="item.href" v-slot="{ isActive }">
              <Icon :icon="item.icon" :class="isActive && activeClass" />
            </router-link>
          </el-tooltip>
        </div>
      </div>
      <div
        class="my-4 flex flex-col gap-4 text-2xl text-gray-500 dark:text-gray-200"
      >
        <el-tooltip effect="dark" content="文档" placement="right">
          <router-link v-slot="{ isActive }">
            <Icon
              icon="hugeicons:document-code"
              :class="isActive && activeClass"
              @click="emitter.emit('show-docs')"
            />
          </router-link>
        </el-tooltip>
        <!-- <el-tooltip effect="dark" content="设置" placement="right">
          <router-link to="/setup" v-slot="{ isActive }">
            <Icon
              icon="ant-design:setting-outlined"
              :class="isActive && activeClass"
            />
          </router-link>
        </el-tooltip> -->
      </div>
    </nav>
    <!-- 内容面板 -->
    <div class="w-[calc(100%-48px)] bg-white dark:bg-gray-900/90">
      <!-- windows 定制化窗口按钮 -->
      <div
        v-if="isNotMac"
        class="flex h-8 bg-[#DAE5EC]"
        style="--wails-draggable: drag"
      >
        <div class="grow"></div>
        <div class="flex">
          <Icon
            icon="mdi:window-minimize"
            class="px-1 w-6 h-8 flex flex-col items-center justify-center hover:bg-[#E9E9E9] dark:hover:bg-[#2D2D2D] hover:text-gray-500"
            @click="WindowMinimise"
            style="--wails-draggable: none"
          />
          <Icon
            icon="mdi:window-maximize"
            class="px-1 w-6 h-8 text flex flex-col items-center justify-center hover:bg-[#E9E9E9] dark:hover:bg-[#2D2D2D]"
            @click="
              isMaximised
                ? (WindowUnmaximise(), (isMaximised = false))
                : (WindowMaximise(), (isMaximised = true))
            "
            style="--wails-draggable: none"
          />
          <Icon
            icon="mdi:window-close"
            class="px-1 w-6 h-8 text flex flex-col items-center justify-center hover:bg-[#C13124] dark:hover:bg-[#C13124] hover:text-white"
            @click="Quit"
            style="--wails-draggable: none"
          />
        </div>
      </div>
      <div v-else class="h-8"></div>
      <!-- 页面内容 -->
      <router-view
        class="overflow-y-auto h-screen"
        style="--wails-draggable: none"
      />
    </div>
  </main>

  <Docs />
</template>
