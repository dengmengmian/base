// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package known

const (
	// XRequestIDKey 用来定义 Gin 上下文中的键，代表请求的 uuid.
	XRequestIDKey = "X-Request-ID"

	// XPrjectIDKey  用来定义 Gin 上下文中的键，代表请求的 project id.
	XPrjectIDKey = "X-Project-ID"

	// XUsernameKey 用来定义 Gin 上下文的键，代表请求的所有者.
	XUsernameKey = "X-Username"

	// 日期格式化.
	TimeFormatDay = "20060102"
	TimeFormat    = "2006-01-02 15:04:05"

	// 公用状态.
	STATUS_OK      = 1
	STATUS_DISABLE = 2
)
