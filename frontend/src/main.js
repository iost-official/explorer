// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import BootstrapVue from "bootstrap-vue"
import App from './App'
import router from './router'
import store from './store'
import './style.less'

//引入nprogress
import NProgress from 'nprogress'
import 'nprogress/nprogress.css' //这个样式必须引入

// 简单配置
NProgress.inc(0.2)
NProgress.configure({ easing: 'ease', speed: 1000, showSpinner: false, minimum:0.1 })


// 全局前置守卫
router.beforeEach((to,from,next) => {
  NProgress.start()
  next()
})

// 全部后置钩子
router.afterEach(() => {
  NProgress.done()
})


// import './style.less'


Vue.use(BootstrapVue)
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App },
  store
})
