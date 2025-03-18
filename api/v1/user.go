package v1

import (
	"doovvvblog/model"
	"doovvvblog/utils/errmsg"
	"doovvvblog/utils/vaildator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询用户列表
func GetUsers(c *gin.Context) {
	userName, ok := c.GetQuery("username")
	if !ok {
	 	userName = ""
	}
	pageSize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		pageSize = -1
	}
	pageNum, err := strconv.Atoi(c.Query("pagenum"))
	if err != nil {
		pageNum = -1
	}
	users,code := model.GetUsers(userName,pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   users,
		"msg":    errmsg.GetErrorMsg(code),
		"total":  len(users),
	})
}

// 查询单个用户
// func GetUser(c *gin.Context) {
	
// }

// 新增用户
func AddUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": errmsg.ERROR,
			"msg":    errmsg.GetErrorMsg(errmsg.ERROR),
			"data":   user,
		})
		return
	}
	msg, code := vaildator.Vaildator(&user)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    msg,
		})
		return
	}
	code = model.CheckUser(user.Username)

	if code == errmsg.SUCCESS {
		model.CreateUser(&user)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   user,
		"msg":    errmsg.GetErrorMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": errmsg.ERROR,
			"msg":    errmsg.GetErrorMsg(errmsg.ERROR),
		})
		return
	}

	code := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrorMsg(code),
	})
}

// 修改用户
func EditUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": errmsg.ERROR,
			"msg":    errmsg.GetErrorMsg(errmsg.ERROR),
		})
	}
	code := model.CheckUser(user.Username)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errmsg.GetErrorMsg(code),
		})
		return;
	}
	if code = model.EditUser(user.ID, &user); code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errmsg.GetErrorMsg(code),
		})
		return;
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrorMsg(code),
	})

}
