// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package category

import (
	"base/internal/pkg/core"
	"base/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

// Get 获取指定的示例.
func (ctrl *CategoryController) Get(c *gin.Context) {
	log.C(c).Infow("Get category function called")

	category, err := ctrl.b.Categoyies().Get(c, c.Param("categoryID"))
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, category)
}