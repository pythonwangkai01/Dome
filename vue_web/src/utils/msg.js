//封装消息框

//导入消息框
import {ElMessage} from 'element-plus'

//success 
export let $msg_success =(message,duration=3000) => {
    ElMessage({
        showClose:true,
        message:message,
        duration:duration,
        type:'success',
    })
}

//warning 
export let $msg_warning =(message,duration=3000) => {
    ElMessage({
        showClose:true,
        message:message,
        duration:duration,
        type:'warning',
    })
}

//error 
export let $msg_error =(message,duration=3000) => {
    ElMessage({
        showClose:true,
        message:message,
        duration:duration,
        type:'error',
    })
}