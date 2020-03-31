package server

import (
	"os"
	"qiqiChat/api"
	"qiqiChat/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	//r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser()) //从session中获取userID，然后从数据库中获取用户信息存储在c.Context中。

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("/ping", api.Ping)

		// 用户注册
		v1.POST("/user/register", api.UserRegister)

		// 用户登录
		v1.POST("/user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("/user/me", api.UserMe)
			auth.GET("/user/logout", api.UserLogout)

			auth.GET("/groups/info", api.GroupInfo)
			auth.POST("/groups/increasement", api.GroupAdd)
			auth.DELETE("/groups/decreasement/:GroupID", api.DelGroup)

			auth.POST("/staff/increasement", api.AddStaff)
			auth.DELETE("/staff/decreasement/:StaffID", api.DelStaff)
			auth.PUT("/staff/modify", api.UpdateStaff)
			auth.GET("/staff/info/:GroupID", api.GetStaffes)

			auth.GET("/products/:characteristic/info", api.GetProducts)
			auth.GET("/products/:characteristic/oneInfo", api.GetProductByName)
			auth.GET("/products/:characteristic/infoByTime", api.GetProductByTime)
			auth.PUT("/products/:characteristic/modification", api.ModifyProductCount)
			auth.POST("/products/:characteristic/increasement", api.AddProduct)
			auth.DELETE("/products/:characteristic/decreasement", api.DelProduct)
			auth.PUT("/products/:characteristic/changeName", api.ModifyProductName)

			auth.GET("/records/:characteristic/info", api.GetReceiveDetail)

			auth.GET("/leadOut/", api.GetRecordExcel)

			auth.POST("/lowValue/setRecord", api.LowValueAddRecords)
			auth.POST("/highValue/setRecord", api.HighValueAddRecords)
		}
	}
	return r
}
