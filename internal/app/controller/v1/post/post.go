// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package post

import (
	"base/internal/app/biz"
	"base/internal/app/store"
)

// PostController 是 post 模块在 Controller 层的实现，用来处理文章模块的请求.
type PostController struct {
	b biz.IBiz
}

// New 创建一个 post controller.
func New(ds store.IStore) *PostController {
	return &PostController{b: biz.NewBiz(ds)}
}
