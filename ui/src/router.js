import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Album from "./views/Album.vue";

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/album/:id',
      name: 'album',
      component: Album
    },
  ]
})
