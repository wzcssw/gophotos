package main

import (
	"fmt"
	"net/http"
	"photos/config"
	"photos/controllers"
	"photos/dao"
	"photos/session"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	config.Load()                // 初始化配置文件
	dao.Initialize()             // 初始化数据库
	controllers.Intialize()      // 初始化模板
	session.InitSessionStorage() // 初始化Session

}

func main() {
	/////////////////////
	// Photos
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/img", controllers.ImgHandler)
	http.HandleFunc("/del", controllers.DelHandler)
	http.HandleFunc("/upload", controllers.UploadHandler)

	// SQL
	http.HandleFunc("/sql", controllers.SQLHandler)
	http.HandleFunc("/sql/edit", controllers.SQLEditHandler)
	http.HandleFunc("/sql/add", controllers.SQLAddHandler)
	http.HandleFunc("/sql/del", controllers.SQLDelHandler)

	// Session
	http.HandleFunc("/session", controllers.SessionHandler)
	http.HandleFunc("/session/del", controllers.SessionDelHandler)

	// 启动
	fmt.Println("starting server on port ", config.Conf.Port, "...")
	http.ListenAndServe(":"+config.Conf.Port, nil)
}
