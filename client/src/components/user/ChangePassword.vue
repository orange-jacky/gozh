<template>
  <mu-col width="100" tablet="70" desktop="70">
    <mu-paper class="article"  :zDepth="2" >
      <div class="edit">
        <div class="edit-title">
          <h1><font-awesome-icon :icon="['fas','lock']"/>请填写你的密码</h1>
          <mu-divider/>
        </div>
        <div>
          <div class="edit-body">
            <label>邮箱：</label>
            <mu-auto-complete disabled hintText="goolge@gmail.com" @input="handleInput" @change="handleChange" />
            <span>设置密码后可使用邮箱登录</span>
            <br/>
            <label>密码：</label>
            <mu-auto-complete hintText="" @input="handleInput" @change="handleChange" /><br/>
            <label>确认：</label>
            <mu-auto-complete  hintText="" @input="handleInput" @change="handleChange" />
          </div>
          <mu-raised-button class="button" label="保存修改" secondary fullWidth/>
        </div>
      </div>
    </mu-paper>
  </mu-col>
</template>

<script>
  export default {
    name: "change_password",
    data () {
      return {

      }
    },
    methods: {
      handleChange () {

      },
      handleInput () {

      }
    }
  }
</script>

<style scoped>
  .edit{
    margin: 15px;
  }
  .edit-title{
    padding-top: 30px;
  }
  .edit-body{
    margin: 20px 0 0 0;
  }
  .avatar img{
    width: 380px;
    height: 380px;
  }
  .button{
    margin: 30px 0 30px 0 ;
  }
</style>
