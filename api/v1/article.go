package v1

import (
	"doovvvblog/model"
	"doovvvblog/utils/errmsg"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询用户列表
func GetArts(c *gin.Context) {
	pageSize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		pageSize = -1
	}
	pageNum, err := strconv.Atoi(c.Query("pagenum"))
	if err != nil {
		pageNum = -1
	}
	arts, total := model.GetArts(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": errmsg.SUCCESS,
		"data":   arts,
		"msg":    errmsg.GetErrorMsg(errmsg.SUCCESS),
		"total":  total,
	})
}

// 查询单个文章
func GetArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	art, code := model.GetArt(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   art,
		"msg":    errmsg.GetErrorMsg(code),
	})
}

// 查询分类下的所有文章
func GetCateArts(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("cid"))
	pageSize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		pageSize = -1
	}
	pageNum, err := strconv.Atoi(c.Query("pagenum"))
	if err != nil {
		pageNum = -1
	}
	arts, code, total := model.GetCateArt(cid, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   arts,
		"msg":    errmsg.GetErrorMsg(code),
		"total":  total,
	})
}

// 新增用户
func AddArt(c *gin.Context) {
	fmt.Println("add art")
	var art model.Article
	err := c.ShouldBindJSON(&art)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": errmsg.ERROR,
			"msg":    errmsg.GetErrorMsg(errmsg.ERROR),
			"data":   art,
		})
		return
	}

	code := model.CreateArt(&art)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   art,
		"msg":    errmsg.GetErrorMsg(code),
	})
}

// 删除文章
func DeleteArt(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": errmsg.ERROR,
			"msg":    errmsg.GetErrorMsg(errmsg.ERROR),
		})
	}

	code := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrorMsg(code),
	})
}

// 修改文章
func EditArt(c *gin.Context) {
	var art model.Article
	err := c.ShouldBindJSON(&art)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": errmsg.ERROR,
			"msg":    errmsg.GetErrorMsg(errmsg.ERROR),
		})
	}
	code := errmsg.SUCCESS
	if code = model.EditArt(art.ID, &art); code != errmsg.SUCCESS {
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
