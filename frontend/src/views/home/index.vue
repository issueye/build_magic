<template>
  <BsHeader title="创建构建计划" description="一个项目的构建计划">
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
            <el-table-column prop="makeType" label="生成类型" width="90" show-overflow-tooltip>
              <template v-slot="{ row }">
                <el-tag effect="plain" :type="row.makeType === 1 ? 'success' : 'danger'">
                  {{ row.makeType === 1 ? "数据源" : "手动生成" }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="schemeId" label="脚本方案" width="100" show-overflow-tooltip>
              <template v-slot="{ row }">
                <el-tag class="mr-1" effect="plain">{{
                  getTPName(row.schemeId)
                }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="mark" label="备注" show-overflow-tooltip />
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
                        <el-dropdown-item :command="{ data: row, type: 'edit_model' }">编辑模型</el-dropdown-item>
                        <el-dropdown-item :command="{ data: row, type: 'make_code' }">生成代码</el-dropdown-item>
                        <el-dropdown-item :command="{ data: row, type: 'variable' }">设置变量</el-dropdown-item>
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
        <el-form-item label="计划标题" prop="title">
          <el-input v-model="dataForm.title" placeholder="请输入计划标题" clearable />
        </el-form-item>
        <el-form-item label="脚本方案">
          <el-select v-model="dataForm.schemeId" placeholder="请选择脚本方案">
            <el-option v-for="(item, index) in programmeTableData" :key="index" :value="item.code"
              :label="item.title" />
          </el-select>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="dataForm.mark" placeholder="请输入备注" type="textarea" :row="2" clearable />
        </el-form-item>
      </el-form>
    </template>
  </BsDialog>

  <!-- <BsDialog :title="variableTitle" :width="800" :visible="variableVisible" @close="onVariableClose"
    @save="onVariableSave" @open="onVariableOpen">
    <template #body>
      <vxe-table border show-overflow keep-source size="mini" ref="tableRef" :data="varTableData"
        max-height="550" empty-text="没有数据"
        :edit-config="{ trigger: 'click', mode: 'row', showStatus: true }">
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
  </BsDialog> -->
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import {
  List,
  Create,
  Delete,
  Modify,
  // GetVariableList,
  // SaveVariable
} from "@/wailsjs/go/main/DataModel";
import {
  ProgrammeList,
} from "@/wailsjs/go/main/Template";
import { model, repository } from "@/wailsjs/go/models";
import { Ref } from "vue";
import { useRouter } from "vue-router";

const nameTitle = "构建计划";
// 标题
const title = ref("构建计划");

// const variableTitle = ref("设置变量");
// const variableVisible = ref(false);
// const selectDM:Ref<model.DataModel> = ref(model.DataModel.createFrom());

// const tableRef = ref();
// 显示弹窗
const visible = ref(false);
// 操作类型
const operationType = ref(0);
// 数据
const dataFormRef = ref();
// 路由
const router = useRouter();

const rules = reactive({
  name: [{ required: true, message: "请输入姓名", trigger: "blur" }],
  account: [{ required: true, message: "请输入账户", trigger: "blur" }],
  groupId: [{ required: true, message: "请选择用户组", trigger: "blur" }],
  tpIds: [{ required: true, message: "请选择脚本模板", trigger: "blur" }],
  tableName: [{ required: true, message: "请输入表名", trigger: "blur" }],
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
const dataForm = reactive<model.DataModel>(
  model.DataModel.createFrom({
    name: "",
    title: "",
    project: "",
    makeType: 0,
    schemeId: "",
    mark: "",
  })
);

onMounted(async () => {
  programmeTableData.value = await ProgrammeList();
});

// const varTableData:Ref<model.Variable[]> = ref([]);
const tableData: Ref<model.DataModel[]> = ref([]);
const programmeTableData: Ref<model.Scheme[]> = ref([]);

const getData = async () => {
  const data = await List(form.condition, pageNum.value, pageSize.value);
  tableData.value = data;
};

// 重置表单数据
const resetForm = () => {
  dataForm.id = "";
  dataForm.title = "";
  dataForm.makeType = 0;
  dataForm.tableName = "";
  dataForm.project = "";
  dataForm.schemeId = "";
  dataForm.mark = "";
};

interface CommandInfo {
  type: string;
  data: model.DataModel;
}

const handleCommand = async (value: CommandInfo) => {
  switch (value.type) {
    case "edit_model": {
      router.push({
        path: "/data_model",
        query: {
          id: value.data.id,
          title: value.data.title,
          makeType: value.data.makeType,
        },
      });
      break;
    }
    case "make_code": {
      router.push({
        path: "/make_code",
        query: {
          id: value.data.id,
          title: value.data.title,
          schemeId: value.data.schemeId,
        },
      });
      break;
    }
    case "variable": {
      // variableTitle.value = `[${value.data.title}]变量`;
      // varTableData.value = await GetVariableList(value.data.id);
      // selectDM.value = value.data;
      // variableVisible.value = true;
      break;
    }
  }
};

// 赋值表单数据
const setForm = (value: model.DataModel) => {
  dataForm.id = value.id;
  dataForm.title = value.title;
  dataForm.tableName = value.tableName;
  dataForm.makeType = value.makeType;
  dataForm.schemeId = value.schemeId;
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

// const onVariableAppendClick = async () => {
//   const $table = tableRef.value;
//   if ($table) {
//     let record = model.Variable.createFrom();
//     const { row: newRow } = await $table.insertAt(record, -1);
//     await $table.setEditRow(newRow, "name");
//   }
// }

// const onVariableDeleteClick = async (value: model.Variable) => {
//   const $table = tableRef.value;
//   if ($table) {
//     await $table.remove(value);

//     const tbdata = $table.getTableData();
//     if (tbdata.tableData.length === 0) {
//       let record = model.Variable.createFrom();
//       const { row: newRow } = await $table.insertAt(record, -1);
//       await $table.setEditRow(newRow, "name");
//     }
//   }
// }

const onChange = () => {
  getData();
};

const onQryClick = () => {
  getData();
};

// 获取脚本模板名称
const getTPName = (value: string): string => {
  return (
    programmeTableData.value.find((item: model.Scheme) => {
      console.log("find->", item, item.code, value);
      return item.code == value;
    })?.title || ""
  );
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

// const onVariableClose = () => {
//   selectDM.value = model.DataModel.createFrom();
//   variableVisible.value = false;
// }

// const onVariableSave = async () => {
//   const $table = tableRef.value;
//   if ($table) {
//     const tbdata = $table.getTableData();
//     try {
//       await SaveVariable(selectDM.value.id, tbdata.tableData);
//       variableVisible.value = false;
//     } catch (error) {
//       ElMessage.error(`保存失败: ${error}`);
//     }
//   }
// }

// const onVariableOpen = async () => {
//   // let  GetVariableList();
//   if (varTableData.value.length === 0) {
//     const $table = tableRef.value;
//     if ($table) {
//       let record = model.Variable.createFrom();
//       const { row: newRow } = await $table.insertAt(record, -1);
//       await $table.setEditRow(newRow, "name");
//     }
//   }
// }

const onSave = () => {
  if (!dataFormRef.value) return;
  dataFormRef.value.validate(async (valid: boolean) => {
    if (valid) {
      switch (operationType.value) {
        case 0: {
          let send = repository.CreateDataModel.createFrom(dataForm);
          await Create(send);
          ElMessage.success("创建成功");
          visible.value = false;
          getData();
          break;
        }
        case 1: {
          let send = repository.RequestModifyDataModel.createFrom(dataForm);
          await Modify(send);
          ElMessage.success("修改成功");
          visible.value = false;
          getData();
          break;
        }
      }
    } else {
      console.log("表单验证失败");
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
