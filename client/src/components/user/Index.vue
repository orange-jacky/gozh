<template>
  <div class="layout">
    <nav-bar :active-tab="activeTab"></nav-bar>
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
          <mu-chip class="notify" @delete="handleClose" showDelete>
            <span>未设置登录密码，请前往 <a href="">修改密码</a> 页面进行设置。设置后将可以在移动设备上使用邮箱登录网站。</span>
          </mu-chip>
        </mu-col>
        <!--notify end-->
        <!--content start-->
        <left-sidebar :is_edit="isEdit"></left-sidebar>
        <router-view></router-view>
      </mu-row>
    </div>
    <div class="footer">
      <my-footer></my-footer>
      Gozh ©2018 Created by Muse-UI and VueJS
    </div>
  </div>
</template>

<script>
  import MyFooter from '../Footer'
  import NavBar from '../Navber'
  import LeftSidebar from './LeftSidebar'

  export default {
    name: 'index',
    components: {
      LeftSidebar,
      NavBar,
      MyFooter
    },
    data() {
      return {
        activeTab: 'tab1',
        activeTabA: 't1',
        activeTabR: 'tab4',
      }
    },
    computed: {
      isEdit() {
        return this.$store.getters.isEditP
      }
    },
    methods: {
      handleTabChange(val) {
        this.activeTab = val
      },
      handleClose() {

      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .body-content {
    margin-left: 10px;
  }
  .layout {
    background-color: rgb(236, 236, 236);
  }
  .content {
    max-width: 1300px;
  }

  .tab {
    margin: 0 auto;
    width: 300px;
    background-color: rgba(0, 0, 0, 0);
  }

  .content {
    width: 90%;
    margin: 0 auto;
  }

  .breadcrumb {
    margin: 10px 0;
  }

  .footer {
    padding: 20px 0;
    text-align: center;
  }

</style>
