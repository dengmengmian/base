// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package v1

// GetTagResponse 指定了 `GET /v1/tag/{tagID}` 接口的返回参数.
type GetTagResponse TagInfo

// TagInfo 指定了分类详细信息.
type TagInfo struct {
	TagID       string `json:"tagID,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
