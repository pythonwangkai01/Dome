<template>
  <div class="container">
      <div class="login">
          <div class="item">
              <h2>Wang测试管理系统</h2>
          </div>
          <div class="item">
              <span>账号:</span>
              <el-input size="mini" v-model="user_name" placeholder="请输入账号" />
          </div>
          <div class="item">
              <span>密码:</span>
              <el-input size="mini" type="password" v-model="password" placeholder="请输入密码" />
          </div>
          <div class="item">
              <span></span>
              <el-button size="mini" type="primary" @click="userLogin">登录</el-button>
              <el-button size="mini">取消</el-button>
          </div>
          <div class="item">
              <span></span>
              <el-button size="mini" type="primary" @click="userRegister">注册</el-button>
          </div>
          <div class="item">
              <span>记住我:</span>
              <el-checkbox v-model="ckme" size="large" />
          </div>
      </div>
  </div>
</template>

<script>
import { reactive, toRefs } from 'vue'
import {useRouter} from 'vue-router'
import {Login} from '../api/user'

export default {
    name: 'Login',
    setup(){
        let $router = useRouter()

        let loginData = reactive({  
            user_name:'',
            password:'',
            ckme:false
        })
        //登录请求
        let userLogin = async () =>{
            let {user_name,password} = loginData
            //获取登录状态
            let code = await Login({user_name,password})
            //成功
            if (code==200){
                $router.push('./layout')
            }
        }
        return{
            ...toRefs(loginData),
            userLogin
        }

    }

}
</script>

<style scoped lang="scss">
.container{
    width: 100vw;
    height: 100vh;
    background: linear-gradient(to bottom,rgb(94, 94, 211),lightblue);

    display: flex;
    justify-content: center;
    align-items: center;
    .login{
        width: 600px;
        border: 1px solid #eee;
        border-radius: 8px;
        color:#eee;
        // font-family: 'bing'; 字体
        padding: 10px;
        .item{
            display: flex;
            font-size: 15px;
            align-items: center;
            margin: 10px 5px;
            h2{
                flex: 1;
                text-align: center;
            }
            span{
                width: 70px;
                text-align: right;
            }
        }
        
    }
}
</style>