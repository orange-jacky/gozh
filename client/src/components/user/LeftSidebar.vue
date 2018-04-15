<template>
  <mu-col width="100" tablet="30" desktop="30">
    <div v-if="is_edit" class="demo-menu-container">
      <mu-paper class="demo-menu" :zDepth="2" >
        <div class="modify_profile">
          <mu-raised-button class="demo-raised-button" backgroundColor="#DB6862" label="个人资料" icon="view_list" fullWidth default/>
        </div>
        <div class="modify_profile">
          <mu-raised-button class="demo-raised-button" label="修稿头像" icon="account_circle" fullWidth default/>
        </div>
        <div class="modify_profile">
          <mu-raised-button class="demo-raised-button" label="消息通知" icon="notifications" fullWidth default/>
        </div>
        <div class="modify_profile">
          <mu-raised-button class="demo-raised-button" label="修改密码" icon="lock_outline" fullWidth default/>
        </div>
      </mu-paper>
    </div>
    <div v-if="!is_edit" class="demo-menu-container">
      <!--个人简介-->
      <mu-paper class="demo-menu" :zDepth="2" >
        <div class="">
          <mu-row gutter>
            <mu-col width="100" tablet="60" desktop="60">
              <div class="profile-avatar">
                <mu-avatar :size="140" slot="left" src="http://www.muse-ui.org/images/uicon.jpg"/>
              </div>
            </mu-col>
            <mu-col width="100" tablet="40" desktop="40">
              <div class="profile-body">
                <h3>Acorns</h3>
                <p>第<span>1</span>个会员</p>
                <p>注册于: <span> 2年前</span></p>
                <p>活跃于: <span> 5分钟</span></p>
              </div>
            </mu-col>
          </mu-row>
          <mu-row gutter>
            <mu-col width="30" tablet="30" desktop="30">
              <div class="profile-flow">
                <p><a href="javascript:;">8</a></p>
                <span class="right-icon">关注者</span>
              </div>
            </mu-col>
            <mu-col width="30" tablet="30" desktop="30">
              <div class="profile-flow">
                <p><a href="javascript:;">8</a></p>
                <span class="right-icon">关注者</span>
              </div>
            </mu-col>
            <mu-col width="30" tablet="30" desktop="30">
              <div class="profile-flow">
                <p><a href="javascript:;">8</a></p>
                <span class="right-icon">关注者</span>
              </div>
            </mu-col>
          </mu-row>
          <div class="profile">
            <mu-raised-button icon="edit" class="profile-button profile-button-follow" label="编辑资料" fullWidth/>
          </div>
        </div>
      </mu-paper>
      <!--功能区域-->
      <mu-paper class="demo-menu" :zDepth="2" >
        <div class="right-content">
          <mu-raised-button class="right-button" label="发起讨论" icon="record_voice_over" Default/>
          <mu-raised-button class="right-button" label="发布文章" icon="edit" Default/>
          <mu-raised-button class="right-button" label="提个问题" icon="help"  Default/>
          <mu-raised-button class="right-button" label="分享此站" icon="share" Default/>
        </div>
      </mu-paper>
      <!--推荐列表-->
      <mu-paper class="demo-menu" :zDepth="2" >
        <div class="">
          <mu-list>
            <mu-list-item title="我发布的文章"><mu-icon slot="left" value="insert_drive_file"/></mu-list-item>
            <mu-divider/>
            <mu-list-item title="我发布的话题"><mu-icon slot="left" value="library_books"/></mu-list-item>
            <mu-divider/>
            <mu-list-item title="我发布的提问"><mu-icon slot="left" value="help"/></mu-list-item>
            <mu-divider/>
            <mu-list-item title="我发表的回复"><mu-icon slot="left" value="new_releases"/></mu-list-item>
            <mu-divider/>
            <mu-list-item title="我关注的用户"><mu-icon slot="left" value="favorite"/></mu-list-item>
            <mu-divider/>
            <mu-list-item title="我参与的翻译"><mu-icon slot="left" value="translate"/></mu-list-item>
            <mu-divider/>
          </mu-list>
        </div>
      </mu-paper>
    </div>
  </mu-col>
</template>

<script>
    import MuRow from "muse-ui/src/grid/row";
    import MuIcon from "muse-ui/src/icon/icon";

    export default {
      components: {
        MuIcon,
        MuRow},
      name: "left-sidebar",
      props: ['is_edit'],
      activeColor: '#DB6862'
    }
</script>

<style scoped>
  .demo-menu-container{
    max-width: 350px;

  }
  .profile-avatar{
    float: left;
    margin: 15px 30px 5px 20px;
  }
  .profile-body {
    display: table-cell;
    vertical-align: top;
    margin-top: 20px;
    color: #606060;
  }
  .profile-body span {
    color: #ccc;
  }
  .right-title{
    text-align: center;
    border-bottom: 1px solid #cccccc;
    height: 40px;
  }
  .right-list a {
    color: #000000;
  }
  .right-list a:hover {
    color: #DB6862;
  }
  .right-icon {
    font-size: 14px;
    color: #cccccc;
  }
  .right-icon :hover{
    color: #DB6862;
  }
  .right-content{
    text-align: center;
  }

  .right-detail a:hover{
    color: #DB6862;
  }
  .right-detail a{
    color: #737373;
  }
  .right-title span {
    color: #cccccc;
    line-height: 40px;
    margin-left: 10px;
    font-size: 14px;
  }
  .right-button{
    margin: 12px;
  }
  .profile{
    margin: 20px 0 20px 0;
    width: 100%;
    text-align: center;
    padding: 20px 0 15px 0;
  }
  .profile-button{
    margin: 3px 0 5px 0;
    width: 80%;
  }
  .profile-button-follow{
    background:#DB6661 ;
  }
  .demo-menu {
    margin: 16px 32px 16px 0;
  }
  .profile-flow{
    font-size: 20px;
    text-align: center;
  }
  .modify_profile{
    margin: 0 auto;
    text-align: center;
    width: 80%;
    padding: 10px 0 10px 0;
    border-bottom: 1px solid #ccc;
  }
  .modify_profile span:hover {
    color: #DB6862;
  }

</style>
