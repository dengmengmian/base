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

// DeleteCollection 批量删除示例.
func (ctrl *ExampleController) DeleteCollection(c *gin.Context) {
	log.C(c).Infow("Batch delete example function called")

	exampleIDs := c.QueryArray("exampleID")
	if err := ctrl.b.Examples().DeleteCollection(c, c.GetString(known.XUsernameKey), exampleIDs); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
