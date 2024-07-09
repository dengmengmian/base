// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package user

import (
	"base/internal/pkg/core"
	"base/internal/pkg/errno"
	"base/internal/pkg/known"
	"base/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

// Delete 删除一个用户.
func (ctrl *UserController) Delete(c *gin.Context) {
	log.C(c).Infow("Delete user function called")

	username := c.Param("name")
	if c.GetString(known.XUsernameKey) != username {
		core.WriteResponse(c, errno.ErrPermissionDeny, nil)
		return
	}

	if err := ctrl.b.Users().Delete(c, username); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
