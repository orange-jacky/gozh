import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import List from '@/components/ArticleList'
import Show from '@/components/Show'
import Login from  '@/components/Login'
import UserCenter from  '@/components/user/Index'
import MyPublish from  '@/components/user/MyPublish'
import Profile from  '@/components/user/Profile'
import WritePost from '@/components/WritePost'

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
      path: '/user',
      component: UserCenter,
      children:[
        {
          path: '/user/:user_id',
          name: 'user',
          component: MyPublish,
          meta: {
            requiredAuth: true
          }
        },
        {
          path: 'profile/:user_id',
          name: 'user_profile',
          component: Profile,
          meta: {
            requiredAuth: true
          }
        },
        {
          path: 'my_article/user_id',
          name: 'user_article',
          meta: {
            requiredAuth: true
          }
        },
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/write_post',
      name: 'write_post',
      component: WritePost,
    }
  ]
})
