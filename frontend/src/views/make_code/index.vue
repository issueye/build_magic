<template>
  <BsHeader
    :title="mdTitle"
    description="数据模型提供基础结构，通过绑定对应的代码模板生成代码"
  >
    <template #actions>
      <el-button @click="onBackClick">返回</el-button>
    </template>
  </BsHeader>
  <BsMain>
    <template #body>
      <el-tabs v-model="activeName" @tab-click="onTabClick">
        <el-tab-pane
          v-for="(item, index) in tableData"
          :label="item.title"
          :key="index"
          :name="item.id"
        >
        </el-tab-pane>
      </el-tabs>
      <div
        class="h-full border border-solid border-[#d9d9d9]"
        style="height: calc(100% - 60px)"
      >
        <Codemirror v-model:value="logData" :options="cmOptions" />
      </div>
    </template>
  </BsMain>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted, reactive, ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import Codemirror from "codemirror-editor-vue3";

import { GetChildrenByCode } from "@/wailsjs/go/main/Template";
import { Ref } from "vue";
import type { TabsPaneContext } from "element-plus";
import { EventsOn, EventsOff } from "@/wailsjs/runtime/runtime";
import { model } from "@/wailsjs/go/models";
import { RunCode } from "@/wailsjs/go/main/DataModel";
// import { VxeColumnPropTypes, VxeTableEvents } from "vxe-table";

// 数据
const router = useRouter();
const route = useRoute();

const activeName = ref();

const id = ref(route.query["id"]);
const mdTitle = ref(route.query["title"]);
const schemeId = ref(route.query["schemeId"]);
console.log("id", id.value);
console.log("schemeId", schemeId.value);

const logData = ref("");
const logDataMap = reactive<Record<string, string>>({});

const cmOptions = {
  mode: "fclog",
  theme: "default",
  lineWrapping: true,
  foldGutter: true,
};

const tableData: Ref<model.CodeTemplate[]> = ref([]);

const getData = async () => {
  const data = await GetChildrenByCode(schemeId.value as string);
  console.log("data", data);
  tableData.value = data;

  if (tableData.value.length > 0) {
    activeName.value = tableData.value[0].id;
  }

  let tData = [...tableData.value];
  tData.forEach((item) => {
    logDataMap[item.id] = "";
  });
};

interface codePushData {
  id: string;
  code_content: string;
}

onMounted(() => {
  getData();
  // 监听代码输入

  EventsOn("code_push", (data: codePushData) => {
    console.log("data", data);
    logDataMap[data.id] = data.code_content;
    console.log("logDataMap", logDataMap);

    if (logData.value == "") {
      logData.value = logDataMap[activeName.value as string];
    }
  });

  RunCode(id.value as string);
});

onUnmounted(() => {
  EventsOff("code_push");
});

const onBackClick = () => {
  router.push("/");
};

const onTabClick = (tab: TabsPaneContext, event: Event) => {
  console.log("value", tab, event);
  logData.value = logDataMap[tab.props.name as string];
  console.log("logData.value", logData.value);
};
</script>


<style scoped>
::v-deep(.el-tabs__content) {
  height: calc(100% - 60px);
}
</style>