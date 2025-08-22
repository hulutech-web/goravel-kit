import {createApp} from 'vue';
import App from './App.vue';
import Antd from 'ant-design-vue';
import setupVxe from "@/plugins/vxetable";
import router from '@/router';
import stepin from 'stepin/es';
import pinia from '@/store';
// 生产打包时可去除 ant-design-vue/dist/antd.variable.less 的引用。
// 开发引入此包是为了加载优化，防止首次打开页面过慢
import 'ant-design-vue/dist/antd.variable.less';
import 'stepin/es/style';
// import 'default-passive-events';
import '@/theme/index.less';
import {AuthPlugin, IconfontPlugin} from '@/plugins';
import {setupPicker} from "@/plugins/colorpicker"
import {setupDirective} from "@/directive"
const app = createApp(App);
app.use(Antd)
app.use(pinia);
app.use(router);
app.use(stepin, {router});
app.use(AuthPlugin, {action: 'disable'});
// iconfont 插件。url为你的 iconfont 图标资源地址（你的iconfont 仓库可获取此地址）
app.use(IconfontPlugin, {
    url: ['//at.alicdn.com/t/c/font_3805284_ulvha6ct7d.js', '//at.alicdn.com/t/c/font_4430217_19gqe4agmcp.js'],
});
app.config.errorHandler = function (err) {
    console.error('未捕获的异常，', err);
};
setupVxe(app);
setupPicker(app);
// 自定义指令
setupDirective(app)
app.mount('#stepin-app');
