/// <reference types="vite/client" />
declare module "*.vue" {
  import { ComponentOptions } from "vue";
  const component: ComponentOptions;
  export default component;
}

declare module 'element-plus/dist/locale/zh-cn.mjs';
declare module '@kangc/v-md-editor/lib/preview';
declare module '@kangc/v-md-editor/lib/theme/vuepress.js';
declare module '@kangc/v-md-editor/lib/plugins/copy-code/index';
declare module '@kangc/v-md-editor/lib/plugins/line-number/index';
declare module '@kangc/v-md-editor/lib/plugins/emoji/index';
declare module '@kangc/v-md-editor/lib/plugins/emoji/index';
