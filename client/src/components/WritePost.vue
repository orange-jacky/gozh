<template>
  <div>
    <nav-bar></nav-bar>
    <div>
      <div>
        <mu-select-field v-model="value" @change="handleSelectTypeChange">
          <mu-menu-item value="1" title="分类一"/>
          <mu-menu-item value="2" title="分类二"/>
          <mu-menu-item value="3" title="分类三"/>
          <mu-menu-item value="4" title="分类四"/>
          <mu-menu-item value="5" title="分类五"/>
          <mu-menu-item value="6" title="分类六"/>
          <mu-menu-item value="7" title="分类七"/>
        </mu-select-field>
        <mu-auto-complete hintText="请随便输入点啥" @input="handleInput" :dataSource="dataSource" @change="handlechange" />
      </div>

    </div>
  </div>
</template>

<script>
  import NavBar from './Navber'
  import MyFooter from './Footer'
  import RightSidebar from './RightSidebar'

  export default {
    name: "WritePost",
    components: {
      NavBar,
      MyFooter,
      RightSidebar
    },
    data() {
      return {
        "value": "1",
      }
    },
    methods: {
      handleSelectTypeChange (value) {
        this.value = value;
      }
    }
  }
</script>

<style scoped>

</style>
