<template>
  <div style="background-color: #e8e8e8">
    <nav-bar></nav-bar>
    <div class="segment">
      <div class="book-header">
        <a class="image" href="#">
          <img class="image" v-bind:src="image">
        </a>
        <div class="content" id="content">
          <div class="header">
            <a href="#">
              {{book_name}}
            </a>
          </div>
          <div class="description">
            {{description}}
            <br/>
            <mu-raised-button label="开始阅读" style="background-color: #00B5AD;"/>
          </div>

        </div>
      </div>
      <mu-divider style="margin-bottom: 1em"/>
      <div>
        文章列表
      </div>
      <mu-list>
        <mu-list-item title="第一章">
          <mu-icon slot="right" value="info"/>
        </mu-list-item>
        <mu-list-item title="第一章">
          <mu-icon slot="right" value="info"/>
        </mu-list-item>
        <mu-list-item title="第一章">
          <mu-icon slot="right" value="info"/>
        </mu-list-item>
        <mu-list-item title="第一章">
          <mu-icon slot="right" value="info"/>
        </mu-list-item>
      </mu-list>
      <div>
      </div>
    </div>
    <my-footer></my-footer>
  </div>
</template>

<script>
  import MyFooter from './Footer'
  import NavBar from './Navber'

  export default {
    name: "CourseContents",
    components: {
      MyFooter,
      NavBar
    },
    data() {
      return {
        book_name: "Laravel 教程 - Web 开发实战入门 ( Laravel 5.5 )",
        image: "https://fsdhub.com/uploads/images/201704/28/1/O2sXyZ5ZH6.jpg?imageView2/1/w/200/h/200",
        description: "ajsghdasgufigasiuyfghyuiasghduiashguidhasduisahuid"
      }
    }
  }
</script>

<style scoped>
  .segment {
    width: 80%;
    padding: 2%;
    margin: 2em auto 2em;
    object-fit: cover;
    background-color: white;
    clear: both;
  }

  .book-header{
    /*display: -webkit-box;*/
    margin-bottom: 1em;
  }

  .image {
    width: 35%;
    float: left;
  }

  .content {
  }

  .content .header {
    padding: 2%;
    font-size: 20px;
  }

  .description {
    padding: 2%;
    font-weight: bold;
    width: 80%;
    word-wrap: break-word;
  }
</style>
