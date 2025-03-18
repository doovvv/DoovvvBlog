<template>
  <el-container style="height: 100vh">
    <!-- 侧边栏 -->
    <el-aside width="200px" style="background-color: #545c64">
      <Nav />
    </el-aside>

    <el-container>
      <!-- 头部 -->
      <el-header
        class="header"
        style="background-color: #409eff; color: #fff; line-height: 60px"
      >
        <Header />
      </el-header>

      <!-- 主内容区域 -->
      <el-main>
        <!-- 静态内容（可选） -->
        <el-card v-if="$route.path === '/admin'">
          <template #header>
            <div class="card-header">
              <span>欢迎使用博客管理后台</span>
            </div>
          </template>
          <div>
            <p>这里是博客管理后台的主内容区域。</p>
            <p>你可以在这里管理文章、设置系统等。</p>
          </div>
        </el-card>

        <!-- 路由内容 -->
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import Nav from "../components/nav.vue";
import Header from "../components/header.vue";

export default {
  name: "BlogAdminLayout",
  components: {
    Nav,
    Header,
  },
};
</script>

<style scoped>
.el-header {
  text-align: center;
  font-size: 22px;
}

.el-aside {
  color: #fff;
}

.el-main {
  background-color: #f0f2f5;
  padding: 20px;
}

.card-header {
  font-size: 18px;
  font-weight: bold;
}

.headerBtn {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
