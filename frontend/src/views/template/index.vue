<template>
  <BsHeader title="代码模板" description="生成对应代码的模板脚本">
    <template #actions>
      <!-- <el-button type="primary" @click="onAddClick">添加</el-button>  -->
    </template>
  </BsHeader>
  <BsMain :usePadding="false" :diffHeight="42">
    <template #body>
      <div class="w-full h-full flex">
        <div class="flex w-[300px] h-full" style="border-right: 1px solid #d9d9d9">
          <Tree />
        </div>
        <div class="flex flex-col w-[calc(100%-300px)] p-[10px]">
          <el-form inline>
            <el-form-item label="检索">
              <el-input v-model="form.condition" placeholder="请输入检索内容" clearable @change="onChange" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="onQryClick">查询</el-button>
            </el-form-item>
          </el-form>
          <div class="h-[calc(100% - 60px)]">
            <el-table border :data="tableData" size="small" highlight-current-row stripe>
              <el-table-column prop="id" label="编码" width="150" show-overflow-tooltip />
              <el-table-column prop="title" label="标题" width="150" show-overflow-tooltip />
              <el-table-column prop="fileName" label="文件名称" min-width="150" show-overflow-tooltip />
              <!-- <el-table-column prop="mark" label="备注" show-overflow-tooltip /> -->
              <el-table-column label="操作" width="130" align="center" fixed="right">
                <template v-slot="{ row }">
                  <el-button type="primary" link size="small" @click="onEditScriptClick(row)">编辑脚本</el-button>
                  <!-- <el-button type="primary" link size="small" @click="onEditClick(row)">编辑</el-button> -->
                  <el-button type="danger" link size="small" @click="onDeleteClick(row)">删除</el-button>
                </template>
              </el-table-column>
            </el-table>
            <div style="margin-top: 10px">
              <el-pagination small background :current-page="pageNum" :page-size="pageSize" :page-sizes="[5, 10, 20]"
                :total="total" layout="total, sizes, prev, pager, next" @size-change="onSizeChange"
                @current-change="onCurrentChange" />
            </div>
          </div>
        </div>
      </div>
    </template>
  </BsMain>

  <BsDialog :title="title" :width="500" :visible="visible" @close="onClose" @save="onSave" @open="onOpen">
    <template #body>
      <el-form label-width="auto" :model="dataForm" :rules="rules" ref="dataFormRef">
        <el-form-item label="标题" prop="title">
          <el-input v-model="dataForm.title" placeholder="请输入数据模型标题" clearable />
        </el-form-item>
        <el-form-item label="文件名称" prop="fileName">
          <el-input v-model="dataForm.fileName" placeholder="请输入文件名称" :disabled="true" clearable />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="dataForm.mark" placeholder="请输入备注" type="textarea" :row="2" clearable />
        </el-form-item>
      </el-form>
    </template>
  </BsDialog>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
// import { nanoid } from "nanoid";
import {
  Create,
  DeleteByCode,
  Modify,
} from "@/wailsjs/go/controller/Template";
import { model } from "@/wailsjs/go/models";
import { useRouter } from "vue-router";
import Tree from './components/tree/index.vue';
import { storeToRefs } from 'pinia';
import { useScriptTemplateStore } from "@/store/script_template"; 

// const nameTitle = "数据模型";
// 标题
const title = ref("数据模型");
// 显示弹窗
const visible = ref(false);
// 操作类型
const operationType = ref(0);
// 数据
const dataFormRef = ref();
// 路由
const router = useRouter();
// store
const scriptTemplateStore = useScriptTemplateStore();
const { pageNum, pageSize, total, tableData } = storeToRefs(scriptTemplateStore);

const rules = reactive({
  name: [{ required: true, message: "请输入姓名", trigger: "blur" }],
  account: [{ required: true, message: "请输入账户", trigger: "blur" }],
  groupId: [{ required: true, message: "请选择用户组", trigger: "blur" }],
});

// 查询表单
const form = reactive({
  condition: "",
});

// 弹窗表单
const dataForm = reactive<model.CodeTemplate>(
  model.CodeTemplate.createFrom({
    id: "",
    title: "",
    fileName: "",
    fileType: "",
    schemeCode: "",
    schemeParentCode: "",
    mark: "",
  })
);

onMounted(async () => {
  
});

const getData = async () => {
  await scriptTemplateStore.getData(form.condition);
  await scriptTemplateStore.getTree();
};

// 重置表单数据
// const resetForm = () => {
//   let id = nanoid();
//   dataForm.id = "";
//   dataForm.title = "";
//   dataForm.fileName = id;
//   dataForm.mark = "";
// };

// 赋值表单数据
// const setForm = (value: any) => {
//   dataForm.id = value.id;
//   dataForm.title = value.title;
//   dataForm.fileName = value.fileName;
//   dataForm.schemeCode = value.schemeCode;
//   dataForm.schemeParentCode = value.schemeParentCode;
//   dataForm.mark = value.mark;
// };

onMounted(() => {
  getData();
});

// 事件
const onChange = () => {
  getData();
};

const onQryClick = () => {
  getData();
};

const onEditScriptClick = (value: any) => {
  // console.log("value", value);
  console.log('value.schemeCode', value.schemeCode);
  router.push({
    path: "/script_edit",
    query: { id: value.id, title: value.title, schemeCode: value.schemeCode },
  });
};

const onDeleteClick = (value: model.CodeTemplate) => {
  // console.log("value", value.value);
  ElMessageBox.confirm("请确认是否要删除数据？", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      await DeleteByCode(value.schemeCode);
      getData();
    })
    .catch(() => {
      ElMessage.info("取消删除");
    });
};

const onSave = () => {
  if (!dataFormRef.value) return;
  dataFormRef.value.validate(async (valid: boolean) => {
    if (valid) {
      switch (operationType.value) {
        case 0: {
          await Create(dataForm);
          visible.value = false;
          getData();
          break;
        }

        case 1: {
          await Modify(dataForm);
          // ElMessage.success(res.message);
          visible.value = false;
          getData();
          scriptTemplateStore.getTree();
          break;
        }
      }
    } else {
      console.log("表单验证失败");
    }
  });
};

// 弹窗打开时
const onOpen = async () => {
  // const res = await apiUserGroupList();
  // if (res.code === 200) {
  //   groupTableData.value = res.data;
  // }
};

const onSizeChange = (value: number) => {
  pageSize.value = value;
  getData();
};

const onCurrentChange = () => {
  getData();
};

const onClose = () => {
  visible.value = false;

  if (!dataFormRef.value) return;
  dataFormRef.value.resetFields();
};
</script>
