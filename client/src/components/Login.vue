<template>
  <div class="login-container">
    <div class="login-bg" id="login-bg">
      <nav-bar></nav-bar>
      <div class="login-form" id="login-from">
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
          <mu-icon-button href="#">
            <font-awesome-icon :icon="['fab', 'github-alt']"/>
          </mu-icon-button>
          <mu-icon-button href="#">
            <font-awesome-icon :icon="['fab', 'weixin']"/>
          </mu-icon-button>
          <mu-icon-button href="#">
            <font-awesome-icon :icon="['fab', 'qq']"/>
          </mu-icon-button>
        </div>
      </div>
    </div>
    <my-footer></my-footer>
  </div>
</template>

<script>
  import NavBar from './Navber'
  import MyFooter from './Footer'

  export default {
    name: "Login",
    components: {
      NavBar,
      MyFooter
    },
    data() {
      return {}
    },
    created() {
      bgAndFooterLayout();
    },
    mounted() {
      bgAndFooterLayout();
    }
  }

  function bgAndFooterLayout() {
    // noinspection UnnecessaryLocalVariableJS
    var height = parseInt(window.innerHeight);
    // noinspection JSValidateTypes
    var panelHeight = parseInt(document.getElementById("login-from").clientHeight) + 162;
    // noinspection JSValidateTypes
    if (panelHeight > height) {
      document.getElementById("login-bg").style.height = panelHeight + 50 + "px";
    } else {
      document.getElementById("login-bg").style.height = height + "px";
    }
  }

</script>

<style scoped>
  body {
    background-color: #e8e8e8;
  }

  .login-bg {
    background-color: #e8e8e8;
    min-height: 100%;
  }

  .login-form {
    width: 400px;
    margin: 100px auto 0;
    background-color: white;
    border-radius: 10px;
  }

  .login-title {
    /*border: 2px solid black;*/
    margin: auto 0px;
    font-size: 200%;
    text-align: center;
    height: 57px;
    padding-top: 15px;
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
    border-top: 1px solid #ccc;
  }

  .sign-in-with-other .txt {
    color: #686868;
    vertical-align: middle;
  }

  .sign-in-with-other-botton {
    width: 154px;
    margin: 0 auto;
  }
</style>
