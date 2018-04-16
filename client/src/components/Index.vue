<template>
  <div class="layout">
    <nav-bar></nav-bar>
        <!--导航-->
    <div class="content">

      <!--面包屑导航-->
      <div class="breadcrumb">
        <mu-breadcrumb>
          <mu-breadcrumb-item href="/">Home</mu-breadcrumb-item>
          <mu-breadcrumb-item href="/">VideoGame</mu-breadcrumb-item>
          <mu-breadcrumb-item>Download</mu-breadcrumb-item>
        </mu-breadcrumb>
      </div>
      <!--面包屑导航-->

      <mu-row class="body-content" gutter>
        <!--notify start-->
        <mu-col width="100" tablet="100" desktop="100">
          <mu-chip class="notify"  @delete="handleClose" showDelete>
            <span>未设置登录密码，请前往 <a href="">修改密码</a> 页面进行设置。设置后将可以在移动设备上使用邮箱登录网站。</span>
          </mu-chip>
        </mu-col>
        <!--notify end-->
        <!--content start-->
         <router-view></router-view>
          <right-sidebar :is_article="true"></right-sidebar>
      </mu-row>
    </div>
    <div class="footer">
      Muse-UI ©2017 Created by Muse-UI
    </div>
  </div>
</template>

<script>
  import NavBar from './Navber'
  import ArticleList from './ArticleList'
  import RightSidebar from './RightSidebar'
export default {
  name: 'home',
  components: {
    ArticleList,
    RightSidebar,
    NavBar
  },
  data () {
    return {
      activeTabA: 't1',
      activeTabR: 'tab4',
    }
  },
  methods: {

    handleClose () {

    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .body-content {
    margin-left: 10px;
  }

  .layout{
    background-color: rgb(236, 236, 236);
  }

  .content{
    width: 90%;
    margin: 0 auto;
    max-width: 1300px;
  }
  .breadcrumb{
    margin: 10px 0;
  }
  .footer{
    padding: 20px 0;
    text-align: center;
  }

</style>
