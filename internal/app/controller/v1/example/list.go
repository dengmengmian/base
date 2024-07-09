// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package example

import (
	"base/internal/pkg/core"
	"base/internal/pkg/errno"
	"base/internal/pkg/known"
	"base/internal/pkg/log"
	v1 "base/pkg/api/v1"

	"github.com/gin-gonic/gin"
)

// List 返回示例列表.
func (ctrl *ExampleController) List(c *gin.Context) {
	log.C(c).Infow("List example function called.")

	var r v1.ListExampleRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	resp, err := ctrl.b.Examples().List(c, c.GetString(known.XUsernameKey), r.Offset, r.Limit)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, resp)
}
