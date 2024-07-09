// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package example

import (
	"base/internal/pkg/core"
	"base/internal/pkg/known"
	"base/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

// Get 获取指定的示例.
func (ctrl *ExampleController) Get(c *gin.Context) {
	log.C(c).Infow("Get example function called")

	example, err := ctrl.b.Examples().Get(c, c.GetString(known.XUsernameKey), c.Param("exampleID"))
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, example)
}
