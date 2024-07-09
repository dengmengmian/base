// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package biz

//go:generate mockgen -destination mock_biz.go -package biz app/internal/app/biz IBiz

import (
	"base/internal/app/biz/category"
	"base/internal/app/biz/column"
	"base/internal/app/biz/config"
	"base/internal/app/biz/example"
	"base/internal/app/biz/post"
	"base/internal/app/biz/project"
	"base/internal/app/biz/tag"
	"base/internal/app/biz/user"
	"base/internal/app/store"
)

// IBiz 定义了 Biz 层需要实现的方法.
type IBiz interface {
	Users() user.UserBiz
	Posts() post.PostBiz
	Examples() example.ExampleBiz
	Configs() config.ConfigBiz
	Columns() column.ColumnBiz
	Categoyies() category.CategoryBiz
	Tags() tag.TagBiz
	Projects() project.ProjectBiz
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*biz)(nil)

// biz 是 IBiz 的一个具体实现.
type biz struct {
	ds store.IStore
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*biz)(nil)

// NewBiz 创建一个 IBiz 类型的实例.
func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

// Users 返回一个实现了 UserBiz 接口的实例.
func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}

// Posts 返回一个实现了 PostBiz 接口的实例.
func (b *biz) Posts() post.PostBiz {
	return post.New(b.ds)
}

// Example 返回一个实现了 ExampleBiz 接口的实例.
func (b *biz) Examples() example.ExampleBiz {
	return example.New(b.ds)
}

// Configs 返回一个实现了 configBiz 接口的实例.
func (b *biz) Configs() config.ConfigBiz {
	return config.New(b.ds)
}

// Columns 返回一个实现了 columnBiz 接口的实例.
func (b *biz) Columns() column.ColumnBiz {
	return column.New(b.ds)
}

// Category 返回一个实现了 categoryBiz 接口的实例.
func (b *biz) Categoyies() category.CategoryBiz {
	return category.New(b.ds)
}

// Tags 返回一个实现了 tagBiz 接口的实例.
func (b *biz) Tags() tag.TagBiz {
	return tag.New(b.ds)
}

// Projects 返回一个实现了 projectBiz 接口的实例.
func (b *biz) Projects() project.ProjectBiz {
	return project.New(b.ds)
}
