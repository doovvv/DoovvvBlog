<template>
  <div>
    <h3>User List Page</h3>
    <el-card>
      <el-row>
        <el-col :span="24">
          <div class="search-container">
            <el-input
              v-model="userName"
              placeholder="Search username"
              clearable
              @keyup.enter="handleSearch"
              @clear="handleSearch"
            />
            <el-button type="primary" @click="handleSearch">
              <el-icon><Search /></el-icon>
            </el-button>
          </div>
        </el-col>
      </el-row>
      <el-table :data="userList">
        <el-table-column prop="username" label="Username" />
        <el-table-column
          prop="role"
          label="role"
          :formatter="formatRole"
          :filters="[
            { text: '管理员', value: '管理员' },
            { text: '订阅者', value: '订阅者' },
          ]"
          :filter-method="filterTag"
          filter-placement="bottom-end"
        >
          <template #default="scope">
            <el-tag
              :type="scope.row.role === 1 ? 'primary' : 'success'"
              disable-transitions
              >{{ formatRole(scope.row) }}</el-tag
            >
          </template>
        </el-table-column>
        <el-table-column prop="CreatedAt" label="Created At" />
        <el-table-column align="right">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.row)">
              Edit
            </el-button>
            <el-button
              size="small"
              type="danger"
              @click="handleDelete(scope.$index, scope.row)"
            >
              Delete
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div>
        <el-button
          class="mt-4"
          style="width: 100%; margin-top: 20px"
          @click="dialogVisible = true"
        >
          Add User
        </el-button>
        <!-- 添加用户弹窗 -->
        <el-dialog v-model="dialogVisible" title="添加用户" width="30%">
          <el-form :model="newUser" label-width="100px">
            <el-form-item label="用户名">
              <el-input v-model="newUser.username" placeholder="请输入用户名" />
            </el-form-item>
            <el-form-item label="密码">
              <el-input
                v-model="newUser.password"
                placeholder="请输入密码"
                type="password"
              />
            </el-form-item>
            <el-form-item label="角色">
              <el-select v-model="newUser.role" placeholder="请选择角色">
                <el-option label="管理员" :value="1" />
                <el-option label="订阅者" :value="2" />
              </el-select>
            </el-form-item>
          </el-form>
          <template #footer>
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmAddUser">确认</el-button>
          </template>
        </el-dialog>
        <!-- 编辑用户弹窗 -->
        <el-dialog v-model="editDialogVisible" title="编辑用户" width="30%">
          <el-form :model="editUser" label-width="100px">
            <el-form-item label="用户名">
              <el-input
                v-model="editUser.username"
                placeholder="请输入用户名"
              />
            </el-form-item>
            <el-form-item label="角色">
              <el-select v-model="editUser.role" placeholder="请选择角色">
                <el-option label="管理员" :value="1" />
                <el-option label="订阅者" :value="2" />
              </el-select>
            </el-form-item>
          </el-form>
          <template #footer>
            <el-button @click="editDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmEditUser">确认</el-button>
          </template>
        </el-dialog>
      </div>
      <el-pagination
        layout="prev, pager, next"
        :total="total"
        :page-size="pageSize"
        :current-page="currentPage"
        @current-change="handleCurrentChange"
      />
    </el-card>
  </div>
</template>
<script setup>
import { ref } from "vue";
import { debounce, size } from "lodash";
import { Search } from "@element-plus/icons-vue";
import axios from "axios";
import router from "../../router";

const userName = ref("");
const userList = ref([]);
const total = ref(0);
const pageSize = ref(10);
const currentPage = ref(1);
const dialogVisible = ref(false); // 控制弹窗显示
const editDialogVisible = ref(false); // 控制编辑用户弹窗显示
const newUser = ref({
  username: "",
  password: "",
  role: 2,
});
const editUser = ref({
  ID: 0,
  username: "",
  role: 2,
});

// 角色映射
const roleMap = {
  2: "订阅者",
  1: "管理员",
};

// 编辑用户
const handleEdit = (row) => {
  editUser.value = {
    ID: row.ID, // 用户 ID
    username: row.username, // 用户名
    role: row.role, // 角色
  };
  editDialogVisible.value = true; // 打开编辑弹窗
};

const confirmEditUser = () => {
  console.log("编辑用户:", editUser.value);
  axios
    .put(`/users`, editUser.value)
    .then((response) => {
      console.log(response.data);
      if (response.data.status !== 200) {
        ElMessage.error(response.data.msg);
        return;
      }
      ElMessage.success("编辑用户成功");
      editDialogVisible.value = false; // 关闭弹窗
      getUserList(""); // 刷新用户列表
    })
    .catch((error) => {
      ElMessage.error("编辑用户失败");
      console.error(error);
    });
};

const confirmAddUser = () => {
  // console.log("添加用户:", newUser.value);
  axios
    .post("/users/add", newUser.value)
    .then((response) => {
      if (response.data.status !== 200) {
        ElMessage.error(response.data.msg);
        // console.log(response.data);
        return;
      }
      ElMessage.success("添加用户成功");
      dialogVisible.value = false; // 关闭弹窗
      getUserList(""); // 刷新用户列表
    })
    .catch((error) => {
      ElMessage.error("添加用户失败");
      console.error(error);
    });
};

const handleDelete = (index, row) => {
  axios.delete(`/users/${row.ID}`).then((response) => {
    if (response.data.status != 200) {
      ElMessage.error(response.data.msg);
      console.log(response.data);
      return;
    } else {
      ElMessage.success("删除成功");
      getUserList("");
    }
  });
};

// 格式化角色
const formatRole = (row) => {
  return roleMap[row.role] || "未知角色";
};

const filterTag = (value, row) => {
  return formatRole(row) === value;
};

const getUserList = (userName) => {
  axios
    .get("/users", {
      params: {
        username: userName,
        pagesize: pageSize.value,
        pagenum: currentPage.value,
      },
    })
    .then((response) => {
      if (response.data.status != 200) {
        ElMessage.error(response.data.msg);
        return;
      }
      userList.value = response.data.data;
      // console.log("用户列表数据:", userList.value);
      userList.value = userList.value.slice(
        (currentPage.value - 1) * pageSize.value,
        currentPage.value * pageSize.value
      );
      console.log("用户列表数据:", userList.value);
      total.value = response.data.total;
      // role.value = response.data.role;
      // console.log(response.data);
      // console.log("用户列表数据:", userList.value);
    })
    .catch((error) => {
      ElMessage.error("获取用户列表失败");
      console.error(error);
    });
};

const handleSearch = debounce(() => {
  console.log("执行搜索:", userName.value);
  // 这里调用你的搜索 API
  getUserList(userName.value);
}, 300);

// 分页事件处理
const handleCurrentChange = (newPage) => {
  currentPage.value = newPage; // 更新当前页码
  getUserList(""); // 重新获取用户列表数据
};

getUserList("");
</script>
<style scoped>
.search-container {
  display: flex;
  gap: 10px;
}

/* 图标样式 */
.search-icon {
  cursor: pointer;
  transition: color 0.2s;
  &:hover {
    color: var(--el-color-primary);
  }
}
</style>
