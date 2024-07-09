// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package routes

import (
	"base/internal/app/controller/v1/project"
	"base/internal/app/store"

	"github.com/gin-gonic/gin"
)

// 注册project路由.
func ProjectRoutes(g *gin.RouterGroup) gin.IRoutes {
	pj := project.New(store.S)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 projects 路由分组
		projectv1 := v1.Group("/projects")
		{
			projectv1.GET(":alias", pj.Get) // 获取project详情
		}
	}
	return nil
}
