<template>
  <div class="layout">
      <mu-appbar title="Title">
        <mu-tabs :value="activeTab" @change="handleTabChange" class="tab">
          <mu-tab value="tab1" title="技术分享"/>
          <mu-tab value="tab2" title="技术翻译"/>
          <mu-tab value="tab3" title="综合文章"/>

        </mu-tabs>
        <mu-text-field icon="search" class="appbar-search-field"  slot="right" hintText="请输入搜索内容"/>
        <mu-flat-button color="white" label="搜索" slot="right"/>
        <mu-icon-menu icon="more_vert" slot="right">
          <mu-menu-item title="菜单 1"/>
          <mu-menu-item title="菜单 2"/>
          <mu-menu-item title="菜单 3"/>
          <mu-menu-item title="菜单 4"/>
          <mu-menu-item title="菜单 5"/>
        </mu-icon-menu>
      </mu-appbar>
        <!--导航-->
    <div class="content">
      <div class="breadcrumb">
        <!--面包屑导航-->
        <mu-breadcrumb>
          <mu-breadcrumb-item href="/">Home</mu-breadcrumb-item>
          <mu-breadcrumb-item href="/">VideoGame</mu-breadcrumb-item>
          <mu-breadcrumb-item>Download</mu-breadcrumb-item>
        </mu-breadcrumb>
      </div>
      <div class="content-body">
        <mu-row class="body-content" gutter>
          <mu-col width="100" tablet="100" desktop="100">
            <div id="notify">
              <span>未设置登录密码，请前往 <a href="">修改密码</a> 页面进行设置。设置后将可以在移动设备上使用邮箱登录网站。</span>
            </div>
          </mu-col>
          <mu-col width="100" tablet="70" desktop="70">
            <mu-paper class="article"  :zDepth="2" >
              <!--标签bar 开始-->
              <div>
                <mu-tabs :value="activeTabA" @change="handleTabChangeA" class="bar">
                  <mu-tab  class="title-bar" value="t1" title="活跃"/>
                  <mu-tab  class="title-bar" value="t2" title="精华"/>
                  <mu-tab class="title-bar" value="t3" title="最新"/>
                  <mu-tab class="title-bar" value="t4" title="最热"/>
                  <mu-tab class="title-bar" value="t4" title="零回复"/>
                </mu-tabs>
              </div>
              <!--标签bar 结束-->

              <!--文章列表 开始-->
              <div class="article-list">
                <mu-list-item title="">
                  <mu-avatar src="http://www.muse-ui.org/images/avatar1.jpg" slot="leftAvatar"/>
                  <span class="btn btn-green">分享</span>
                  <span>散落在指尖的阳光，我试着轻轻抓住光影的踪迹</span>
                  <span class="right-icon" slot="right"> 20/10/15k 5小时前</span>
                </mu-list-item>
              </div>
              <div class="article-list">
                <mu-list-item title="">
                  <span class="btn">分享</span>
                  <mu-avatar src="http://www.muse-ui.org/images/avatar1.jpg" slot="leftAvatar"/>
                  <span>散落在指尖的阳光，我试着轻轻抓住光影的踪迹</span>
                  <span class="right-icon" slot="right"> 20/10/15k 5小时前</span>
                </mu-list-item>
              </div>
              <div class="article-list">
                <mu-list-item title="">
                  <span class="btn btn-best">分享</span>
                  <mu-avatar src="http://www.muse-ui.org/images/avatar1.jpg" slot="leftAvatar"/>
                  <span>散落在指尖的阳光，我试着轻轻抓住光影的踪迹</span>
                  <span class="right-icon" slot="right"> 20/10/15k 5小时前</span>
                </mu-list-item>
              </div>
              <!--文章列表 结束-->
              <mu-card-header title="Myron Avatar" subTitle="sub title">
                <mu-avatar src="http://www.muse-ui.org/images/uicon.jpg" slot="avatar"/>
              </mu-card-header>
              <mu-card-title title="Content Title" subTitle="Content Title"/>
              <mu-card-text>
                散落在指尖的阳光，我试着轻轻抓住光影的踪迹，它却在眉宇间投下一片淡淡的阴影。调皮的阳光掀动了四月的心帘，温暖如约的歌声渐起。似乎在诉说着，我也可以在漆黑的角落里，找到阴影背后的阳光，找到阳光与阴影奏出和谐的旋律。我要用一颗敏感赤诚的心迎接每一缕滑过指尖的阳光！
              </mu-card-text>
            </mu-paper>
          </mu-col>
          <mu-col width="100" tablet="30" desktop="30">
            <div class="demo-menu-container">
              <mu-paper class="demo-menu" :zDepth="2" >

                <div class="right-content">
                  <mu-raised-button class="right-button" label="发起讨论" icon="record_voice_over" Default/>
                  <mu-raised-button class="right-button" label="发布文章" icon="edit" Default/>
                  <mu-raised-button class="right-button" label="提个问题" icon="help"  Default/>
                  <mu-raised-button class="right-button" label="分享此站" icon="share" Default/>
                </div>
              </mu-paper>
            </div>
            <div class="demo-menu-container">
              <mu-paper class="demo-menu" :zDepth="2" >
                <div class="right-title">
                  <span>讨论列表</span>
                </div>
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
            </div>
          </mu-col>
        </mu-row>
      </div>
    </div>
    <div class="footer">
      Muse-UI ©2017 Created by Muse-UI
    </div>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
      activeTabA: 't1',
      activeTab: 'tab1',
      activeTabR: 'tab4',
      activeList: 'list1'
    }
  },
  methods: {
    handleTabChange (val) {
      this.activeTab = val
    },
    handleTabChangeA (val) {
      this.activeTabA = val
    },
    handleTabChangeR (val) {
      this.activeTabR = val
    },
    handleListChange (val) {
      this.activeList = val
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .body-content {
    margin-left: 28px;
  }
  .content-body{
    min-height: 500px;
    border-radius: 5px;
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
  .right-title span {
    color: #cccccc;
    line-height: 40px;
    margin-left: 10px;
    font-size: 20px;
  }
  .right-button{
    margin: 12px;
  }

  .article{
    display: inline-block;
    margin: 16px 32px 16px 0;
  }
  .demo-menu {
    margin: 16px 32px 16px 0;
  }
  .layout{
    background-color: rgb(236, 236, 236);
  }
  .tab{
    margin: 0 auto;
    width: 300px;
    background-color: rgba(0, 0, 0, 0);
  }
  .content{
    width: 90%;
    margin: 0 auto;
  }
  .breadcrumb{
    margin: 10px 0;
  }
  .footer{
    padding: 20px 0;
    text-align: center;
  }
  .title-bar{
    background-color: white;
    color: #ccc;
  }
  .bar{
    width: 350px;
  }
  #notify{
    color: #AA7B3E;
    background: #FBF8E4;
    line-height: 40px;
    width: 90%;
  }
  #notify span{
    margin-left: 35px;
  }
  .article-list{
    margin-top: 5px;
    border-bottom: 1px solid #cccccc;
  }
  .btn{
    background: #ccc;
    padding: 5px;
    color: white;
    font-size: 10px;
  }
  .btn-green{
    background: #7BB75F;
  }
  .btn-default{
    background: #7BB75F;
  }
  .btn-best{
    background: #DB6661;
  }
</style>
