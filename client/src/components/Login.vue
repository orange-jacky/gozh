<template>
  <div class="login-bg">
    <nav-bar></nav-bar>
    <div class="login">
      <!--头部登录标识-->
      <div class="login-title">
        登录
      </div>
      <!--输入框-->
      <div class="input">
        <mu-text-field label="邮箱" type="email" full-width="true" hint-text="请输入邮箱地址" icon="email" name="email"/>
        <br/>
        <mu-text-field label="密码" type="password" full-width="true" hint-text="请输入密码" icon="remove_red_eye"
                       name="password"/>
        <br/>
        <!--按钮-->
        <div style="width:200px; margin:0px auto;">
          <mu-raised-button class="login-button" full-width="true" label="开始吧!"/>
        </div>
      </div>
      <!--第三方登录   横线、文字处-->
      <div class="sign-in-with-other">
        <span style="white-space:pre;">  </span><span class="line"></span>
        <span style="white-space:pre;">  </span><span class="txt">第三方登录</span>
        <span style="white-space:pre;">  </span><span class="line"></span>
      </div>
      <!--第三方登录按钮-->
      <div class="sign-in-with-other-botton">
        <div class="sign-in-with-other-icon">
          <mu-icon value="home"/>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
  import NavBar from './Navber'

  export default {
    name: "Login",
    components: {
      NavBar
    },
    data() {
      return {}
    }
  }
</script>

<style scoped>
  .login-bg {
    background-color: #e8e8e8;
    height: 1000px;
  }

  .login {
    width: 30%;
    height: 40%;
    margin: 100px auto 0;
    background-color: white;
    border-radius: 5px;
  }

  .login-title {
    border: 2px solid black;
    margin: auto 0;
    font-size: 140%;
  }

  .input {
    width: 80%;
    margin: 0 auto;
  }

  .input .login-button {
    /*width:*/
    margin: 0px auto;
  }

  .sign-in-with-other {
    height: 60px;
    line-height: 60px;
    text-align: center;
  }
  .sign-in-with-other .line {
    display: inline-block;
    width: 150px;
    border-top: 1px solid #ccc ;
  }
  .sign-in-with-other .txt {
    color: #686868;
    vertical-align: middle;
  }

  .sign-in-with-other-botton {
    width: 100%;
    position: absolute;
    margin: 0 auto;
  }
</style>
