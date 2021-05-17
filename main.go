package main

import (
	"TwoProject/common"
	"TwoProject/config"
	"TwoProject/controler"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func init() {
	var err error
	//DSN : username:password@protocol(address)/dbname?param=value
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", config.User, config.Password, config.Host, config.Port, config.Dbname)
	//fmt.Println(connStr)
	//设置连接方式：mysql
	common.Db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//访问数据库
	ctx := context.Background()
	err = common.Db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("数据库链接成功!")
}
func test() {
	//http.Handler()

}

func main() {
	server := http.Server{
		Addr:    ":8888",
		Handler: nil, // 添加中间层的。过滤。认证啥功能
		//Handler: &middleware.BasicAuthMiddleware{},
	}
	//http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("./wwwroot"))))
	// 开启路由对应的函数是controler->router.go
	controler.Route()
	log.Println("服务器已经启动了...")

	//go server.ListenAndServe()
	server.ListenAndServe()
}

/*
todo list
1.网页页面模版
2.数据库中间层
3.页面认证问题
4.网页接口api
*/
