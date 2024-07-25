import { createRouter, createWebHistory } from "vue-router";
import Home from "./views/home/index.vue";
import Project from "./views/project/index.vue";
import Root from "./views/Root.vue";
import DataSource from './views/data_source/index.vue';
import DataModel from './views/data_model/index.vue';
import MakeCode from './views/make_code/index.vue';
import Template from './views/template/index.vue';
import ScriptEdit from './views/script_edit/index.vue';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: "root", component: Root },
    { path: "/home", name: "home", component: Home },
    { path: "/project", name: "project", component: Project },
    { path: "/data_source", component: DataSource },
    { path: "/data_model", component: DataModel },
    { path: "/make_code", component: MakeCode },
    { path: "/templates", component: Template },
    { path: "/script_edit", component: ScriptEdit },
    { path: "/setup", component: () => import("./views/Setup.vue") },
  ],
});

export default router;
