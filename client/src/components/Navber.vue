<template>
  <mu-appbar title="Title">
    <mu-tabs :value="activeTab" @change="handleTabChange" class="tab">
      <mu-tab value="tab1" title="技术分享" href="/"/>
      <mu-tab value="tab2" title="技术翻译"href="/documents" />
      <mu-tab value="tab3" title="综合文章" href="#"/>
      <mu-tab value="tab4" title="中文教程" href="/courses"/>
    </mu-tabs>
    <mu-text-field icon="search" class="appbar-search-field"  slot="right" hintText="请输入搜索内容"/>
    <mu-flat-button color="white" label="搜索" slot="right"/>
    <mu-icon-menu icon="more_vert" slot="right">
      <mu-menu-item @click="userCenter" title="个人中心"/>
      <mu-menu-item @click="publish" title="发布文章"/>
      <mu-menu-item @click="publish" title="提个问题"/>
    </mu-icon-menu>
  </mu-appbar>
</template>

<script>
  export default {
    name: "navber",
    props: ['activeTab'],
    data() {
      return {
        // activeTab: 'tab1',
      }
    },
    methods: {
      handleTabChange(val) {
        this.activeTab = val
      },

      userCenter() {
        this.$router.push('/user/sad')
      },
      publish() {
        this.$router.push('create')
      }
    }
  }

</script>

<style scoped>
  .tab{
    margin: 0 auto;
    width: 300px;
    background-color: rgba(0, 0, 0, 0);
  }
</style>
