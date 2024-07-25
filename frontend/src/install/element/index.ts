import ElementPlus from 'element-plus'
import zhCN from "element-plus/dist/locale/zh-cn.mjs" //引入中文
import 'element-plus/dist/index.css'

import * as ElementPlusIconsVue from '@element-plus/icons-vue'

export const elementInstall = (app: any) => {
    app.use(ElementPlus, {locale:zhCN})

    for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
        app.component(key, component)
    }
}