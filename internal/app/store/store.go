// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package store

import (
	"context"
	"sync"

	"gorm.io/gorm"
)

var (
	once sync.Once
	// 全局变量，方便其它包直接调用已初始化好的 S 实例.
	S *datastore
)

// IStore 定义了 Store 层需要实现的方法.
type IStore interface {
	TX(context.Context, func(ctx context.Context) error) error
	DB() *gorm.DB
	Users() UserStore
	Posts() PostStore
	Examples() ExampleStore
	Configs() ConfigStore
	Categories() CategoryStore
	Columns() ColumnStore
	Tags() TagStore
	Projects() ProjectStore
}

// datastore 是 IStore 的一个具体实现.
type datastore struct {
	db *gorm.DB
}

// 确保 datastore 实现了 IStore 接口.
var _ IStore = (*datastore)(nil)

// NewStore 创建一个 IStore 类型的实例.
func NewStore(db *gorm.DB) *datastore {
	// 确保 S 只被初始化一次
	once.Do(func() {
		S = &datastore{db}
	})

	return S
}

// DB 返回存储在 datastore 中的 *gorm.DB.
func (ds *datastore) DB() *gorm.DB {
	return ds.db
}

// 事务相关代码.
type transactionKey struct{}

func (ds *datastore) Core(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(transactionKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return ds.db
}

func (ds *datastore) TX(ctx context.Context, fn func(ctx context.Context) error) error {
	return ds.db.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			ctx = context.WithValue(ctx, transactionKey{}, tx)
			return fn(ctx)
		},
	)
}

// Users 返回一个实现了 UserStore 接口的实例.
func (ds *datastore) Users() UserStore {
	return newUsers(ds.db)
}

// Posts 返回一个实现了 PostStore 接口的实例.
func (ds *datastore) Posts() PostStore {
	return newPosts(ds.db)
}

// Examples 返回一个实现了 ExampleStore 接口的实例.
func (ds *datastore) Examples() ExampleStore {
	return newExamples(ds.db)
}

// Configs 返回一个实现了 configStore 接口的实例.
func (ds *datastore) Configs() ConfigStore {
	return newConfigs(ds.db)
}

// Categories 返回一个实现了 categoryStore 接口的实例.
func (ds *datastore) Categories() CategoryStore {
	return newCategories(ds.db)
}

// Columns 返回一个实现了 columnStore 接口的实例.
func (ds *datastore) Columns() ColumnStore {
	return newColumns(ds.db)
}

// Tags 返回一个实现了 tagStore 接口的实例.
func (ds *datastore) Tags() TagStore {
	return newTags(ds.db)
}

// Projects 返回一个实现了 projectStore 接口的实例.
func (ds *datastore) Projects() ProjectStore {
	return newProjects(ds.db)
}
