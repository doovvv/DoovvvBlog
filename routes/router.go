package routes

import (
	v1 "doovvvblog/api/v1"
	"doovvvblog/middleware"
	"doovvvblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppConfig.Server.Appmode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	auth := r.Group("/api/v1")
	auth.Use(middleware.JWTMiddleware())
	visitor := r.Group("/api/v1")
	{
		authUser := auth.Group("/users")
		{
			authUser.PUT("", v1.EditUser)
			authUser.DELETE("/:id", v1.DeleteUser)
			authUser.POST("/upload", v1.Upload)
		}
		visitorUser := visitor.Group("/users")
		{
			visitorUser.POST("/add", v1.AddUser)
			visitorUser.GET("", v1.GetUsers)
			visitorUser.GET("/:id", v1.GetUser)
			visitorUser.POST("/login", v1.Login)
		}
		authCate := auth.Group("/category")
		{
			authCate.POST("/add", v1.AddCate)
			authCate.PUT("/:id", v1.EditCate)
			authCate.DELETE("/:id", v1.DeleteCate)
		}
		visitorCate := visitor.Group("/category")
		{
			visitorCate.GET("", v1.GetCates)
		}
		authArt := auth.Group("/article")
		{
			authArt.POST("/add", v1.AddArt)
			authArt.PUT("/:id", v1.EditArt)
			authArt.DELETE("/:id", v1.DeleteArt)
		}
		visitorArt := visitor.Group("/article")
		{
			visitorArt.GET("", v1.GetArts)
			visitorArt.GET("/:id", v1.GetArt)
			visitorArt.GET("/cate/:cid", v1.GetCateArts)
		}
	}
	r.Run(utils.AppConfig.Server.Host + ":" + utils.AppConfig.Server.Port)

}
