import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import List from '@/components/ArticleList'
import Show from '@/components/Show'
import Login from  '@/components/Login'
import UserCenter from  '@/components/user/Index'
import MyPublish from  '@/components/user/MyPublish'
import Profile from  '@/components/user/Profile'
import ModifyAvatar from  '@/components/user/ModifyAvatar'
import Join from '@/components/Join'
import ChangePassword from '@/components/user/ChangePassword'
import Create from '@/components/user/Create'
import Courses from '@/components/Courses'
import CourseContents from '@/components/CourseContents'

Vue.use(Router)

export default new Router({
  mode: "history",
  routes: [
    {
      path: '/',
      component: Index,
      redirect: 'list',
      children: [
        { //默认首页
          path: "/list",
          name: 'index',
          component: List
        },
        { // 文章详情
          path: '/list/:user/:article_flag',
          name: 'article',
          component: Show
        },
      ]
    },
    {// 发帖
      path: '/create',
      name: 'create',
      component: Create,
      meta: {
        requiredAuth: true
      }
    },
    { // 用户中心
      path: '/user',
      component: UserCenter,
      children:[
        { // 默认用户中心页面
          path: '/user/:user_id',
          name: 'user',
          component: MyPublish,
          meta: {
            requiredAuth: true
          }
        },
        { // 用户修改个人资料
          path: 'profile/:user_id',
          name: 'user_profile',
          component: Profile,
          meta: {
            requiredAuth: true
          }
        },
        { // 用户修改个人资料
          path: 'change_password/:user_id',
          name: 'change_password',
          component: ChangePassword,
          meta: {
            requiredAuth: true
          }
        },
        { // 用户修改头像
          path: 'edit_avatar/:user_id',
          name: 'user_profile',
          component: ModifyAvatar,
          meta: {
            requiredAuth: true
          }
        },
        { // 我的文章
          path: 'my_article/:user_id',
          name: 'user_article',
          meta: {
            requiredAuth: true
          }
        }
      ]
    },
    { // 登录
      path: '/login',
      name: 'login',
      component: Login
    },

    {// 注册
      path: '/join',
      name: 'join',
      component: Join
    },

    {// 教程页面
      path: '/courses',
      name: 'courses',
      component: Courses
    },
    {// 教程目录
      path: '/courses/:course_id',
      name: 'course_contents',
      component: CourseContents
    },
  ]
})
