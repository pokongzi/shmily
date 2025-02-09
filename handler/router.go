package handler

import (
	"github.com/gocraft/web"
)

type Context struct {
	HelloCount int
}

// NewRouter 创建并配置路由
func NewRouter() *web.Router {
	// 创建路由
	router := web.New(Context{})

	// 注册中间件
	router.Middleware(web.LoggerMiddleware)
	router.Middleware(web.ShowErrorsMiddleware)

	// 注册路由
	router.Get("/", (*Context).Login)

	return router
}
