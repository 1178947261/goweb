package main

import (
	"awesomeProject/goweb/bootstrap"
	"awesomeProject/goweb/http/middlewares"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

// Article  对应一条文章数据
type Article struct {
	Title, Body string
	ID          int64
}

var router = mux.NewRouter()

var Db *gorm.DB

func main() {

	bootstrap.SetupDB() //初始化数据库
	router = bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
