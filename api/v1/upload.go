package v1

import (
	"doovvvblog/utils"
	"doovvvblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	dst := utils.AppConfig.Server.FilePath + "/" + file.Filename
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"status": errmsg.SUCCESS,
		"msg":    errmsg.GetErrorMsg(errmsg.SUCCESS),
	})
}
