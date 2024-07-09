// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package project

import (
	"base/internal/app/biz"
	"base/internal/app/store"
)

// ProjectController 是 project 模块在 Controller 层的实现，用来处理示例模块的请求.
type ProjectController struct {
	b biz.IBiz
}

// New 创建一个 project controller.
func New(ds store.IStore) *ProjectController {
	return &ProjectController{b: biz.NewBiz(ds)}
}
