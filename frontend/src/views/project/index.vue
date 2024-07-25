<template>
  <BsHeader title="项目列表" description="项目列表">
    <template #actions>
      <el-button type="primary" @click="onAddClick">添加</el-button>
    </template>
  </BsHeader>
  <BsMain>
    <template #body>
      <div>
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
            <el-table-column prop="name" label="项目名称" width="150" show-overflow-tooltip />
            <el-table-column prop="version" label="版本" width="150" show-overflow-tooltip />
            <el-table-column prop="build_type" label="构建类型" width="150" show-overflow-tooltip />
            <el-table-column prop="build_version" label="构建版本" width="150" show-overflow-tooltip />
            <el-table-column prop="project_path" label="项目路径" min-width="200" show-overflow-tooltip />
            <el-table-column prop="mark" label="备注" min-width="200" show-overflow-tooltip />
            <el-table-column label="操作" width="150" align="center" fixed="right">
              <template v-slot="{ row }">
                <div class="flex items-center">
                  <el-button type="primary" link size="small" @click="onEditClick(row)">编辑</el-button>
                  <el-button type="danger" link size="small" @click="onDeleteClick(row)">删除</el-button>
                  <el-dropdown @command="handleCommand" class="h-[16px] flex items-center ml-[12px]">
                    <span
                      class="flex items-center text-[12px] text-[--el-color-primary] hover:text-[--el-color-primary-light-3] outline-none currsor-pointer">更多
                      <el-icon class="el-icon--right"><arrow-down /></el-icon></span>
                    <template #dropdown>
                      <el-dropdown-menu>
                        <el-dropdown-item :command="{ data: row, type: 'build_code' }">构建代码</el-dropdown-item>
                        <el-dropdown-item :command="{ data: row, type: 'injet_variable' }">注入变量</el-dropdown-item>
                        <el-dropdown-item :command="{ data: row, type: 'env_variable' }">环境变量</el-dropdown-item>
                      </el-dropdown-menu>
                    </template>
                  </el-dropdown>
                </div>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <div style="margin-top: 10px">
          <el-pagination small background :current-page="pageNum" :page-size="pageSize" :page-sizes="[5, 10, 20]"
            :total="total" layout="total, sizes, prev, pager, next" @size-change="onSizeChange"
            @current-change="onCurrentChange" />
        </div>
      </div>
    </template>
  </BsMain>

  <BsDialog :title="title" :width="500" :visible="visible" @close="onClose" @save="onSave">
    <template #body>
      <el-form label-width="auto" :model="dataForm" :rules="rules" ref="dataFormRef">
        <el-form-item label="标题" prop="title">
          <el-input v-model="dataForm.title" placeholder="请输入标题" clearable />
        </el-form-item>
        <el-form-item label="项目名称" prop="name">
          <el-input v-model="dataForm.name" placeholder="请输入项目名称" clearable />
        </el-form-item>
        <el-form-item label="项目路径" prop="project_path">
          <el-input v-model="dataForm.project_path" placeholder="请输入项目路径" clearable />
        </el-form-item>
        <el-form-item label="版本" prop="version">
          <el-input v-model="dataForm.version" placeholder="请输入版本" clearable />
        </el-form-item>
        <el-form-item label="构建类型" prop="build_type">
          <el-select v-model="dataForm.build_type" placeholder="请选择构建类型">
            <el-option :value="0" label="alpha" />
            <el-option :value="1" label="beta" />
            <el-option :value="2" label="RC" />
          </el-select>
        </el-form-item>
        <el-form-item label="构建版本" prop="build_version">
          <el-input-number v-model="dataForm.build_version" :min="1" :max="100" controls-position="right"
            placeholder="请输入构建版本" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="dataForm.mark" placeholder="请输入备注" type="textarea" :row="2" clearable />
        </el-form-item>
      </el-form>
    </template>
  </BsDialog>

  <BsDialog :title="variableTitle" :width="800" :visible="variableVisible" @close="onVariableClose"
    @save="onVariableSave" @open="onVariableOpen">
    <template #body>
      <vxe-table border show-overflow keep-source size="mini" ref="tableRef" :data="varTableData" max-height="550"
        empty-text="没有数据" :edit-config="{ trigger: 'click', mode: 'row', showStatus: true }">
        <vxe-column type="seq" title="序号" width="70" />
        <vxe-column field="key" title="参数名称" :edit-render="{ name: 'ElInput' }" />
        <vxe-column field="value" title="参数值" :edit-render="{ name: 'ElInput' }" />
        <vxe-column field="mark" title="备注" :edit-render="{ name: 'ElInput' }" />
        <vxe-column title="操作" width="120">
          <template #default="{ row }">
            <el-button link type="primary" @click="onVariableAppendClick">新增</el-button>
            <el-button link type="danger" @click="onVariableDeleteClick(row)">删除</el-button>
          </template>
        </vxe-column>
      </vxe-table>
    </template>
  </BsDialog>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from "vue";
import { ElMessage, ElMessageBox, FormInstance } from "element-plus";
import {
  List,
  Create,
  Delete,
  Modify,
} from "@/wailsjs/go/controller/BuildProject";
import {
  GetVariableList,
  SaveVariable
} from '@/wailsjs/go/controller/Variable';
import {
  ProgrammeList,
} from "@/wailsjs/go/main/Template";
import { model, repository } from "@/wailsjs/go/models";
import { Ref } from "vue";
// import { useRouter } from "vue-router";

const nameTitle = "项目";
// 标题
const title = ref("项目");

const variableTitle = ref("设置变量");
const variableVisible = ref(false);
const selectData: Ref<model.BuildProject> = ref(model.BuildProject.createFrom());

