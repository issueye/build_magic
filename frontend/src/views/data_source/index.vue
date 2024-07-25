<template>
  <BsHeader title="数据源" description="数据库源">
    <template #actions>
      <el-button type="primary" @click="onAddClick">添加</el-button>
    </template>
  </BsHeader>
  <BsMain>
    <template #body>
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
            <el-table-column prop="title" label="标题" width="200" show-overflow-tooltip />
            <el-table-column prop="db_type" label="数据类型" width="130" show-overflow-tooltip />
            <el-table-column prop="mark" label="备注" show-overflow-tooltip />
            <el-table-column label="操作" width="160" align="center" fixed="right">
              <template v-slot="{ row }">
                <el-button type="primary" link size="small" @click="onEditClick(row)">编辑</el-button>
                <el-button type="danger" link size="small" @click="onDeleteClick(row)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <div style="margin-top: 10px">
          <el-pagination small background :current-page="pageNum" :page-size="pageSize" :page-sizes="[5, 10, 20]"
            :total="total" layout="total, sizes, prev, pager, next" @size-change="onSizeChange"
            @current-change="onCurrentChange" />
        </div>
    </template>
  </BsMain>

  <BsDialog :title="title" :width="500" :visible="visible" @close="onClose" @save="onSave">
    <template #body>
      <el-form label-width="auto" :model="dataForm" :rules="rules" ref="dataFormRef">
        <el-form-item label="标题" prop="title">
          <el-input v-model="dataForm.title" placeholder="请输入数据模型标题" clearable />
        </el-form-item>
        <el-form-item label="数据库类型" prop="db_type">
          <el-select v-model="dataForm.db_type" placeholder="请选择数据库类型" clearable @change="onDBTypeChange"
            :disabled="operationType == 1">
            <el-option value="sqlserver" label="sqlserver" />
            <el-option value="mysql" label="mysql" />
            <el-option value="postgresql" label="postgresql" />
            <el-option value="sqlite3" label="sqlite3" />
          </el-select>
        </el-form-item>
        <el-row v-if="dataForm.db_type !== 'sqlite3'">
          <el-col :span="14">
            <el-form-item label="服务地址" prop="host">
              <el-input v-model="dataForm.host" placeholder="请输入服务地址" clearable />
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label="端口号" prop="port">
              <el-input v-model.number="dataForm.port" placeholder="请输入服务地址" clearable />
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="数据库" prop="database" v-if="dataForm.db_type !== 'sqlite3'">
          <el-input v-model="dataForm.database" placeholder="请输入数据库" clearable />
        </el-form-item>
        <el-form-item label="模式" prop="schema" v-if="dataForm.db_type === 'postgresql'">
          <el-input v-model="dataForm.schema" placeholder="请输入模式" clearable />
        </el-form-item>
        <el-form-item label="用户名" prop="userName" v-if="dataForm.db_type !== 'sqlite3'">
          <el-input v-model="dataForm.user_name" placeholder="请输入用户名" clearable />
        </el-form-item>
        <el-form-item label="密码" prop="password" v-if="dataForm.db_type !== 'sqlite3'">
          <el-input v-model="dataForm.password" placeholder="请输入密码" clearable />
        </el-form-item>
        <el-form-item label="数据库路径" prop="path" v-if="dataForm.db_type === 'sqlite3'">
          <el-input v-model="dataForm.path" placeholder="请输入数据库路径" clearable />
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
import {
  List,
  Create,
  Delete,
  Modify,
} from "@/wailsjs/go/main/DataSource";
import { model } from "@/wailsjs/go/models";
import { Ref } from "vue";

const nameTitle = "数据源";
// 标题
const title = ref("数据源");
// 显示弹窗
const visible = ref(false);
// 操作类型
const operationType = ref(0);
// 数据
const dataFormRef = ref();

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
const dataForm = reactive<model.DataSource>(
  model.DataSource.createFrom({
    title: "", // 标题
    host: "", // 数据库服务地址
    port: 0, // 端口号
    username: "", // 用户
    password: "", // 密码
    database: "", // 数据库
    schema: "", // 数据库模式
    db_type: "", // 数据库类型
    path: "", // 数据库路径
    mark: "",
  })
);

const tableData: Ref<model.DataSource[]> = ref([]);

const getData = async () => {
  const data = await List(form.condition, pageNum.value, pageSize.value);
  tableData.value = data;
};

// 重置表单数据
const resetForm = () => {
  dataForm.id = "";
  dataForm.title = "";
  dataForm.host = "";
  dataForm.port = 0;
  dataForm.database = "";
  dataForm.user_name = "";
  dataForm.password = "";
  dataForm.schema = "";
  dataForm.path = "";
  dataForm.db_type = "";
  dataForm.mark = "";
};

// 赋值表单数据
const setForm = (value: model.DataSource) => {
  dataForm.id = value.id;
  dataForm.db_type = value.db_type;
  dataForm.title = value.title;
  dataForm.host = value.host;
  dataForm.port = value.port;
  dataForm.database = value.database;
  dataForm.user_name = value.user_name;
  dataForm.password = value.password;
  dataForm.schema = value.schema;
  dataForm.path = value.path;
  dataForm.mark = value.mark;
};

onMounted(() => {
  getData();
});

// 事件

const onDBTypeChange = (value: any) => {
  console.log("dataForm.db_type", dataForm.db_type, value);
}

/**
 * 添加事件
 */
const onAddClick = () => {
  operationType.value = 0;
  title.value = `[添加]${nameTitle}`;
  resetForm();
  visible.value = true;
};

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

const onSave = () => {
  if (!dataFormRef.value) return;
  dataFormRef.value.validate(async (valid: boolean) => {
    if (valid) {
      console.log('dataForm', dataForm);

      switch (operationType.value) {
        case 0: {
          await Create(dataForm);
          ElMessage.success("创建成功");
          visible.value = false;
          getData();
          break;
        }
        case 1: {
          await Modify(dataForm);
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