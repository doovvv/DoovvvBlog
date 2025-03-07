package v1

import (
	"doovvvblog/model"
	"doovvvblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询用户列表
func GetCates(c *gin.Context) {
	pageSize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		pageSize = -1
	}
	pageNum, err := strconv.Atoi(c.Query("pagenum"))
	if err != nil {
		pageNum = -1
	}
	cates := model.GetCates(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": errmsg.SUCCESS,
		"data":   cates,
		"msg":    errmsg.GetErrorMsg(errmsg.SUCCESS),
	})
}

// 新增用户
func AddCate(c *gin.Context) {
	var cate model.Category
	err := c.ShouldBindJSON(&cate)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": errmsg.ERROR,
			"msg":    errmsg.GetErrorMsg(errmsg.ERROR),
			"data":   cate,
		})
		return
	}
	code := model.CheckCategory(cate.Name)

	if code == errmsg.SUCCESS {
		model.CreateCate(&cate)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   cate,
		"msg":    errmsg.GetErrorMsg(code),
	})
}

// 删除用户
func DeleteCate(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": errmsg.ERROR,
			"msg":    errmsg.GetErrorMsg(errmsg.ERROR),
		})
	}

	code := model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrorMsg(code),
	})
}

// 修改用户
func EditCate(c *gin.Context) {
	var cate model.Category
	err := c.ShouldBindJSON(&cate)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": errmsg.ERROR,
			"msg":    errmsg.GetErrorMsg(errmsg.ERROR),
		})
	}
	code := model.CheckCategory(cate.Name)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errmsg.GetErrorMsg(code),
		})
	}
	if code = model.EditCate(cate.Id, &cate); code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errmsg.GetErrorMsg(code),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrorMsg(code),
	})

}
