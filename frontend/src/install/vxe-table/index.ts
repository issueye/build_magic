// import { VxeUI, VxeIcon, VxeLoading } from 'vxe-pc-ui'
import { VxeTable, VxeColumn, VxeUI } from 'vxe-table';

import VxeUIPluginRenderElement from '@vxe-ui/plugin-render-element'

// 导入主题变量，也可以重写主题变量
import 'vxe-table/styles/cssvar.scss'
import '@vxe-ui/plugin-render-element/dist/style.css'
// import 'vxe-pc-ui/styles/cssvar.scss'



export const vxeTableInstall = (app: any) => {
    // 添加兼容 element plus 的插件
    VxeUI.use(VxeUIPluginRenderElement)

    app.use(VxeTable)
    app.use(VxeColumn)
}