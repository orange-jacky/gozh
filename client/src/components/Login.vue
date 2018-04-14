<template>
  <!-- 背景 -->
  <div class="login">
    <mu-paper class="demo-paper" :z-depth="1" rounded=true>
      <div class="login-title">
        登录
      </div>
    </mu-paper>

  </div>
</template>

<script>
    export default {
        name: "Login",
        data() {
          return {
          }
        }
    }
</script>

<style scoped>
.login{
  margin: 0 auto;
}
.login-title{
  border-top: 10px;
  border-color: black;
  width: 350px;
}
</style>
