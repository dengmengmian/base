// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package routes

import (
	"base/internal/app/controller/v1/tag"
	"base/internal/app/store"

	"github.com/gin-gonic/gin"
)

// 注册tag路由.
func TagRoutes(g *gin.RouterGroup) gin.IRoutes {
	cf := tag.New(store.S)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 tags 路由分组
		tagv1 := v1.Group("/tags")
		{
			tagv1.GET(":tagID", cf.Get) // 获取tag详情
		}
	}
	return nil
}
