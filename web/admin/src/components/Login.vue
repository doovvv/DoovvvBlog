<template>
  <div class="container">
    <div class="loginbox">
      <h2 class="title">用户登录</h2>
      <el-form
        ref="loginFormRef"
        :model="formdata"
        :rules="rules"
        class="loginform"
        label-width="90px"
      >
        <el-form-item label="Username" prop="username">
          <el-input
            v-model="formdata.username"
            placeholder="请输入用户名"
          ></el-input>
        </el-form-item>
        <el-form-item label="Password" prop="password">
          <el-input
            v-model="formdata.password"
            type="password"
            placeholder="请输入密码"
            @keyup.enter.native="login"
          ></el-input>
        </el-form-item>
        <el-form-item class="loginbtn">
          <el-button @click="login" type="primary">登录</el-button>
          <el-button @click="resetForm" type="info">取消</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<style scoped>
.container {
  width: 100%;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(135deg, #4a90e2, #6a82fb);
}
.loginbox {
  width: 450px;
  box-shadow: 0px 10px 30px rgba(0, 0, 0, 0.2);
  padding: 20px;
  border-radius: 12px;
  background: #fff;
  text-align: center;
}
.loginform {
  width: 100%;
}
.loginbtn {
  display: flex;
  justify-content: space-between;
}
</style>
<script setup>
import { ref } from "vue";
import axios from "axios";
import router from "../router";

const formdata = ref({
  username: "",
  password: "",
});
const rules = ref({
  username: [
    { required: true, message: "Please input user name", trigger: "blur" },
    { min: 4, max: 12, message: "Length should be 4 to 12", trigger: "blur" },
  ],
  password: [
    { required: true, message: "Please input password", trigger: "blur" },
    { min: 6, max: 20, message: "Length should be 6 to 20", trigger: "blur" },
  ],
});
const loginFormRef = ref(null);

const resetForm = () => {
  formdata.value = {
    username: "",
    password: "",
  };
};
const login = () => {
  loginFormRef.value.validate((valid) => {
    if (valid) {
      // console.log('Form Data:', JSON.stringify(formdata.value, null, 2));
      axios
        .post("/users/login", formdata.value, {
          headers: {
            "Content-Type": "application/json",
          },
        })
        .then((response) => {
          if (response.data.status == 200) {
            ElMessage({
              message: "登录成功",
              type: "success",
              position: "top",
            });
            window.localStorage.setItem("token", response.data.token);
            router.push("admin/index");
          } else {
            ElMessage({
              message: "登录失败: " + response.data.msg,
              type: "error",
              position: "top",
            });
          }
        })
        .catch((err) => {
          ElMessage({
            message: "登录失败：" + err.message,
            type: "error",
            position: "top",
          });
        });
    } else {
      ElMessage({
        message: "输入非法数据，请重新输入",
        type: "error",
        position: "top",
      });
    }
  });
};
</script>
