import { defineStore } from "pinia";
import { Ref, ref } from "vue";

import {
    List,
    GetTree
} from "@/wailsjs/go/controller/Template";
import { model, repository } from "@/wailsjs/go/models";


export const useScriptTemplateStore = defineStore(
    "script_template",
    () => {
        const selectNode = ref('');
        const tableData: Ref<model.CodeTemplate[]> = ref([]);
        const treeData: Ref<repository.SchemeTree[]> = ref([]);

        const pageNum = ref(1);
        const pageSize = ref(10);
        const total = ref(0);

        const getData = async (condition?: string) => {
            const send = repository.QryTemplate.createFrom({
                condition: condition,
                parentCode: selectNode.value,
            });
            const data = await List(send, 0, 0);
            tableData.value = data;
        }

        const getTree = async () => {
            const data = await GetTree();
            treeData.value = data;
        }

        return {
            selectNode,
            tableData,
            pageNum,
            pageSize,
            total,
            treeData,

            getData,
            getTree,
        }
    },
    {
        persist: {
            key: 'script_template',
            storage: localStorage, // or sessionStorage
        }
    }
);