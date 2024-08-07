// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package routes

import (
	"base/internal/app/controller/v1/post"
	"base/internal/app/store"
	mw "base/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// 注册Post路由.
func PostRoutes(g *gin.RouterGroup) gin.IRoutes {
	pc := post.New(store.S)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 posts 路由分组
		postv1 := v1.Group("/posts")
		{
			postv1.GET(":postID", pc.Get) // 获取文章详情
			postv1.GET("", pc.List)       // 获取文章列表
			postv1.Use(mw.Authn())
			postv1.POST("", pc.Create)             // 创建文章
			postv1.PUT(":postID", pc.Update)       // 更新文章
			postv1.DELETE("", pc.DeleteCollection) // 批量删除文章
			postv1.DELETE(":postID", pc.Delete)    // 删除文章
		}
	}
	return nil
}
