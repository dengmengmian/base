// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package column

import (
	"base/internal/app/biz"
	"base/internal/app/store"
)

// ColumnController 是 column 模块在 Controller 层的实现，用来处理示例模块的请求.
type ColumnController struct {
	b biz.IBiz
}

// New 创建一个 column controller.
func New(ds store.IStore) *ColumnController {
	return &ColumnController{b: biz.NewBiz(ds)}
}
