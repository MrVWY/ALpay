package middleware

import (
	"ALpay/config"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)
var (
	Enforcers *casbin.Enforcer
)


func C()  {
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// You can also use an already existing gorm instance with gormadapter.NewAdapterByDB(gormInstance)
	a, _ := gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/") // Your driver and data source.
	Enforcers, _ = casbin.NewEnforcer("casb/rbac_model.conf", a)

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// Load the policy from DB.
	_ = Enforcers.LoadPolicy()

	// Check the permission.
	//_, _ = Enforcers.Enforce("alice", "data1", "read")

	// Modify the policy.
	//Enforcers.AddPolicy("test1","root")
	//Enforcers.RemovePolicy("test1","root")

	// Save the policy back to DB.
	//_ = Enforcers.SavePolicy()
}

func EnableCookieSession() gin.HandlerFunc {
	Store = cookie.NewStore([]byte(config.KEY))
	return sessions.Sessions("SESSION",Store)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}

func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context){
		season := sessions.Default(c)
		sessionValue := season.Get("username")
		if sessionValue == nil {
			c.Abort()
			//c.JSON(http.StatusUnauthorized,gin.H{"queue":"访问未授权"})
			c.Redirect(http.StatusMovedPermanently,"/index")
		}
		c.Next()
		fmt.Println(c.Request.RemoteAddr)
		fmt.Println("I am after next")
	}
}

func CasbinWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		season := sessions.Default(c)
		sessionValue := season.Get("username")
		if sessionValue == "" {
			c.JSON(200,gin.H{
				"code":    401,
				"queue": "Unauthorized",
				"data":    "",
			})
			c.Abort()
			return
		}

		path := c.Request.URL.Path
		method := c.Request.Method
		result, err := Enforcers.Enforce(sessionValue,path,method)
		fmt.Println("CasbinWare")
		if err != nil {
			c.JSON(200,gin.H{
				"code":    401,
				"queue": "Unauthorized",
				"data":    "",
			})
			c.Abort()
			return
		}
		if !result {
			fmt.Println("permission check failed")
			c.JSON(200, gin.H{
				"code":    401,
				"queue": "Unauthorized",
				"data":    "",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}