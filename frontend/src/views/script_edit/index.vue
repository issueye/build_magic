<template>
  <BsHeader :title="mdTitle" description="生成对应代码的模板脚本">
    <template #actions>
      <div class="flex w-[400px] justify-end">
        <el-button class="mr-5" type="primary" @click="onTestRunClick">测试</el-button>
        <el-button @click="onBackClick">返回</el-button>
      </div>
    </template>
  </BsHeader>
  <BsMain :usePadding="false" :diffHeight="44">
    <template #body>
      <div class="flex h-full">
        <div class="h-full w-[20%] px-1" style="border-right: 1px solid #d9d9d9;">
          <CodeTree  />
        </div>
        <div class="flex flex-col h-full w-[80%]">
          <div ref="editorContainer" style="height: 70%"></div>
          <div class="h-[30%]">
            <Codemirror v-model:value="logData" :options="cmOptions" />
          </div>
        </div>
      </div>
    </template>
  </BsMain>
</template>

<script setup lang="ts">
import { ref, onMounted, toRaw, Ref } from "vue";
import * as monaco from "monaco-editor";
import { useRouter, useRoute } from "vue-router";
import { GetCode, SaveCode } from "@/wailsjs/go/controller/Template";
import { ElMessage } from "element-plus";
import CodeTree from './components/tree/index.vue';

import { TestRunCode } from "@/wailsjs/go/controller/Code";

import Codemirror, { createLog } from "codemirror-editor-vue3";
import { EventsOn, EventsOff } from "@/wailsjs/runtime/runtime";

const router = useRouter();
const route = useRoute();

const dmId = ref("");

const id = ref(route.query["id"]);
const mdTitle = ref(route.query["title"]);

const editorContainer = ref<any>(null);
const editor = ref<any>(null);

onMounted(async () => {
  const code = await GetCode(id.value as string);
  
var initCode = `export function main() {
    /// 你的代码...

    return 'return data'
}`

  editor.value = getEditer(code || initCode, editorContainer);

  // 设置快捷键 ctrl + s
  editor.value.addAction({
    id: "save",
    label: "保存",
    keybindings: [monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyS],
    run: () => {
      let code = toRaw(editor.value).getValue();
      // console.log("code", code);
      SaveCode(id.value as string, code).then(() => {
        ElMessage.success("保存成功");
      });
    },
  });
});

const logData = ref("=========测试日志==========\n");

const cmOptions = {
  mode: "fclog",
  theme: "default",
  lineWrapping: true,
  foldGutter: true,
};

const getEditer = (
  code: string,
  editor: Ref<any>,
  readOnly: boolean = false
) => {
  return monaco.editor.create(editor.value, {
    value: code,
    language: "javascript",
    folding: true, // 是否折叠
    foldingHighlight: true, // 折叠等高线
    foldingStrategy: "indentation", // 折叠方式  auto | indentation
    showFoldingControls: "always", // 是否一直显示折叠 always | mouseover
    disableLayerHinting: true, // 等宽优化
    emptySelectionClipboard: false, // 空选择剪切板
    selectionClipboard: false, // 选择剪切板
    automaticLayout: true, // 自动布局
    codeLens: false, // 代码镜头
    scrollBeyondLastLine: false, // 滚动完最后一行后再滚动一屏幕
    colorDecorators: true, // 颜色装饰器
    accessibilitySupport: "off", // 辅助功能支持  "auto" | "off" | "on"
    lineNumbers: "on", // 行号 取值： "on" | "off" | "relative" | "interval" | function
    lineNumbersMinChars: 5, // 行号最小字符   number
    readOnly: readOnly, //是否只读  取值 true | false
    theme: "vs-dark",
    minimap: {
      enabled: true, // 是否启用预览图
    },
  });
};

const onBackClick = () => {
  router.push("/project");
};

const onTestRunClick = async () => {
  try {
    EventsOn("console", (data: any) => {
      logData.value += `${createLog(`${data}\n`, "info")}`;
    });

    await TestRunCode(dmId.value, id.value as string);
  } finally {
    console.log("结束，关闭事件");
    EventsOff("console");
  }
};
</script>
