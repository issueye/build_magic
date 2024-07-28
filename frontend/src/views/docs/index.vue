<template>
  <BsDialog
    title="脚本文档"
    :width="900"
    :visible="visible"
    @close="onClose"
    @save="onSave"
    :showFooter="false"
    :showPadding="false"
  >
    <template #body>
      <div class="h-[60vh] flex">
        <div class="w-[200px]" style="border-right: 1px solid #d9d9d9">
          <el-tree
            :data="docsTree"
            :props="defaultProps"
            accordion
            @node-click="handleNodeClick"
          />
        </div>
        <div class="grow overflow-y-auto">
          <v-md-preview :text="docs"></v-md-preview>
        </div>
      </div>
    </template>
  </BsDialog>
</template>

<script setup lang="ts">
import emitter from "@/utils/eventbus";
import { ref, onMounted } from "vue";
import { GetDocsTree, GetDocsContent } from "@/wailsjs/go/controller/Docs";

const docs = ref("");
const visible = ref(false);
const docsTree = ref();

const defaultProps = {
  children: "children",
  label: "label",
};

const onClose = () => {
  visible.value = false;
};

const onSave = () => {
  visible.value = false;
};

onMounted(() => {
  emitter.on("show-docs", async () => {
    // 加载文档目录
    let data = await GetDocsTree();
    // console.log("tree", data);
    let treeObj = JSON.parse(data);
    // console.log("treeObj", treeObj);
    docsTree.value = treeObj;
    visible.value = true;
  });

  emitter.on("close-docs", () => {
    visible.value = false;
  });
});

const handleNodeClick = async (data: any) => {
  if (data.file) {
    let content = await GetDocsContent(data.file);
    // console.log("content", content);
    docs.value = content;
  }
};
</script>