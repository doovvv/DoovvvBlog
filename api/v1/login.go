package v1

import (
	"doovvvblog/middleware"
	"doovvvblog/model"
	"doovvvblog/utils/errmsg"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.User
	var token string
	fmt.Println("login")
	c.ShouldBindJSON(&user)
	code := model.CheckLogin(user.Username, user.Password)
	fmt.Println("login")
	if code == errmsg.SUCCESS {
		token, _ = middleware.CreateJwt(user.Username, user.Password)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrorMsg(code),
		"token":  token,
	})

}
