// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
	"github.com/gin-gonic/gin"
	controller0 "go-project/app/api/home/asset/controller"
	controller1 "go-project/app/api/home/asset/controller"
	controller2 "go-project/app/api/home/user/controller"
)
var controllers []interface{}

func RegisterController(controller interface{}) {
	controllers = append(controllers, controller)
}

func init() {
	RegisterController(controller0.AssetController{})
	RegisterController(controller1.Asset2Controller{})
	RegisterController(controller2.UserController{})
}

func Register(router *gin.Engine) {
	for _, v := range controllers {
		AutoRegisterRoutes(router, v)
	}
}

// 自动注册路由
func AutoRegisterRoutes(router *gin.Engine, controller interface{}) {
	controllerType := reflect.TypeOf(controller)
	controllerValue := reflect.ValueOf(controller)

	// 获取基础路由前缀
	baseRoute := buildBaseRoute(controllerType)

	// 遍历控制器的所有方法并注册路由
	for i := 0; i < controllerType.NumMethod(); i++ {
		method := controllerType.Method(i)
		methodName := getControllerRouterName(method.Name)

		// 注册方法为 Gin 的 GET 路由
		route := fmt.Sprintf("api/%s/%s", baseRoute, methodName)
		router.POST(route, func(ctx *gin.Context) {
			method.Func.Call([]reflect.Value{controllerValue, reflect.ValueOf(ctx)})
		})
	}
}

// 构建控制器的基础路由前缀
func buildBaseRoute(controllerType reflect.Type) string {
	// 解析包路径
	pkgPath := controllerType.PkgPath()
	pkgParts := strings.Split(pkgPath, "/")

	// 构建路由路径
	var routeBuilder strings.Builder
	flag := false

	for _, part := range pkgParts {
		if part == "controller" {
			break
		}
		if flag {
			routeBuilder.WriteString(part + "/")
		}
		if part == "home" {
			flag = true
		}
	}

	// 添加控制器名
	controllerName := getControllerRouterName(controllerType.Name())
	routeBuilder.WriteString(controllerName)

	return routeBuilder.String()
}

// 去掉字符串末尾的 "Controller" 并将首字母变为小写
func getControllerRouterName(input string) string {
	input = strings.TrimSuffix(input, "Controller")
	if input == "" {
		return input
	}

	runes := []rune(input)
	runes[0] = unicode.ToLower(runes[0])

	return string(runes)
}
