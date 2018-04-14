<template>
  <mu-col width="100" tablet="70" desktop="70">
    <mu-paper class="article"  :zDepth="2" >
      <mu-card>
        <!--标签bar 开始-->
        <div>
          <mu-tabs :value="activeTabA" @change="handleTabChangeA" class="bar">
            <mu-tab  class="title-bar" value="t1" title="活跃"/>
            <mu-tab  class="title-bar" value="t2" title="精华"/>
            <mu-tab class="title-bar" value="t3" title="最新"/>
            <mu-tab class="title-bar" value="t4" title="最热"/>
            <mu-tab class="title-bar" value="t5" title="零回复"/>
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
        <div class="article-list">
          <mu-list-item title="">
            <span class="btn btn-best">分享</span>
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
        <div class="article-list">
          <mu-list-item title="">
            <span class="btn btn-best">分享</span>
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
        <div class="article-list">
          <mu-list-item title="">
            <span class="btn btn-best">分享</span>
            <mu-avatar src="http://www.muse-ui.org/images/avatar1.jpg" slot="leftAvatar"/>
            <span>散落在指尖的阳光，我试着轻轻抓住光影的踪迹</span>
            <span class="right-icon" slot="right"> 20/10/15k 5小时前</span>
          </mu-list-item>
        </div>
        <!--文章列表 结束-->
        <!--分页开始-->
        <div class="pagination">
          <mu-pagination :total="total" :current="current" @pageChange="handleClick">
          </mu-pagination>
        </div>
        <!--分页结束-->
        <!--网站提示-->
        <mu-card-text>
          <div class="notice">
            <span>GOZH是一个开源的社区，社区的建设都是技术类免费贡献自己的时间去设计和开发的，我们致力于打造一个活跃度高的，能真正解决问题的，高质量的Golang中文社区。这里需要你的加入为社区贡献你的一份力量。</span>
          </div>
        </mu-card-text>
        <!--网站提示-->
      </mu-card>
    </mu-paper>
  </mu-col>
</template>
<script>
    export default {
      name: "article-list",
      data() {
        return {
          activeTabA: 't1',
          total: 40,
          current: 1,
        }
      },
      methods: {
        handleTabChangeA (val) {
          this.activeTabA = val
        },
      }
    }
</script>

<style scoped>
  .title-bar{
    font-size: 12px;
    background-color: white;
    color: #ccc;
  }
  .bar{
    width: 350px;
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
  .btn-best{
    background: #DB6661;
  }
  .notice{
    font-size: 12px;
    color: #8B6C3E;
    background: #FBF8E4;
  }
  .notice span{
    line-height: 28px;
  }
  .pagination{
    margin-top: 15px;
  }

  .article{
    margin: 16px 32px 16px 0;
  }
  .right-icon {
    font-size: 10px;
    color: #cccccc;
  }
  .right-icon :hover{
    color: #DB6862;
  }
</style>
