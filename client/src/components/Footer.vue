<template>
  <mu-row gutter>
    <mu-col width="100" tablet="25" desktop="25">
      <div class="footer">
        <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;  GOZH是一个开源的社区，社区的建设都是技术类免费贡献自己的时间去设计和开发的，我们致力于打造一个活跃度高的，能真正解决问题的，高质量的Golang中文社区。这里需要你的加入为社区贡献你的一份力量。</p>
        <div class="icon">
          <a href=""><font-awesome-icon :icon="['fab','github']"/></a>
          <a href=""><font-awesome-icon :icon="['fab','weixin']"/></a>
          <a href=""><font-awesome-icon :icon="['fab','weibo']"/></a>
        </div>
        <p>Designed by <span>❤</span>
          <a href="">Acorns</a> <span>❤</span>
          <a href="">冷</a> <span>❤</span>
          <a href="">大嘴的候</a>
        </p>
      </div>
    </mu-col>
    <mu-col width="100" tablet="25" desktop="25">
      <h2>赞助商</h2>
    </mu-col>
    <mu-col width="100" tablet="25" desktop="25">
      <h2>信息统计</h2>
    </mu-col>
    <mu-col width="100" tablet="25" desktop="25">
      <h2>其他信息</h2>
    </mu-col>
  </mu-row>
</template>

<script>
    export default {
        name: "footer"
    }
</script>

<style scoped>
  h2{
    color: #8A8A8A;
  }
  .footer{
    margin: 0 0 0 20px;
    text-align: left;
    font-size: 12px;
    color: #8A8A8A;
  }
  .footer span{
    color: #f4645f;
  }
  .footer a{
    color: #8A8A8A;
  }
  .footer a:hover{
    color: #f4645f;
  }
  .icon{
    font-size: 20px;
  }
</style>
