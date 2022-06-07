// axios 封装请求方法
import axios from 'axios'

import {BASE_URL,TimeOut} from '../config/config'

const instance = axios.create({
    baseURL:BASE_URL,
    timeout:TimeOut,
})

import Nprogress from 'nprogress'
  //导入nprogress样式
import 'nprogress/nprogress.css'

//config
instance.interceptors.request.use(function(config) {
    // 在发送请求之前做些什么
    Nprogress.start();
    return config;
}, function (error){
    Nprogress.done();
    return Promise.reject(error);
});

//response
instance.interceptors.response.use(function (response) {
    // 对响应数据做点什么
    Nprogress.done();
    return response;
}, function (error) {
    // 对响应错误做点什么
    Nprogress.done();
    return Promise.reject(error);
});

export let $get =async (url,params)=>{
    let {data} =await instance.get(url,(params))
    return data
}

export let $post =async (url,params)=>{
    let {data} =await instance.post(url,params)
    return data
}

export let $delete =async (url,params)=>{
    let {data} =await instance.delete(url,params)
    return data
}

//设置token的方法，该方法会将浏览器缓存中的token信息，添加到请求头中
export let $setToken =()=>{
    instance.defaults.headers.common['Authorization']=sessionStorage.getItem('token');

}
export default instance