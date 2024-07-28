import { createRouter, createWebHistory } from "vue-router";
import Project from "./views/project/index.vue";
import Root from "./views/Root.vue";
import Template from './views/template/index.vue';
import ScriptEdit from './views/script_edit/index.vue';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: "root", component: Root },
    { path: "/project", name: "project", component: Project },
    { path: "/templates", component: Template },
    { path: "/script_edit", component: ScriptEdit },
  ],
});

export default router;
