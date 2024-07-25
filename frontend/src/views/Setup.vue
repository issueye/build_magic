<template>
  <BsHeader title="设置" description="一些设置的内容">
    <template #actions>
      <!-- <el-button type="primary" @click="onAddClick">添加</el-button>  -->
    </template>
  </BsHeader>
  <BsMain :usePadding="false" :diffHeight="42">
    <template #body>
      <div class="w-full h-full flex">
        <el-tabs tab-position="left" class="h-full" v-model="activeName">
          <el-tab-pane label="字典设置" name="first">User</el-tab-pane>
        </el-tabs>
      </div>
    </template>
  </BsMain>

  <!-- <BsDialog :title="title" :width="500" :visible="visible" @close="onClose" @save="onSave" @open="onOpen">
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
  </BsDialog> -->
</template>

<script lang="ts" setup>
import { ref } from "vue";

const activeName = ref("first");
</script>

<style scoped>
::v-deep(.el-tabs__nav) {
  width: 200px;
}
</style>