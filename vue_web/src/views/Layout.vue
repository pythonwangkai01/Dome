<!-- 布局页 -->
<template>
  <div class="container">
    <div class="left">
      <el-menu
        active-text-color="#ffd04b"
        background-color="#545c64"
        class="el-menu-vertical-demo"
        default-active="2"
        text-color="#fff"
      >
        <el-sub-menu index="1">
          <template #title>
            <el-icon><location /></el-icon>
            <span>Navigator One</span>
          </template>
          <el-menu-item-group title="Group One">
            <el-menu-item index="1-1">item one</el-menu-item>
            <el-menu-item index="1-2">item one</el-menu-item>
          </el-menu-item-group>
          <el-menu-item-group title="Group Two">
            <el-menu-item index="1-3">item three</el-menu-item>
          </el-menu-item-group>
          <el-sub-menu index="1-4">
            <template #title>item four</template>
            <el-menu-item index="1-4-1">item one</el-menu-item>
          </el-sub-menu>
        </el-sub-menu>
        <el-menu-item index="2">
          <el-icon><icon-menu /></el-icon>
          <span>Navigator Two</span>
        </el-menu-item>
        <el-menu-item index="3" disabled>
          <el-icon><document /></el-icon>
          <span>Navigator Three</span>
        </el-menu-item>
        <el-menu-item index="4">
          <el-icon><setting /></el-icon>
          <span>Navigator Four</span>
        </el-menu-item>
      </el-menu>
    </div>
    <div class="right">
        <div class="top">
          <el-menu
    :default-active="1"
    class="el-menu-demo"
    mode="horizontal"
    background-color="#545c64"
    text-color="#fff"
    active-text-color="#ffd04b"
  >
    <el-menu-item index="1">首页</el-menu-item>
    <el-menu-item index="2">邮件</el-menu-item>
    <el-menu-item index="3">账单</el-menu-item>
    <el-sub-menu index="4">
      <template #title>切换皮肤</template>
      <el-menu-item index="2-1">蓝色主题</el-menu-item>
      <el-menu-item index="2-2">绿色主题</el-menu-item>
      <el-menu-item index="2-3">黄色主题</el-menu-item>
      <el-menu-item index="2-4">黑色主题</el-menu-item>
    </el-sub-menu>
    <el-sub-menu index="5">
      <template #title>管理员</template>
      <el-menu-item index="3-1">个人中心</el-menu-item>
      <el-menu-item index="3-2">修改密码</el-menu-item>
      <el-menu-item index="3-3">退出系统</el-menu-item>
    </el-sub-menu>
    
  </el-menu>
        </div>
        <div class="bottom">3</div>
    </div>
  </div>
</template>

<script>
import {reactive,toRefs} from 'vue'

export default {
    name:'Layout',
    setup(){
        
        }
      

      
    }
</script>

<style scoped lang="scss">

.container{
  width: 100vw;
  height: 100vh;
  display: flex;
  .left{
    width: 200px;
  }
  .right{
    flex: 1;
    display: flex;
    flex-direction: column;
    .top{
      height: 60px;
    }
    .bottom{
      flex:1;
    }
  }

}
</style>