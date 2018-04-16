<template>
  <mu-col width="100" tablet="70" desktop="70">
    <mu-paper class="article"  :zDepth="2" >
      <div class="profile">
        <h1>编辑你的个人资料</h1>
        <mu-divider/>
        <div>
          <span class="field">用户名：</span>
          <font-awesome-icon icon="user"/>
          <mu-text-field label="用户名" labelFloat/><span class="description">说明</span><br/>
          <span>说  明：</span>
          <font-awesome-icon icon="transgender"/>
          <mu-text-field label="性别" labelFloat/><span class="description">说明</span><br/>
          <label>github Name</label>
          <font-awesome-icon :icon="['fab', 'github']" />
          <mu-text-field label="Github Name" labelFloat/><span class="description">说明</span><br/>

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
    margin: 0 0 0 20px;
  }
  .profile h1{
    padding: 20px 0 0 0;
  }
  .description{
    color: #ccc;
    font-size: 12px;
  }
  .field{
    width: 300px;
  }
</style>
