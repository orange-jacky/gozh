<template>
  <mu-col width="100" tablet="30" desktop="30">
    <div class="demo-menu-container">
      <!--个人简介-->
      <mu-paper v-if="is_article" class="demo-menu" :zDepth="2" >
        <div class="right-content">
          <div class="right-title">
          <span>作者：Acorns</span>
          </div>
          <div class="profile">
            <mu-avatar :size="80" src="http://www.muse-ui.org/images/uicon.jpg"/>
          </div>
          <div class="profile"><span>男</span></div>
          <div class="profile"><span class="btn btn-success">身份</span></div>
          <div class="profile"><span>A Life-long learner.</span></div>
          <div class="profile">
            <mu-chip class="demo-chip">
              <mu-icon value="location_on" :size="18"/>北京
            </mu-chip>
          </div>
          <div class="profile">
            <mu-raised-button icon="add" class="profile-button profile-button-follow" label="他关注" fullWidth/>
            <mu-raised-button icon="email" class="profile-button" label="发私信" fullWidth/>
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
        <div class="right-content">
          <mu-tabs :value="activeTabR" @change="handleTabChangeR">
            <mu-tab value="tab4" title="最新"/>
            <mu-tab value="tab5" title="最热"/>
            <mu-tab value="tab6"  title="更多"/>
          </mu-tabs>
          <div class="right-list">
            <mu-list-item title="" disabled>
              <mu-avatar  slot="left" src="http://www.muse-ui.org/images/uicon.jpg" :size="30"/>
              <a href=""><p>我试着轻轻抓住光影的踪迹，它却在眉宇间投下一片淡淡的阴影...</p></a>
              <span class="right-icon">5小时前 </span>
              <span class="right-icon"><mu-icon value="favorite" :size="18"/>2</span>
              <span class="right-icon"><mu-icon value="insert_comment" :size="18"/>23</span>
            </mu-list-item>
            <mu-list-item title="" disabled>
              <mu-avatar  slot="left" src="http://www.muse-ui.org/images/uicon.jpg" :size="30"/>
              <a href=""><p>我试着轻轻抓住光影的踪迹，它却在眉宇间投下一片淡淡的阴影...</p></a>
              <span class="right-icon">5小时前 </span>
              <span class="right-icon"><mu-icon value="favorite" :size="18"/>2</span>
              <span class="right-icon"><mu-icon value="insert_comment" :size="18"/>23</span>
            </mu-list-item>
          </div>
        </div>
      </mu-paper>
      <!--推荐招聘-->
      <mu-paper class="demo-menu" :zDepth="2" >
        <div class="right-content">
          <div class="right-title">
            <span>推荐招聘</span>
          </div>
          <div class="right-detail">
            <p><a href="">[北京][20K][中国航空]招聘2年以上go语言工程师</a></p>
            <p><a href="">[北京][20K][中国航空]招聘2年以上go语言工程师</a></p>
            <p><a href="">[北京][20K][中国航空]招聘2年以上go语言工程师</a></p>
            <p><a href="">[北京][20K][中国航空]招聘2年以上go语言工程师</a></p>
          </div>
        </div>
      </mu-paper>
    </div>
  </mu-col>
</template>

<script>
    export default {
      name: "right-sidebar",
      props:['is_article'],
      data () {
          return {
            activeTabR: 'tab4',
          }
      },
      methods: {
        handleTabChangeR (val) {
          this.activeTabR = val
        },
      }
    }
</script>

<style scoped>
  .demo-menu-container{
    max-width: 350px;

  }
  .right-title{
    text-align: center;
    border-bottom: 1px solid #cccccc;
    height: 40px;
  }
  .right-list{
    text-align: left;
  }
  .right-list a {
    color: #000000;
  }
  .right-list a:hover {
    color: #DB6862;
  }
  .right-icon {
    font-size: 10px;
    color: #cccccc;
  }
  .right-icon :hover{
    color: #DB6862;
  }
  .right-content{
    text-align: center;
  }
  .right-detail{
    font-size: 13px;
    color: #737373;
    margin: 10px 0 20px 0;
    padding: 10px 0 20px 0;
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
    margin: 7px 0 7px 0;
    padding: 5px 0 5px 0 ;
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
  .btn{
    background: #ccc;
    padding: 5px;
    color: white;
    font-size: 14px;
  }
  .btn-success{
    background: #5cb85c;
  }
</style>
