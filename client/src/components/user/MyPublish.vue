<template>
  <mu-col width="100" tablet="70" desktop="70">
    <mu-paper class="article"  :zDepth="2" >
      <div class="content">
        this is my publish
      </div>
    </mu-paper>
  </mu-col>
</template>

<script>
    import MuPaper from "muse-ui/src/paper/paper";

    export default {
      components: {MuPaper},
      name: "my-publish"
    }
</script>

<style scoped>
  .content{
    margin: 15px 0 0 0;
  }
</style>
