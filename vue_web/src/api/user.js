//用户api-实现增删改查

import {$get,$post,$setToken} from '../utils/request'
import {$msg_success,$msg_error} from '../utils/msg'
import md5 from 'js-md5'

//定义登录方法
export let Login = async (params) => {
    params.password = md5(params.password).toString()
    let {code,msg,data} = await $post('/user/login', params)
    if (code ==200){
        //浏览器缓存两种：sessionStorage 随着浏览器关闭清空
        // localStorage缓存数据，除了手动清除，否则会一直缓存在浏览器中
        sessionStorage.setItem('token', data.token)
        $setToken()
        $msg_success(msg)
    }else{
        $msg_error(msg)
    }
    return code
}

//定义注册方法
export let UserRegister = async (params) => {
    params.password = md5(params.password)
    params.password_confirm = md5(params.password_confirm)
    let {code,msg,data} = await $post('/user/register',params)
    if (code==200){
        $msg_success(msg)
    }else{
        $msg_error(msg)
    }
    return data.data
}