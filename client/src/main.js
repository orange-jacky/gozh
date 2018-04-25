// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
// 引入字体图标
import fontawesome from '@fortawesome/fontawesome'
import '@fortawesome/fontawesome/styles.css';
import FontAwesomeIcon from '@fortawesome/vue-fontawesome'
import solid from '@fortawesome/fontawesome-free-solid'
import regular from '@fortawesome/fontawesome-free-regular'
import brands from '@fortawesome/fontawesome-free-brands'
fontawesome.config = {
  autoAddCss: false,
};
// 引入markdown编辑器及样式
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
Vue.use(mavonEditor)


fontawesome.library.add(solid)
fontawesome.library.add(regular)
fontawesome.library.add(brands)

Vue.component('font-awesome-icon', FontAwesomeIcon)
// 引入muse-UI 组件
import MuseUI from 'muse-ui'
import 'muse-ui/dist/muse-ui.css'
import 'muse-ui/dist/theme-carbon.min.css'
import store from './store/store'
Vue.use(MuseUI)
Vue.config.productionTip = false

Vue.prototype.axios = axios

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
