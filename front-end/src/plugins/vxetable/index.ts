import VxeUIAll, { VxeUI } from 'vxe-pc-ui'
import 'vxe-pc-ui/lib/style.css'
import VxeUITable from 'vxe-table'
import 'vxe-table/lib/style.css'
import VxeUIPluginRenderAntd from '@vxe-ui/plugin-render-antd'
import '@vxe-ui/plugin-render-antd/dist/style.css'
VxeUI.use(VxeUIPluginRenderAntd, {
    // prefixCls: 'ant'
})
const setupVxe = (app) => {
    app.use(VxeUIAll)
    app.use(VxeUITable)
}
export default setupVxe;