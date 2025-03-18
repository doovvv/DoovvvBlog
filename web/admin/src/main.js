import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import axios from "axios";
import router from "./router";
const app = createApp(App);
// 配置全局 axios 默认 URL
axios.defaults.baseURL = "http://localhost:3000/api/v1";

app.use(router);
app.mount("#app");
