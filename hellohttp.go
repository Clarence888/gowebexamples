package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

/*
基本的http服务 需要具备以下能力
处理动态请求：处理浏览网站，登录帐户或发布图像的用户的传入请求。
提供静态资产：为浏览器提供JavaScript，CSS和图像，为用户创造动态体验。
接收连接： HTTP Server必须在特定端口上侦听才能接受来自Internet的连接。
*/

func main() {

	//处理动态请求
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, you ve request : %s \n", r.URL.Path)
	})
	/*
		注意 http.Request 包含请求参数等信息 可以拿到
	*/

	//处理静态资产

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// http.StripPrefix函数的作用之一，就是在将请求定向到你通过参数指定的请求处理处之前，将特定的prefix从URL中过滤出去。下面是一个浏览器或HTTP客户端请求资源的例子：
	///static/example.txt
	//StripPrefix 函数将会过滤掉/static/，并将修改过的请求定向到http.FileServer所返回的Handler中去，因此请求的资源将会是：
	///example.txt
	//http.FileServer 返回的Handler将会进行查找，并将与文件夹或文件系统有关的内容以参数的形式返回给你（在这里你将"static"作为静态文件的根目录）。因为你的"example.txt"文件在静态目录中，你必须定义一个相对路径去获得正确的文件路径。

	//使用 go get -u github.com/gorilla/mux 路由

	//创建一个新的路由器
	r := mux.NewRouter()
	//注册请求处理程序
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		// get the book
		// navigate to the page
		vars := mux.Vars(r)
		title := vars["title"] // the book title slug
		page := vars["page"]   // the page
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	r.HandleFunc("/nihaoget", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "nihao  get")
	}).Methods("GET")
	//限制为特定的HTTP方法。
	r.HandleFunc("/nihaopost", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "nihao  post")
	}).Methods("POST")

	r.HandleFunc("/xianzhidomain", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "限制特定域名")
	}).Host("localhost").Schemes("http") //限制为http

	//接收连接 启动两个端口http服务
	go http.ListenAndServe(":9999", nil)
	http.ListenAndServe(":9998", r)

}
