package internal

import (
	"github.com/gin-gonic/gin"
	"shop/user/internal/pkg/core"
	"shop/user/internal/pkg/errno"
	"shop/user/internal/shop/biz"
	"shop/user/internal/shop/contorller"
	"shop/user/internal/shop/store"
)

// installRouters 安装 miniblog 接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		//log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	// 注册 pprof 路由
	//pprof.Register(g)

	//authz, err := auth.NewAuthz(store.S.DB())
	//if err != nil {
	//	return err
	//}

	userBiz := biz.NewBiz(store.S)
	userController := contorller.NewUserController(userBiz)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", userController.Create) // 创建用户
		}
	}
	return nil
}
