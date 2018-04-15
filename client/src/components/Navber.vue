<template>
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
</template>

<script>
  export default {
    name: "navber",
    data() {
      return {
        activeTab: 'tab1',
      }
    },
    methods: {
      handleTabChange (val) {
        this.activeTab = val
      },
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
