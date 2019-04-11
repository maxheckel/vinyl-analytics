import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import dotenv from 'dotenv'
dotenv.config()

Vue.config.productionTip = false

console.log(process.env)
new Vue({
  router,
  store,
  render: function (h) { return h(App) }
}).$mount('#app')
