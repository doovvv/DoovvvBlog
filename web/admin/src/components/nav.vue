<template>
  <div class="blogName">
    <span> Doovvv's Blog </span>
  </div>
  <el-radio-group v-model="isCollapse" style="margin-bottom: 20px">
    <el-radio-button :value="false">expand</el-radio-button>
    <el-radio-button :value="true">collapse</el-radio-button>
  </el-radio-group>
  <el-menu
    default-active="2"
    class="el-menu-vertical-demo"
    :collapse="isCollapse"
    @select="goToPage"
  >
    <el-menu-item index="1" key="index">
      <el-icon><Home /></el-icon>
      <template #title>仪表盘</template>
    </el-menu-item>
    <el-sub-menu index="2">
      <template #title>
        <el-icon><Document /></el-icon>
        <span>文章管理</span>
      </template>
      <el-menu-item-group>
        <template #title><span>文章管理</span></template>
        <el-menu-item index="2-1" key="addart">写文章</el-menu-item>
        <el-menu-item index="2-2" key="artlist">文章列表</el-menu-item>
      </el-menu-item-group>
    </el-sub-menu>
    <el-menu-item index="3" key="catelist">
      <el-icon><Operation /></el-icon>
      <template #title>分类列表</template>
    </el-menu-item>
    <el-menu-item index="4" key="userlist">
      <el-icon><User /></el-icon>
      <template #title>用户列表</template>
    </el-menu-item>
  </el-menu>
</template>
<script setup>
import { ref } from "vue";
import router from "../router";
import {
  Document,
  User,
  Operation,
  House as Home,
} from "@element-plus/icons-vue";

const isCollapse = ref(true);

// const handleOpen = (key, keyPath) => {
//   console.log(key, keyPath);
// };

// const handleClose = (key, keyPath) => {
//   console.log(key, keyPath);
// };
const goToPage = (e) => {
  const routeMap = {
    1: "/index", // 仪表盘
    "2-1": "/addart", // 写文章
    "2-2": "/artlist", // 文章列表
    3: "/catelist", // 分类列表
    4: "/userlist", // 用户列表
  };

  const route = routeMap[e];
  if (route) {
    router.push("/admin" + route);
  } else {
    console.error("Invalid route index:", e);
  }
};
</script>
<style scoped>
.blogName {
  text-align: center;
  font-size: 22px;
  font-weight: bold;
}
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>
