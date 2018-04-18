<template>
  <div class="layout">
    <nav-bar></nav-bar>
    <div class="container">
      <mu-row gutter>
        <mu-col width="100" tablet="100" desktop="100">
          <mu-paper  :zDepth="2" >
            <div class="title">
              <h1><font-awesome-icon :icon="['far','comment']"/> 创建新的文章/话题</h1>
            </div>
            <div class="description">
              <div>
                <span>选择分类</span>

              </div>
            </div>
          </mu-paper>
        </mu-col>
      </mu-row>
    </div>
  </div>
</template>

<script>
  import NavBar from '../Navber'
  export default {
    name: "create",
    components: {
      NavBar
    }
  }
</script>

<style scoped>
  .container{
    max-width: 1300px;
    width: 90%;
    margin: 0 auto;
  }
  .layout{
    background-color: rgb(236, 236, 236);
  }
  .title {
    margin: 30px;
    padding: 30px 30px 0 30px;
    border-bottom: 1px solid #ccc;
    text-align: center;
  }
  .description {

  }
</style>
