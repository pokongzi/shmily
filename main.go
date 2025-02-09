package main

import (
	"net/http"
	"shmily/common/mysql"
	"shmily/common/redis"
	"shmily/handler"
)

func main() {
	// 初始化数据库连接
	mysql.Init()
	redis.Init()

	// 创建路由
	router := handler.NewRouter()

	// 启动服务器
	http.ListenAndServe("localhost:3000", router)
}
