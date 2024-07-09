// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package errno

// ErrTagNotFound 表示未找到分类信息.
var ErrTagNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.ErrTagNotFound", Message: "Tag was not found."}
