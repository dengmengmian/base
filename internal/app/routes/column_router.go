// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package routes

import (
	"base/internal/app/controller/v1/column"
	"base/internal/app/store"

	"github.com/gin-gonic/gin"
)

// 注册column路由.
func ColumnRoutes(g *gin.RouterGroup) gin.IRoutes {
	cf := column.New(store.S)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 columns 路由分组
		columnv1 := v1.Group("/columns")
		{
			columnv1.GET(":columnID", cf.Get) // 获取内容详情
			columnv1.GET("", cf.List)         // 获取内容列表
		}
	}
	return nil
}
