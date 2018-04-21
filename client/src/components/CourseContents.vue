<template>
  <div>
    <nav-bar></nav-bar>
    <my-footer></my-footer>
  </div>
</template>

<script>
  import MyFooter from './Footer'
  import NavBar from './Navber'

  export default {
    name: "CourseContents",
    components: {
      MyFooter,
      NavBar
    },
  }
</script>

<style scoped>

</style>
