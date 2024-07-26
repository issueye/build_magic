<template>
  <el-tree
    :data="treeData"
    class="w-full py-[5px]"
    :expand-on-click-node="false"
    accordion
  >
    <template #default="{ data }">
      <div class="flex items-center justify-between w-full">
        <span class="flex grow items-center" @click="onNodeSelectClick(data)">
          <Icon :icon="data.icon" class="mr-[2px]" v-if="data.icon" /><span>{{
            data.title
          }}</span>
        </span>
        <span class="flex gap-1">
          <el-link
            type="primary"
            v-if="data.type !== 2"
            @click="onAddClick(data)"
            >添加</el-link
          >
          <el-link
            type="primary"
            v-if="data.type == 1"
            @click="onEditClick(data)"
            >编辑</el-link
          >
          <el-link
            type="danger"
            v-if="data.type !== 0"
            @click="onDeleteClick(data)"
            >删除</el-link
          >
        </span>
      </div>
    </template>
  </el-tree>

  <BsDialog
    :title="title"
    :width="500"
    :visible="visible"
    @close="onClose"
    @save="onSave"
    @open="onOpen"
  >
    <template #body>
      <el-form
        label-width="auto"
        :model="dataForm"
        :rules="rules"
        ref="dataFormRef"
      >
        <el-form-item label="名称" prop="title">
          <el-input
            v-model="dataForm.title"
            placeholder="请输入名称"
            clearable
          />
        </el-form-item>
        <el-form-item label="图标">
          <el-input
            v-model="dataForm.icon"
            placeholder="请输入图标"
            clearable
          />
        </el-form-item>
        <el-form-item label="类型" prop="nodeType">
          <el-select
            v-model="dataForm.nodeType"
            placeholder="请选择类型"
            :disabled="operationType == 1"
            clearable
          >
            <el-option label="文件夹" :value="1" />
            <el-option label="模板" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item
          label="文件类型"
          prop="fileType"
          v-if="dataForm.nodeType == 2"
        >
          <el-select
            v-model="dataForm.fileType"
            placeholder="请选择类型"
            :disabled="operationType == 1"
            clearable
          >
            <el-option label="javascript" :value="1" />
            <el-option label="typescript" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item label="备注">
          <el-input
            v-model="dataForm.mark"
            placeholder="请输入备注"
            type="textarea"
            :row="2"
            clearable
          />
        </el-form-item>
      </el-form>
    </template>
  </BsDialog>
</template>
<script setup lang="ts">
import { onMounted, ref } from "vue";
import { repository, model } from "@/wailsjs/go/models";
import {
  Create,
  Modify,
  DeleteByCode,
} from "@/wailsjs/go/controller/Template";
import { reactive } from "vue";
import { Icon } from "@iconify/vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { useScriptTemplateStore } from "@/store/script_template";
import { storeToRefs } from "pinia";

const scriptTemplateStore = useScriptTemplateStore();
const { selectNode, treeData } = storeToRefs(scriptTemplateStore);

const operationType = ref(0);

const title = ref("");
const visible = ref(false);
// 数据
const dataFormRef = ref();

const dataForm = reactive<model.CodeTemplate>(
  model.CodeTemplate.createFrom({
    id: "",
    title: "",
    fileName: "",
    icon: "",
    schemeCode: "",
    schemeParentCode: "",
    nodeType: 2,
    fileType: 1,
    mark: "",
  })
);

const resetForm = () => {
  dataForm.id = "";
  dataForm.title = "";
  dataForm.fileName = "";
  dataForm.icon = "";
  dataForm.schemeCode = "";
  dataForm.schemeParentCode = "";
  dataForm.nodeType = 2;
  dataForm.fileType = 1;
  dataForm.mark = "";
};

const setForm = (value: repository.SchemeTree) => {
  dataForm.id = value.id;
  dataForm.title = value.title;
  dataForm.fileName = "";
  dataForm.icon = value.icon;
  dataForm.schemeCode = value.id;
  dataForm.schemeParentCode = value.parentCode;
  dataForm.nodeType = value.type;
};

const rules = reactive({
  fileName: [{ required: true, message: "请输入文件名称", trigger: "blur" }],
});

onMounted(() => {
  getData();
});

const getData = async () => {
  await scriptTemplateStore.getTree();
};

const onNodeSelectClick = (value: repository.SchemeTree) => {
  selectNode.value = value.id;
  console.log("selectNode.value", selectNode.value);

  // 查询数据
  scriptTemplateStore.getData();
};

const onAddClick = (value: repository.SchemeTree) => {
  title.value = "添加";
  operationType.value = 0;
  resetForm();
  dataForm.schemeParentCode = value.parentCode;
  dataForm.schemeCode = value.id;
  visible.value = true;
};

const onEditClick = (value: repository.SchemeTree) => {
  title.value = "修改";
  operationType.value = 1;
  setForm(value);
  visible.value = true;
};

const onDeleteClick = (value: repository.SchemeTree) => {
  ElMessageBox.confirm("请确认是否要删除数据？", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      await DeleteByCode(value.id);
      getData();
    })
    .catch(() => {
      ElMessage.info("取消删除");
    });
};

const onClose = () => {
  visible.value = false;
};

const onSave = () => {
  if (!dataFormRef.value) return;
  dataFormRef.value.validate(async (valid: boolean) => {
    if (valid) {
      switch (operationType.value) {
        case 0: {
          await Create(dataForm);
          break;
        }
        case 1: {
          await Modify(dataForm);
          break;
        }
      }
      visible.value = false;

      getData();
    } else {
      console.log("表单验证失败");
    }
  });
};

const onOpen = () => {};
</script>