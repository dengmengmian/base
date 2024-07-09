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

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// Update 更新示例.
func (ctrl *ExampleController) Update(c *gin.Context) {
	log.C(c).Infow("Update example function called")

	var r v1.UpdateExampleRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	if err := ctrl.b.Examples().Update(c, c.GetString(known.XUsernameKey), c.Param("exampleID"), &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
