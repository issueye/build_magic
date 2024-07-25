<template>
  <BsHeader :title="mdTitle" description="数据模型信息">
    <template #actions>
      <el-button @click="onBackClick">返回</el-button>
    </template>
  </BsHeader>
  <BsMain>
    <template #body>
      <div class="flex justify-between mb-3">
        <div class="w-[500px] flex items-center gap-2" v-if="makeType == '1'">
          <el-select placeholder="请选择数据源" v-model="selectDataSource" clearable>
            <el-option v-for="(item, index) in dsData" :key="index" :value="item.id" :label="item.title" />
          </el-select>
          <el-badge is-dot :type="linkState">
            <el-button type="primary" @click="onLinkClick">连接</el-button>
          </el-badge>

          <el-select class="ml-20" placeholder="请选择数据表" v-model="selectDataTable" clearable
            @change="onSelectTableChange">
            <el-option v-for="(item, index) in tableNameData" :key="index" :value="item.table_name"
              :label="item.table_name" />
          </el-select>
        </div>
        <div>
          <el-button type="primary" @click="onAddOneCellClick">新增一行</el-button>
          <el-button type="primary" @click="onSaveClick">保存</el-button>
        </div>
      </div>

      <div class="h-[calc(100% - 60px)]">
        <vxe-table border show-overflow keep-source size="mini" ref="tableRef" :data="tableData"
          :menu-config="menuConfig" :edit-rules="validRules" max-height="550" empty-text="没有数据"
          :edit-config="{ trigger: 'click', mode: 'row', showStatus: true }" @menu-click="onCtxMenuClickEvent">
          <vxe-column type="seq" title="序号" width="70" />
          <vxe-column field="title" title="标题" :edit-render="{ name: 'ElInput' }" />
          <vxe-column field="name" title="名称" :edit-render="{ name: 'ElInput' }" />
          <vxe-column field="mark" title="备注" :edit-render="{ name: 'ElInput' }" />
          <vxe-column field="columnType" title="数据类型" :edit-render="columnTypeRender" width="180" />
          <vxe-column field="size" title="长度" width="100"
            :edit-render="{ name: 'ElInputNumber', props: { controlsPosition: 'right' } }" />
          <vxe-column field="isPk" title="主键" width="80"
            :cell-render="{ name: 'ElSwitch', props: { activeValue: 1, inactiveValue: 0 } }" />
          <vxe-column field="extension" title="拓展" width="190" :edit-render="extensionRender" />
        </vxe-table>
      </div>
    </template>
  </BsMain>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import {
  GetModelInfo,
  SaveModelInfo,
} from "@/wailsjs/go/main/DataModel";
import { List as getDataSourceList, LinkDataSource, GetTableList, GetColumns } from "@/wailsjs/go/main/DataSource";
import { model, main } from "@/wailsjs/go/models";
import { Ref } from "vue";
import { ElMessage } from "element-plus";
import { VxeTableEvents } from "vxe-table/types/all";
// import { VxeColumnPropTypes, VxeTableEvents } from "vxe-table";

// 数据
const tableRef = ref();

const selectDataSource = ref();
const selectDataTable = ref();
const linkState = ref('danger');

const router = useRouter();
const route = useRoute();

const menuConfig = reactive({
  body: {
    options: [
      [
        { code: "insertAt", name: "插入", disabled: false },
        { code: "remove", name: "删除", disabled: false },
      ],
    ],
  },
});

const columnTypeRender = reactive({
  name: "ElSelect",
  props: {
    placeholder: "请选择数据类型",
  },
  options: [
    { value: "int", label: "int" },
    { value: "nvarchar", label: "nvarchar" },
    { value: "bool", label: "bool" },
    { value: "text", label: "text" },
    { value: "datetime", label: "datetime" },
    { value: "decimal", label: "decimal" },
  ],
});

const extensionRender = reactive({
  name: "ElSelect",
  props: {
    multiple: true,
    collapseTags: true,
    placeholder: '拓展信息',
  },
  options: [
    { value: "isLike", label: "模糊查询" },
    { value: "isSearch", label: "精确查询" },
    { value: "isOrder", label: "排序" },
    { value: "isShow", label: "显示" },
    { value: "isEdit", label: "编辑" },
  ],
});

const id = ref(route.query["id"]);
const mdTitle = ref(route.query["title"]);
const makeType = ref(route.query['makeType'] as string);

const validRules = ref({
  title: [{ required: true, message: "标题不能为空" }],
  name: [{ required: true, message: "名称不能为空" }],
  columnType: [{ required: true, message: "数据类型不能为空" }],
});

onMounted(() => { });

const tableData: Ref<model.ModelInfo[]> = ref([]);
const dsData: Ref<model.DataSource[]> = ref([]);
const tableNameData: Ref<main.TablelInfo[]> = ref([]);

const getData = async () => {
  const data = await GetModelInfo(id.value as string);
  console.log("data", data);
  tableData.value = data;
};

const onLinkClick = () => {
  linkDataSource();
}

const linkDataSource = async () => {
  if (selectDataSource.value) {
    const res = await LinkDataSource(selectDataSource.value);
    console.log('res', res);
    if (res == null) {
      linkState.value = 'success';
      const list = await GetTableList(selectDataSource.value);
      console.log('tableNameData', list);
      tableNameData.value = list;
    }
  }
};

const getDsData = async () => {
  const data = await getDataSourceList("", 0, 0);
  dsData.value = data;
}

onMounted(() => {
  getData();
  getDsData();
});

const onSelectTableChange = async () => {
  if (selectDataTable.value) {
    const list = await GetColumns(selectDataSource.value, selectDataTable.value);
    if (list) {
      tableData.value = [];
      list.forEach(e => {
        tableData.value.push(
          model.ModelInfo.createFrom({
            title: e.name,
            name: e.name,
            columnType: e.type,
            size: e.size,
            extension: ['isShow'],
          })
        )
      })
    }
  }
}

const onBackClick = () => {
  router.push("/");
};

const onSaveClick = async () => {
  await validEvent();
};

const validEvent = async () => {
  const $table = tableRef.value;
  if ($table) {
    const errMap = await $table.validate();
    if (!errMap) {
      const tbdata = $table.getTableData();
      console.log("$table", tbdata.tableData);
      tableData.value = tbdata.tableData;
      await SaveModelInfo(id.value as string, tableData.value);
      getData();
      ElMessage.success("保存成功");
    }
  }
};

const onAddOneCellClick = async () => {
  const $table = tableRef.value;
  if ($table) {
    const record = model.ModelInfo.createFrom({
      title: "",
      name: "",
      columnType: "",
      length: 0,
      extension: ['isShow'],
    });
    const { row: newRow } = await $table.insertAt(record, -1);
    await $table.setEditRow(newRow, "name");
  }
};

const onCtxMenuClickEvent: VxeTableEvents.MenuClick<model.ModelInfo> = ({
  menu,
  row,
  column,
}) => {
  const $table = tableRef.value;
  if ($table) {
    switch (menu.code) {
      case "insertAt": {
        $table.insertAt({}, row || -1).then(({ row }) => {
          $table.setEditCell(row, column || "name");
        });
        break;
      }
      case "remove":
        $table.remove(row);
        break;
    }
  }
};
</script>