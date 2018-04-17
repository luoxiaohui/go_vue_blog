import App from './App.vue'
import './api'

// 按需引用element
import { Loading } from 'element-ui'

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
    // ---------------------------------------------配置axios结束-----------------------------------------
Vue.prototype.$axios = axios

const article = r => require.ensure([], () => r(require('./components/article')), 'group-home')
const sidebar = r => require.ensure([], () => r(require('./components/sidebar')), 'group-home')
const about = r => require.ensure([], () => r(require('./components/about')), 'group-home')
const articleDetail = r => require.ensure([], () => r(require('./components/articleDetail')), 'group-home')

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