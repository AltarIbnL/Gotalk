package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetupRouter(mode string) *gin.Engine {
	//if mode == gin.ReleaseMode {
	//	gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	//}
	r := gin.New()
	// 全局限流
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(1*time.Second, 2))
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := r.Group("/api/v1")

	// 注册
	v1.POST("/signup", controller.SignUpHandler)
	// 登录
	// 可以获取token
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件

	{
		// 因为是装在wsl里面，虚拟机的ip地址会开机一次重启一次，所以每次都得重新输入
		v1.GET("/community", controller.CommunityHandler, middlewares.RateLimitMiddleware(1*time.Second, 2))
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		// 使用教程：填充好token以后，比如172.21.19.250:8081/api/v1/post/1 就是查询帖子id为1的帖子
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		// 使用教程 172.21.19.250:8081/api/v1/posts，查询所有的ID
		v1.GET("/posts/", controller.GetPostListHandler)
		v1.GET("/posts2/", controller.GetPostListHandler2)

		// 投票
		// 需要在body里面以json格式传入post_id 和 direction（-1,1 表示反对或者赞成）
		/*
			{
			    "post_id":"292926547168333824",
			    "direction":"1"
			}
		*/
		v1.POST("/vote", controller.PostVoteController)

	}
	//r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
	//	// 如果是登录的用户,判断请求头中是否有 有效的JWT  ？
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "ok",
	//	})
	//})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
