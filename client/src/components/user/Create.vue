<template>
  <div class="layout">
    <nav-bar :active-tab="activeTab"></nav-bar>
    <div class="container">
      <mu-row gutter>
        <mu-col width="100" tablet="100" desktop="100">
          <mu-paper  :zDepth="2" >
            <div class="title">
              <h1><font-awesome-icon :icon="['far','comment']"/> 创建新的文章/话题</h1>
            </div>
            <div class="description">
              <div>
                <span>选择分类</span>
                <select class="select" name="" id="" v-model="value">
                  <option value="1">招聘</option>
                  <option value="2">问答</option>
                  <option value="3">分享</option>
                  <option value="4">教程</option>
                  <option value="5">生活</option>
                </select>
                <input class="input" type="text" placeholder="请填写标题">
              </div>
              <div class="mark-description">
                <ul>
                  <li>请注意单词拼写，以及中英文排版，<a href="https://github.com/sparanoid/chinese-copywriting-guidelines">参考此页</a></li>
                  <li>支持 Markdown 格式, **粗体**、~~删除线~~、`单行代码`, 更多语法请见这里 <a href="https://github.com/riku/Markdown-Syntax-CN/blob/master/syntax.md">Markdown</a> 语法</li>
                  <li>上传图片, 支持拖拽和剪切板黏贴上传, 格式限制 - jpg, png, gif</li>
                  <li>发布框支持本地存储功能，会在内容变更时保存，「提交」按钮点击时清空</li>
                </ul>
              </div>
            </div>
            <div class="editor">
              <mavon-editor id="edit" ref="md" @imgAdd="$imgAdd" @imgDel="$imgDel" v-model="input" :ishljs = "true"></mavon-editor>
            </div>
            <div class="editor">
              <mu-raised-button  class="button" label="发 布" secondary fullWidth/>
            </div>
          </mu-paper>
        </mu-col>
      </mu-row>
    </div>
  </div>
</template>

<script>
  import NavBar from '../Navber'
  export default {
    name: "create",
    data() {
      return {
        value: '1',
        input: '',
        activeTab:'tab1',
      }
    },
    components: {
      NavBar
    },
    methods: {
      handleChange (value) {
        this.value = value
      },
      $imgDel(pos){
        delete this.img_file[pos];
      },
      // 绑定@imgAdd event
      $imgAdd(pos, $file){
        // 第一步.将图片上传到服务器.
        var formdata = new FormData();
        formdata.append('image', $file);
        this.axios({
          url: '/upload',
          method: 'post',
          data: formdata,
          headers: { 'Content-Type': 'multipart/form-data' },
        }).then((url) => {
          // 第二步.将返回的url替换到文本原位置![...](./0) -> ![...](url)
          /**
           * $vm 指为mavonEditor实例，可以通过如下两种方式获取
           * 1. 通过引入对象获取: `import {mavonEditor} from ...` 等方式引入后，`$vm`为`mavonEditor`
           * 2. 通过$refs获取: html声明ref : `<mavon-editor ref=md ></mavon-editor>，`$vm`为 `this.$refs.md`
           */
          this.$refs.md.$img2Url(pos, url);
        })
      }
    }
  }
</script>

<style scoped>
  #edit{
    min-height: 500px;
  }
  .container{
    max-width: 1300px;
    width: 90%;
    margin: 0 auto;
  }
  .layout{
    background-color: rgb(236, 236, 236);
  }
  .title {
    margin: 30px;
    padding: 30px 30px 0 30px;
    border-bottom: 1px solid #ccc;
    text-align: center;
  }
  .editor {
    margin: 30px;
    padding-bottom: 30px;
  }
  .description {
    margin: 0 30px 0 30px;
  }
  .mark-description{
    width: 100%;
    margin: 20px 20px 0 0 ;
    border: 1px solid #ccc;
    color: #606060;
  }
  input[type="text"]{
    box-sizing: border-box;
    border-radius:4px;
    line-height: 25px;
    border:1px solid #c8cccf;
    color:#6a6f77;
    -web-kit-appearance:none;
    -moz-appearance: none;
    outline:0;
    padding:0 1em;
    text-decoration:none;
    width:80%;
  }
  input[type="text"]:focus{
    border:1px solid #ff7496;
  }
</style>
