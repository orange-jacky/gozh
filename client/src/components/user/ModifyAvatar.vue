<template>
  <mu-col width="100" tablet="70" desktop="70">
    <mu-paper class="article"  :zDepth="2" >
      <div class="edit">
        <div class="edit-title">
          <h1><font-awesome-icon :icon="['fas','image']"/>请选择头像图片</h1>
          <mu-divider/>
        </div>
        <div class="edit-body">
          <h4>请选择图片： </h4>
          <div class="avatar">
            <img src="https://lccdn.phphub.org/uploads/avatars/6900_1524049601.jpg?imageView2/1/w/380/h/380" alt="">
          </div>
          <div>
            <mu-raised-button class="demo-raised-button" labelPosition="before" icon="folder" secondary label="选择文件">
              <input type="file" class="file-button">
            </mu-raised-button>
          </div>
          <mu-raised-button class="button" label="上传头像" secondary fullWidth/>
        </div>
      </div>
    </mu-paper>
  </mu-col>
</template>

<script>
    export default {
        name: "modify-avatar"
    }
</script>

<style scoped>
  .edit{
    margin: 15px;
  }
  .edit-title{
    padding-top: 30px;
  }
  .avatar{
    border: 1px solid #ccc;
    width: 386px;
    height: 386px;
    text-align: center;
    padding-top: 2px;
  }
  .avatar img{
    width: 380px;
    height: 380px;
  }
  .file-button{
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    opacity: 0;
  }
  .button{
    margin: 30px 0 30px 0 ;
  }
</style>
