// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"base/internal/app/routes"
	"base/internal/pkg/core"
	"base/internal/pkg/errno"
	"base/internal/pkg/log"
	mw "base/internal/pkg/middleware"
)

// installRouters 安装 base 接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})
	// 路由分组
	apiGroup := g.Group("/" + viper.GetString("url-path-prefix"))
	// 注册 /healthz handler.
	apiGroup.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	// project路由
	routes.ProjectRoutes(apiGroup)
	// healthz 不需要projectID，在 healthz之后添加 PojectID 中间件
	apiGroup.Use(mw.ProjectID())
	// 注册 pprof 路由
	//pprof.Register(g)
	//RABC
	//authz, err := auth.NewAuthz(store.S.DB())
	//if err != nil {
	//	return err
	//}

	// 用户路由
	routes.UserRoutes(apiGroup)
	// post路由
	routes.PostRoutes(apiGroup)
	// 示例路由
	routes.ExampleRoutes(apiGroup)
	// config路由
	routes.ConfigRoutes(apiGroup)
	// column路由
	routes.ColumnRoutes(apiGroup)
	// category路由
	routes.CategoryRoutes(apiGroup)
	// tag路由
	routes.TagRoutes(apiGroup)
	return nil
}
