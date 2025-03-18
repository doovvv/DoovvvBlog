<template>
              博客管理后台
              <el-button class="headerBtn" type="danger" size="small" style="float: right;" @click="loginOut">退出</el-button>
</template>
<script>
export default {
    methods: {
        loginOut() {
            localStorage.removeItem('token')
            location.href = '/login'
        }
    }

}
</script>