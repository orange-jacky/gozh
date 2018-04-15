<template>
  <div class="layout">
    <mu-appbar title="Title">
      <mu-tabs :value="activeTab" @change="handleTabChange" class="tab">
        <mu-tab value="tab1" title="技术分享"/>
        <mu-tab value="tab2" title="技术翻译"/>
        <mu-tab value="tab3" title="综合文章"/>
      </mu-tabs>
      <mu-text-field icon="search" class="appbar-search-field"  slot="right" hintText="请输入搜索内容"/>
      <mu-flat-button color="white" label="搜索" slot="right"/>
      <mu-icon-menu icon="more_vert" slot="right">
        <mu-menu-item title="菜单 1"/>
        <mu-menu-item title="菜单 2"/>
        <mu-menu-item title="菜单 3"/>
        <mu-menu-item title="菜单 4"/>
        <mu-menu-item title="菜单 5"/>
      </mu-icon-menu>
    </mu-appbar>
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
  import RightSidebar from '../RightSidebar'
  export default {
    name: 'index',
    components: {
      RightSidebar
    },
    data () {
      return {
        activeTabA: 't1',
        activeTab: 'tab1',
        activeTabR: 'tab4',
      }
    },
    methods: {
      handleTabChange (val) {
        this.activeTab = val
      },
      handleClose (){

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
  .tab{
    margin: 0 auto;
    width: 300px;
    background-color: rgba(0, 0, 0, 0);
  }
  .content{
    width: 90%;
    margin: 0 auto;
  }
  .breadcrumb{
    margin: 10px 0;
  }
  .footer{
    padding: 20px 0;
    text-align: center;
  }

</style>
