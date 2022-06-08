<template>
    <div class="container">
        <div class="register">
            <div class="item">
                <h3>注册界面</h3>
            </div>
            <div class="item">
              <span>账号:</span>
              <el-input size="mini" v-model="user_name" placeholder="请输入注册账号" />
            </div>
            <div class="item">
              <span>密码:</span>
              <el-input size="mini" type="password" v-model="password" placeholder="请输入注册密码" />
            </div>
            <div class="item">
              <span>确定密码:</span>
              <el-input size="mini" type="password" v-model="password_confirm" placeholder="请再次输入密码" />
            </div>
            <div class="item">
              <span>手机号:</span>
              <el-input size="mini" type="phone" v-model="phone" placeholder="请输入手机号" />
            </div>
            <div class="item">
              <span></span>
              <el-button size="mini" type="primary" @click="userRegister">注册</el-button>
              <el-button size="mini">取消</el-button>
            </div>
        </div>
    </div>
</template>

<script>
import { reactive, toRefs } from 'vue'
import {useRouter} from 'vue-router'
import {Register} from '../api/user'
export default {
    name: 'Register',
    setup(){
        let $router = useRouter()

        let userRegisterData = reactive({
            user_name:'',
            password:'',
            password_confirm:'',
            phone:''
        })
        //注册请求
        let userRegister = async () =>{
            let{user_name,password,password_confirm,phone} = userRegisterData
            //注册
            let code = await Register({
                user_name,
                password,
                password_confirm,
                phone,
            })
            if (code==200){
                
                    await $router.push('./login')
            }
        }
        return{
            ...toRefs(userRegisterData), //让数据保持响应式。将多个数据变成响应式数据。
            userRegister
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