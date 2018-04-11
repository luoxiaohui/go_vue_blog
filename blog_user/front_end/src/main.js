import Vue from 'vue'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'
import App from './App.vue'
import axios from 'axios'
import './api'

// 按需引用element
import { Loading, Button, Message, MessageBox, Notification, Popover, Tag, Input } from 'element-ui'
import 'element-ui/lib/theme-default/index.css'
import sidebar from './components/sidebar.vue'
import article from './components/article.vue'
import about from './components/about.vue'
import articleDetail from './components/articleDetail.vue'

Vue.use(VueRouter)
Vue.use(VueResource)


//---------------------------------------------配置axios开始-----------------------------------------
// 超时时间，单位为毫秒
axios.defaults.timeout = 60000
    // axios默认不发送cookie，需要设置此属性为true才可以
axios.defaults.withCredentials = true
    // http请求拦截器
var loadinginstace
axios.interceptors.request.use(config => {
        // element ui Loading方法
        loadinginstace = Loading.service({ fullscreen: true })
        return config
    }, error => {
        loadinginstace.close()
        Message.error({
            message: '加载超时'
        })
        return Promise.reject(error)
    })
    // http响应拦截器
axios.interceptors.response.use(response => { // 响应成功关闭loading
        loadinginstace.close()
            // console.log(response.data)
        var code = JSON.parse(response.data)["code"]
            // 登录失效
            // if (code == "001") {

        //     this.$router.push({ path: '/admin/signin', component: signin })
        // } else {
        //     return response
        // }
        return response
    }, error => {
        loadinginstace.close()
        Message.error({
            message: '加载失败'
        })
        return Promise.reject(error)
    })
    //---------------------------------------------配置axios结束-----------------------------------------
Vue.prototype.$axios = axios


const components = [Button, Message, MessageBox, Notification, Popover, Tag, Input]

components.forEach((item) => {
    Vue.component(item.name, item)
})

const MsgBox = MessageBox
Vue.prototype.$msgbox = MsgBox
Vue.prototype.$alert = MsgBox.alert
Vue.prototype.$confirm = MsgBox.confirm
Vue.prototype.$prompt = MsgBox.prompt
Vue.prototype.$message = Message
Vue.prototype.$notify = Notification

const router = new VueRouter({
    routes: [
        { path: '/', components: { default: article, sidebar: sidebar } },
        { path: '/about', components: { default: about, sidebar: sidebar } },
        { path: '/articleDetail/:articleId', components: { default: articleDetail, sidebar: sidebar } },
        { path: '/article', components: { default: article, sidebar: sidebar } }
    ],
    mode: 'history'
})

new Vue({
    el: '#app',
    router: router,
    render: h => h(App)
})