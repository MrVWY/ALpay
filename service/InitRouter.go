package service

import (
	"ALpay/middleware"
	"ALpay/service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterInit() (router *gin.Engine) {

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")
	//middleware
	router.Use(middleware.EnableCookieSession())
	router.Use(middleware.Cors())

	//basic page
	router.GET("/home", model.Home)
	router.GET("/index",model.Index)
	router.POST("/login", model.LoginHandler)

	//pay
	router.GET("/Jump", func(c *gin.Context) {
		money := c.Query("MilkyTea")
		if money == "1" {
			c.Redirect(http.StatusMovedPermanently,"/MilkyTea/1")
		}else if money == "10" {
			c.Redirect(http.StatusMovedPermanently,"/MilkyTea/10")
		}else if money == "100" {
			c.Redirect(http.StatusMovedPermanently,"/MilkyTea/100")
		} else {
			c.JSON(200,gin.H{"status":"ok"})
		}
	})

	//ReturnURL
	router.GET("/CompleteOrder", model.CompleteOrder)
	//NotifyURL
	router.POST("Verity", model.Verity)

	MilkyTea := router.Group("/MilkyTea")
	{
		MilkyTea.GET("/1", model.MakeOrder1Handler)
		MilkyTea.GET("/10", model.MakeOrder10Handler)
		MilkyTea.GET("/100", model.MakeOrder100Handler)
	}

	//Casbin-Backstage
	Backstage := router.Group("/Backstage")
	Backstage.Use(middleware.IsLogin(),middleware.CasbinWare())
	{
		Backstage.GET("/lo",model.BackBasicHandler)
	}

	return router
}