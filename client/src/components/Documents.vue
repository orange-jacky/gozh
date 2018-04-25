<template>  <div class="courses-bg">
  <nav-bar :active-tab="activeTab"></nav-bar>
  <div>
    <div class="courses-title">Golang中文文档</div>
    <br/>
    <div class="courses-items">
      <div class="courses-item" v-for="item in items">
        <a class="image" v-bind:href="item.src">
          <img class="image" v-bind:src="item.image">
        </a>
        <div class="content" id="content">
          <div class="header">
            <a v-bind:href="item.src">
              {{item.header}}
            </a>
          </div>
          <div class="description">
            {{item.description}}
          </div>
          <mu-raised-button label="开始阅读" style="background-color: #00B5AD; padding: 2%;"/>
        </div>
      </div>
    </div>
  </div>
  <my-footer></my-footer>
</div>

</template>

<script>
  import MyFooter from './Footer'
  import NavBar from './Navber'
  export default {
    name: "Documents",
    components: {
      MyFooter,
      NavBar
    },
    data() {
      return {
        activeTab: 'tab2',
        items: [
          {
            image: "https://www.baidu.com",
            header: "Laravel 教程 - Web 开发实战入门 ( Laravel 5.5 )",
            src:"http://www.baidu.com",
            description: "本书将教你如何使用 Laravel 一步一步构建一个类似新浪微博的应用，让你从实际开发中体会到 Laravel 开发的敏捷、愉悦与轻松。<br>" +
            "通过阅读本教程，你将学到如 HTML、CSS、JavaScript、PHP 和 Laravel 等 Web 开发相关的基础知识。本书还会对这些基础知识点进行延伸扩展，为你讲解一些在 Web" +
            "开发中更为专业、实用的技能，如 Git 工作流、Gulp 前端工作流、Bootstrap" +
            "框架基本使用等。这些知识将为你未来的编程开发奠定下坚实的基础。使你不论是在做自己的个人项目，或是构建一个伟大的商业产品时，都能得心应手。"
          },
          {
            image: "https://fsdhub.com/uploads/images/201704/28/1/O2sXyZ5ZH6.jpg?imageView2/1/w/200/h/200",
            header: "Laravel 教程 - Web 开发实战入门 ( Laravel 5 )",
            src:"http://www.baidu.com",
            description: "本书将教你如何使用 Laravel 一步一步构建一个类似新浪微博的应用，让你从实际开发中体会到 Laravel 开发的敏捷、愉悦与轻松。<br>" +
            "通过阅读本教程，你将学到如 HTML、CSS、JavaScript、PHP 和 Laravel 等 Web 开发相关的基础知识。本书还会对这些基础知识点进行延伸扩展，为你讲解一些在 Web" +
            "开发中更为专业、实用的技能，如 Git 工作流、Gulp 前端工作流、Bootstrap" +
            "框架基本使用等。这些知识将为你未来的编程开发奠定下坚实的基础。使你不论是在做自己的个人项目，或是构建一个伟大的商业产品时，都能得心应手。"
          },
          {
            image: "https://fsdhub.com/uploads/images/201704/28/1/O2sXyZ5ZH6.jpg?imageView2/1/w/200/h/200",
            header: "Laravel 教程 - Web 开发实战入门 ( Laravel 5.7 )",
            src:"http://www.baidu.com",
            description: "本书将教你如何使用 Laravel 一步一步构建一个类似新浪微博的应用，让你从实际开发中体会到 Laravel 开发的敏捷、愉悦与轻松。<br>" +
            "通过阅读本教程，你将学到如 HTML、CSS、JavaScript、PHP 和 Laravel 等 Web 开发相关的基础知识。本书还会对这些基础知识点进行延伸扩展，为你讲解一些在 Web" +
            "开发中更为专业、实用的技能，如 Git 工作流、Gulp 前端工作流、Bootstrap" +
            "框架基本使用等。这些知识将为你未来的编程开发奠定下坚实的基础。使你不论是在做自己的个人项目，或是构建一个伟大的商业产品时，都能得心应手。" +
            "asddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd" +
            "asdddddddddddddddddddddasdsadasdasdssaddddddddddddddddddddddddddddddddddddddddddddddddddd"
          }
        ]
      }
    },
    created() {
      window.onload = SetItemsHeight;
      window.onresize = SetItemsHeight;
    }
  }

  function SetItemsHeight() {
    // console.log(document.getElementsByClassName("content")[0].clientHeight);
    var contents = document.getElementsByClassName("content");
    var items_ = document.getElementsByClassName("courses-item");
    // console.log("items:"+items_);
    var size  = contents.length;
    // alert("size:" + size);
    for(var i = 0; i < size; i = i+1) {
      var contentHeight = contents[i].clientHeight;
      // console.log(contentHeight);
      // console.log(items_[i]);
      items_[i].style.height =  contentHeight * (1 + 0.4) + "px";
    }
  }
</script>

<style scoped>
  .courses-bg {
    background-color: #e8e8e8;
    min-height: 100%;
  }

  .courses-title {
    text-align: center;
    font-size: 26px;
    font-weight: normal;
    margin-top: 16px;
    margin-bottom: 2%;
    color: #636B6F;
  }

  .courses-items {
    margin: 0 auto;
    width: 100%;
  }

  .courses-item {
    width: 80%;
    height: 300px;
    padding: 2%;
    margin: 0 auto;
    margin-bottom: 2%;
    object-fit: cover;
    background-color: white;
    clear: both;
  }

  .image {
    width: 30%;
    position: absolute;
  }

  .content {
    margin-left: 13%;
    width: 65%;
    position: absolute;
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
