// Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://www.base.cn

package main

import (
	"base/internal/app"
	_ "go.uber.org/automaxprocs"
	"os"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := app.NewBaseCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
