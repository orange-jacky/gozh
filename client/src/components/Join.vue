<template>
  <div class="join-container">

    <div id="join-bg">
      <nav-bar></nav-bar>
      <div class="join-form" id="join-from">
        <!--头部注册标识-->
        <div class="join-title">
          注册
        </div>
        <!--输入框-->
        <div class="input">
          <mu-text-field type="text" hint-text="昵称"
                         name="nick_name" icon="account_circle"/>
          <br/>
          <mu-text-field type="email" full-width="true" hint-text="请输入邮箱地址" icon="email" name="email"/>
          <br/>
          <mu-text-field type="password" full-width="true" hint-text="请输入密码" icon="remove_red_eye"
                         name="password"/>
          <br/>
          <mu-text-field type="password" full-width="true" hint-text="请再次输入密码" icon="remove_red_eye"
                         name="second_password"/>
          <br/>
          <mu-text-field type="text" hint-text="请输入验证码"
                         name="validate" max-length="4" icon="keyboard"/>
          <img src="http://120.79.195.26/validate_code" height="30px" width="60px"/>
          <br/>

          <mu-text-field type="text" hint-text="请输入邮箱验证码"
                         name="validate" max-length="4" icon="keyboard"/>
          <mu-flat-button label="获取邮箱验证码" class="demo-flat-button"/>
          <br/>


          <!--按钮-->
          <div style="width:200px; margin:0px auto;">
            <mu-raised-button class="login-button" full-width="true" label="注册"/>
          </div>
        </div>
      </div>
      <my-footer></my-footer>
    </div>
  </div>
</template>

<script>
  import NavBar from './Navber'
  import MyFooter from './Footer'

  export default {
    name: "Join",
    components: {
      NavBar,
      MyFooter
    },
    data() {
      return {}
    },
    mounted() {
      bgAndFooterLayout();
    }
  }

  function bgAndFooterLayout() {
    // noinspection UnnecessaryLocalVariableJS
    var height = parseInt(window.innerHeight);
    // noinspection JSValidateTypes
    var panelHeight = parseInt(document.getElementById("join-from").clientHeight) + 162;
    // noinspection JSValidateTypes
    if (panelHeight > height) {
      document.getElementById("join-bg").style.height = panelHeight + 50 + "px";
    } else {
      document.getElementById("join-bg").style.height = height + "px";
    }
  }
</script>
<style scoped>
  body {
    background-color: #e8e8e8;
  }

  #join-bg {
    background-color: #e8e8e8;
    min-height: 100%;
  }

  .join-form {
    width: 400px;
    margin: 100px auto 5%;
    background-color: white;
    border-radius: 10px;
  }

  .join-title {
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
  .demo-flat-button {
    margin: 12px;
  }
</style>
