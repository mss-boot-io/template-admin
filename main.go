package main

import (
	_ "admin/apis"
	"admin/cmd"
)

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2024/1/25 10:40:18
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2024/1/25 10:40:18
 */

// @title chainide-admin API
// @version 0.0.1
// @description chainide-admin接口文档
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath
func main() {
	cmd.Execute()
}
