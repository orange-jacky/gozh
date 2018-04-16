<template>
  <mu-col width="100" tablet="70" desktop="70">
    <mu-paper class="article"  :zDepth="2" >
      <div class="profile">
        <h1>编辑个人资料</h1>
        <mu-divider/>
        <div class="edit-profile">
          <span class="field">用户名：</span>
          <font-awesome-icon icon="user"/>
          <mu-text-field label="Username" labelFloat/><span class="description">如：Acorns</span><br/>
          <span>性  别： </span>
          <font-awesome-icon icon="transgender"/>
          <mu-select-field v-model="gander" label="Gander">
            <mu-menu-item value="1" title="男"/>
            <mu-menu-item value="2" title="女"/>
          </mu-select-field><span class="description">如：男</span><br/>
          <label>Github：</label>
          <font-awesome-icon :icon="['fab', 'github']" />
          <mu-text-field label="Github Name" labelFloat/><span class="description">如：Panda</span><br/>
          <label>Email： </label>
          <font-awesome-icon :icon="['fas','envelope']" />
          <mu-text-field label="Email" labelFloat/><span class="description">如：acorndu@gmail.com</span><br/>
          <label>城  市： </label>
          <font-awesome-icon :icon="['fas','university']" />
          <mu-text-field label="City" labelFloat/><span class="description">如：北京</span><br/>
          <label>姓  名： </label>
          <font-awesome-icon :icon="['fas','universal-access']" />
          <mu-text-field label="Name" labelFloat/><span class="description">如：张三，Taylor</span><br/>
          <label>公   司： </label>
          <font-awesome-icon :icon="['fas','building']" />
          <mu-text-field label="Company" labelFloat/><span class="description">如：Google</span><br/>
          <label>简   介： </label>
          <mu-text-field hintText="Describe" multiLine :rows="3" :rowsMax="6" /><span class="description">个人简介</span><br/>
        </div>
        <div class="profile-button">
          <mu-raised-button class="demo-raised-button" backgroundColor="#DB6862" label="保存修改" icon="save" fullWidth default/>
        </div>
      </div>
    </mu-paper>
  </mu-col>
</template>

<script>
  import FontAwesomeIcon from '@fortawesome/vue-fontawesome';

  export default {
    name: "profile",
    components: {FontAwesomeIcon}
  }
</script>

<style scoped>
  .profile{
    margin: 0 20px 0 20px;
  }
  .profile h1{
    padding: 20px 0 0 0;
  }
  .description{
    color: #ccc;
    font-size: 12px;
  }
  .edit-profile {
    padding: 0 0 20px 0;
  }
  .profile-button{
    padding: 10px 0 20px 0;
  }
</style>
