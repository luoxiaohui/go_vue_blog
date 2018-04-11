<template>
  <div id="sign_wrap">
    <h1>后台管理</h1>
    <el-input v-model="name" placeholder="请输入用户名"></el-input>
    <el-input v-model="password" placeholder="请输入密码" type="password"></el-input>
    <el-button type="primary" @click="signin">登录</el-button>
  </div>
</template>

<script>
  export default {
    name: 'signin',
    data() {
      return {
        name: '',
        password: '',
        hasName: false, // 用户名被占
      }
    },
    mounted: function () {

    },
    methods: {
      signin: function () {
        let _this = this;
        if (this.name.length < 6) {
          this.$message.error('用户名不能为空或小于六个字符')
          return
        }

        if (this.password.length < 6) {
          this.$message.error('密码不能为空或小于六个字符')
          return
        }
        this.$axios.post(global.signin,{
          'name': _this.name,
          'password': _this.password
        }).then(response => {
          console.log(response.data)
          var code = JSON.parse(response.data)["code"];
          var data = JSON.parse(response.data)["data"];
          if(code != "000"){
            _this.$message({
              message: data,
              type: 'success'
            })
          }else{
            _this.$message({
              message: data,
              type: 'success'
            })
            _this.$router.go(-1)
          }
          delete _this.password;
        }, error => {
          _this.$message({
              message: error,
              type: 'success'
            })
        })
      }
    }
  }
</script>

<style>
  #sign_wrap {
    width: 300px;
    margin: 200px auto;
  }

  #sign_wrap h1 {
    color: #383a42;
    padding: 10px;
  }

  #sign_wrap div {
    padding-bottom: 20px;
  }
</style>