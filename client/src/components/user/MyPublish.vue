<template>
  <mu-col width="100" tablet="70" desktop="70">
    <mu-paper class="article"  :zDepth="2" >
      <div class="content">
        <div class="content-title">
          <h4><font-awesome-icon :icon="['fas','newspaper']"/> 最近发表的话题</h4>
        </div>
        <div class="list">
          <mu-row>
            <mu-col width="100" tablet="70" desktop="70">
              <a href="">一个 Excel 处理包，希望能有人加入共同</a>
            </mu-col>
            <mu-col width="100" tablet="25" desktop="25">
              <span>分享[12].点赞[666].3小时前</span>
            </mu-col>
          </mu-row>
          <mu-divider/>
        </div>
        <div class="list">
          <mu-row>
            <mu-col width="100" tablet="70" desktop="70">
              <a href="">一个 Excel 处理包，希望能有人加入共同</a>
            </mu-col>
            <mu-col width="100" tablet="25" desktop="25">
              <span>分享[12].点赞[666].3小时前</span>
            </mu-col>
          </mu-row>
          <mu-divider/>
        </div>
        <div class="list">
          <mu-row>
            <mu-col width="100" tablet="70" desktop="70">
              <a href="">一个 Excel 处理包，希望能有人加入共同</a>
            </mu-col>
            <mu-col width="100" tablet="25" desktop="25">
              <span>分享[12].点赞[666].3小时前</span>
            </mu-col>
          </mu-row>
          <mu-divider/>
        </div>
      </div>
    </mu-paper>
    <mu-paper class="article"  :zDepth="2" >
      <div class="content">
        <div class="content-title">
          <h4><font-awesome-icon :icon="['fas','comment-dots']"/> 最近发表的评论</h4>
        </div>
        <div class="list">
          <div>
            <a href="">使用 docker 部署你的 Laravel 项目 </a><span>at 1个月前</span>
            <p>可以的，只要你的服务器有docker</p>
            <mu-divider/>
          </div>
        </div>
        <div class="list">
          <div>
            <a href="">使用 docker 部署你的 Laravel 项目 </a><span>at 1个月前</span>
            <p><a style="color: #f4645f" href="">@Acorns</a> 可以的，只要你的服务器有docker</p>
            <mu-divider/>
          </div>
        </div>
        <div class="list">
          <div>
            <a href="">使用 docker 部署你的 Laravel 项目 </a><span>at 1个月前</span>
            <p>可以的，只要你的服务器有docker</p>
            <mu-divider/>
          </div>
        </div>
      </div>
    </mu-paper>
  </mu-col>
</template>

<script>
    import MuPaper from "muse-ui/src/paper/paper";
    import MuRow from "muse-ui/src/grid/row";
    export default {
      components: {
        MuRow,
        MuPaper},
      name: "my-publish"
    }
</script>

<style scoped>
  .content{
    margin: 15px 0 0 0;
  }
  h4{
    line-height: 35px;
    color: #606060;
  }
  .content-title{
    margin: 15px 0 0 0;
    text-align: center;
    border-bottom: 1px solid #ccc ;
    height: 35px;
  }
  .list span{
    font-size: 11px;
    color: #ccc;
  }
  .list a {
    color:#606060 ;
    line-height: 35px;
  }
  .list a:hover {
    color: #f4645f ;
  }
  .list {
    margin: 5px 20px 0 20px;
  }
  .content-list a {
    color: #606060 ;
  }
  .content-list span {
    color: #ccc;
    font-size: 12px;
    float: right;
  }
  .content-list a:hover {
    color:#f4645f ;
  }
  .alt_who {
    color: #f4645f;
  }
</style>
