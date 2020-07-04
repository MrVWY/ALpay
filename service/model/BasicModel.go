package model

import (
	"ALpay/config/payconfig"
	"ALpay/dao"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitModel()  {
	payconfig.Inits()
}

func Home(c *gin.Context) {
	c.HTML(200,"defualt.html",gin.H{})
}

func Index(c *gin.Context) {
	c.HTML(200,"index.html",gin.H{})
}

func LoginHandler(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.JSON(404,gin.H{"queue":"illegal Request"})
	}
	username := c.Query("username")
	password := c.Query("password")
	User := dao.CheckUser(username)
	if password == User.Pwd {
		session := sessions.Default(c)
		session.Set("username",username)
		_ = session.Save()
		c.HTML(200,"defualt.html",gin.H{})
	}else  {
		c.String(401,"Your Password is wrong")
	}
}

func BackBasicHandler(c *gin.Context)  {
	history := dao.Checkinventory_history()
	fmt.Println(history)
	//history := Paginator(5,10,50)
	c.HTML(200,"t.html", gin.H{"inventory_history":history,"ti":"ti"})
}