const tableRef = ref();
// 显示弹窗
const visible = ref(false);
// 操作类型
const operationType = ref(0);
// 数据
const dataFormRef = ref<FormInstance>();
// 路由
// const router = useRouter();

const rules = reactive({
  title: [{ required: true, message: "请输入标题", trigger: "blur" }],
  name: [{ required: true, message: "请输入姓名", trigger: "blur" }],
  version: [{ required: true, message: "请输入版本", trigger: "blur" }],
  build_type: [{ required: true, message: "请选择构建类型", trigger: "blur" }],
  build_version: [{ required: true, message: "请输入构建版本", trigger: "blur" }],
  project_path: [{ required: true, message: "请输入项目路径", trigger: "blur" }],
});

// 分页
const pageNum = ref(1);
const pageSize = ref(10);
const total = ref(0);

// 查询表单
const form = reactive({
  condition: "",
});

// 弹窗表单
const dataForm = reactive<model.BuildProject>(
  model.BuildProject.createFrom({})
);

onMounted(async () => {
  programmeTableData.value = await ProgrammeList();
});

const varTableData: Ref<model.Variable[]> = ref([]);
const tableData: Ref<model.BuildProject[]> = ref([]);
const programmeTableData: Ref<model.Scheme[]> = ref([]);

const getData = async () => {
  let sendData = repository.CreateBuildProject.createFrom(dataForm)
  const data = await List(sendData, pageNum.value, pageSize.value);
  tableData.value = data;
};

// 重置表单数据
const resetForm = () => {
  dataForm.id = "";
  dataForm.title = "";
  dataForm.name = "";
  dataForm.project_path = "";
  dataForm.version = "";
  dataForm.build_type = 0;
  dataForm.build_version = 1;
  dataForm.mark = "";
};

interface CommandInfo {
  type: string;
  data: model.BuildProject;
}

const handleCommand = async (value: CommandInfo) => {

  selectData.value = value.data;
  
  switch (value.type) {
    case "build_code": {
      // router.push({
      //   path: "/data_model",
      //   query: {
      //     id: value.data.id,
      //     title: value.data.title,
      //     // makeType: value.data.makeType,
      //   },
      // });
      break;
    }
    case "injet_variable": {
      variableTitle.value = `[注入变量]`;
      varTableData.value = await GetVariableList(value.data.id);
      variableVisible.value = true;
      break;
    }
    case "env_variable": {
      variableTitle.value = `[环境变量]`;
      varTableData.value = await GetVariableList(value.data.id);
      variableVisible.value = true;
      break;
    }
  }
};

// 赋值表单数据
const setForm = (value: model.BuildProject) => {
  dataForm.id = value.id;
  dataForm.title = value.title;
  dataForm.name = value.name;
  dataForm.project_path = value.project_path;
  dataForm.version = value.version;
  dataForm.build_type = value.build_type;
  dataForm.build_version = value.build_version;
  dataForm.mark = value.mark;
};

onMounted(() => {
  getData();
});

// 事件

/**
 * 添加事件
 */
const onAddClick = () => {
  operationType.value = 0;
  title.value = `[添加]${nameTitle}`;
  resetForm();
  visible.value = true;
};

const onVariableAppendClick = async () => {
  const $table = tableRef.value;
  if ($table) {
    let record = model.Variable.createFrom();
    const { row: newRow } = await $table.insertAt(record, -1);
    await $table.setEditRow(newRow, "name");
  }
}

const onVariableDeleteClick = async (value: model.Variable) => {
  const $table = tableRef.value;
  if ($table) {
    await $table.remove(value);

    const tbdata = $table.getTableData();
    if (tbdata.tableData.length === 0) {
      let record = model.Variable.createFrom();
      const { row: newRow } = await $table.insertAt(record, -1);
      await $table.setEditRow(newRow, "name");
    }
  }
}

const onChange = () => {
  getData();
};

const onQryClick = () => {
  getData();
};

const onEditClick = (value: any) => {
  operationType.value = 1;
  title.value = `[编辑]${nameTitle}`;
  setForm(value);
  visible.value = true;
};

const onDeleteClick = (value: any) => {
  console.log("value", value.value);

  ElMessageBox.confirm("请确认是否要删除数据？", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      await Delete(value.id);
      getData();
    })
    .catch(() => {
      ElMessage.info("取消删除");
    });
};

const onVariableClose = () => {
  selectData.value = model.BuildProject.createFrom();
  variableVisible.value = false;
}

const onVariableSave = async () => {
  const $table = tableRef.value;
  if ($table) {
    const tbdata = $table.getTableData();
    try {
      await SaveVariable(selectData.value.id, tbdata.tableData);
      variableVisible.value = false;
    } catch (error) {
      ElMessage.error(`保存失败: ${error}`);
    }
  }
}

const onVariableOpen = async () => {
  // let  GetVariableList();
  if (varTableData.value.length === 0) {
    const $table = tableRef.value;
    if ($table) {
      let record = model.Variable.createFrom();
      const { row: newRow } = await $table.insertAt(record, -1);
      await $table.setEditRow(newRow, "name");
    }
  }
}

const onSave = async () => {
  if (!dataFormRef.value) return;
  await dataFormRef.value.validate(async (valid, fields) => {
    if (valid) {
      switch (operationType.value) {
        case 0: {
          let send = repository.CreateBuildProject.createFrom(dataForm);
          await Create(send);
          ElMessage.success("创建成功");
          visible.value = false;
          getData();
          break;
        }
        case 1: {
          let send = repository.ModifyBuildProject.createFrom(dataForm);
          await Modify(send);
          ElMessage.success("修改成功");
          visible.value = false;
          getData();
          break;
        }
      }
    } else {
      console.log("表单验证失败", fields);
    }
  });
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
