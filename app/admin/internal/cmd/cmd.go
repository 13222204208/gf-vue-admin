package cmd

import (
	"context"
	"gf-vue-admin/app/admin/internal/controller/auth"
	"gf-vue-admin/app/admin/internal/controller/menus"
	"gf-vue-admin/app/admin/internal/controller/roles"
	"gf-vue-admin/app/admin/internal/controller/users"
	"gf-vue-admin/app/admin/internal/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/admin", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)

				// 认证相关路由（不需要JWT验证）
				authController := auth.NewV1()
				group.POST("/auth/login", authController.Login)
				group.POST("/auth/refresh_token", authController.RefreshToken)

				// 需要JWT验证的路由
				group.Group("", func(protectedGroup *ghttp.RouterGroup) {
					protectedGroup.Middleware(middleware.JWTAuth)
					protectedGroup.Bind(
						users.NewV1(),
						roles.NewV1(),
						menus.NewV1(),
					)
					// 用户信息接口（需要JWT验证）
					protectedGroup.GET("/user/info", authController.GetUserInfo)
				})
			})
			s.Run()
			return nil
		},
	}
)
