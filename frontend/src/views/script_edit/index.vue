<template>
  <BsHeader :title="mdTitle" description="生成对应代码的模板脚本">
    <template #actions>
      <div class="flex w-[400px] justify-end">
        <div class="flex grow mr-10">
          <el-select v-model="dmId" class="mr-1" placeholder="请选择数据模型">
            <el-option v-for="(item, index) in tableData" :key="index" :label="item.title" :value="item.id" />
          </el-select>
          <el-button type="primary" @click="onTestRunClick">测试</el-button>
        </div>
        <el-button @click="onBackClick">返回</el-button>
      </div>
    </template>
  </BsHeader>
  <BsMain :usePadding="false" :diffHeight="44">
    <template #body>
      <div class="flex h-full w-full">
        <div class="w-[240px]" style="border-right: 1px solid #d9d9d9;">
          <el-tree :data="schemeTree" class="w-[calc(100% - 5px)] py-[5px]" node-key="code" highlight-current
            :expand-on-click-node="false" accordion @node-click="onNodeClick">
            <template #default="{ data }">
              <div class="flex items-center justify-between w-full">
                <span class="flex grow items-center">
                  <Icon :icon="data.icon" class="mr-[2px]" v-if="data.icon" /><span>{{
                    data.title
                  }}</span>
                </span>
                <el-link type="primary" v-if="showInport(data)" @click.stop="onImportClick(data)" class="mr-1">导入</el-link>
              </div>
            </template>
          </el-tree>
        </div>
        <div class="flex flex-col h-full p-1" style="width: calc(100% - 240px);">
          <el-tabs v-model="activeName" @tab-click="onTabClick" class="h-full">
            <el-tab-pane :label="mdTitle" name="001" class="h-full">
              <div ref="editorContainer1" class="mr-1" style="height: 70%"></div>
              <div class="h-[30%]">
                <Codemirror v-model:value="logData" :options="cmOptions" />
              </div>
            </el-tab-pane>
            <el-tab-pane :label="selectTitle" name="002" class="h-full">
              <div ref="editorContainer2" class="mr-1 h-full"></div>
            </el-tab-pane>
          </el-tabs>
        </div>
      </div>

    </template>
  </BsMain>
</template>

<script setup lang="ts">
import { ref, onMounted, toRaw, Ref } from "vue";
import * as monaco from "monaco-editor";
import { useRouter, useRoute } from "vue-router";
import { GetCode, SaveCode, GetTreeByCode, GetTpByCode } from "@/wailsjs/go/main/Template";
import { ElMessage } from "element-plus";

import { List, TestRunCode } from "@/wailsjs/go/main/DataModel";

import Codemirror, { createLog } from "codemirror-editor-vue3";
import { model, repository } from "@/wailsjs/go/models";
import { EventsOn, EventsOff } from "@/wailsjs/runtime/runtime";

const router = useRouter();
const route = useRoute();

const dmId = ref("");
const tableData: Ref<model.DataModel[]> = ref([]);
const schemeTree: Ref<repository.SchemeTree[]> = ref([]);
const activeName = ref('001');

const id = ref(route.query["id"]);
const schemeCode = ref(route.query['schemeCode']);
const mdTitle = ref(route.query["title"]);
const selectCode = ref('');
const selectTitle = ref(mdTitle.value);

const getData = async () => {
  const data = await List("", 0, 0);
  tableData.value = data;
};

const getTree = async () => {
  // console.log('getTree -> schemeCode.value', schemeCode.value);
  const data = await GetTreeByCode(schemeCode.value as string);
  schemeTree.value = data;
};

// 获取代码
const getCodeBySchemeCode = async (value: string) => {
  const data = await GetTpByCode(value);
  const code = await GetCode(data.id);

  selectCode.value = code;
  toRaw(editor2.value).setValue(code);
};

onMounted(() => {
  getData();
  getTree();
});

const editorContainer1 = ref<any>(null);
const editorContainer2 = ref<any>(null);
const editor = ref<any>(null);
const editor2 = ref<any>(null);

onMounted(async () => {
  const code = await GetCode(id.value as string);

  editor.value = getEditer(code || '', editorContainer1);
  editor2.value = getEditer(selectCode.value, editorContainer2, true);

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

// 当组件销毁时，关闭 editer
// onUnmounted(() => {
//   editor.value.dispose();
// });


const getEditer = (code: string, editor: Ref<any>, readOnly: boolean = false) => {
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
}

const showInport = (data: repository.SchemeTree) => {
  // 类型为2 path 中只能包含 common

  if (data.type == 2) {
    return data.path.indexOf("common") > -1;
  }
  return false;
}

const onTabClick = (value: any) => {
  console.log(value);
}

const onBackClick = () => {
  router.push("/templates");
};

const onImportClick = (value: model.Scheme) => {
  console.log('value', value);
  let code = toRaw(editor.value).getValue();

  let path = value.path
  // 将 \ 替换为 /
  path = path.replace(/\\/g, "/");
  // 将文件后缀去除
  path = path.replace(/\.[^/.]+$/, "");
  // 将文件路径替换为导入的路径

  code = `import {  } from '${path}'\n${code}`

  toRaw(editor.value).setValue(code);

  activeName.value = "001"
}

const onNodeClick = (value: any) => {
  console.log('value', value);
  selectTitle.value = value.title;
  getCodeBySchemeCode(value.id);
  activeName.value = "002"
}

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

<style scoped>
::v-deep(.el-tabs__content) {
  height: calc(100% - 60px);
}
</style>