<template>
  <mu-col width="100" tablet="70" desktop="70">
    <mu-paper class="article"  :zDepth="2" >
      <div>
        <font-awesome-icon :icon="image"/>
        <h1>请选择头像图片</h1>
      </div>
    </mu-paper>
  </mu-col>
</template>

<script>
    export default {
        name: "modify-avatar"
    }
</script>

<style scoped>

</style>
