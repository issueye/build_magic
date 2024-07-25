import { iconifyInstall } from './iconify'
import { vxeTableInstall } from './vxe-table'
import { elementInstall } from './element'
import { customComponentsInstall } from './custom_components'
import { piniaInstall } from './pinia'


// 安装
export const install = (app: any) => {
    // element 
    elementInstall(app)

    // 图标库
    iconifyInstall()

    // vxe-table
    vxeTableInstall(app)

    // 自定义组件
    customComponentsInstall(app)

    // pinia
    piniaInstall(app)
}