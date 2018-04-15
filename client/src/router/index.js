import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import List from '@/components/ArticleList'
import Show from '@/components/Show'
import Login from  '@/components/Login'

Vue.use(Router)

export default new Router({
  mode: "history",
  routes: [
    {
      path: '/',
      component: Index,
      redirect: 'list',
      children: [
        {
          path: "/list",
          name: 'index',
          component: List
        },
        {
          path: '/list/:user/:article_flag',
          name: 'article',
          component: Show
        },

      ]
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    }
  ]
})
