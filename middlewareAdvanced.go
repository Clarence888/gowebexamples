package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":9994", nil)
}

//Chain将middlewares应用于http.HandlerFunc 管理众多中间件
//... 代表 可变参数  是指函数传入的参数个数是可变的 ...type本质上是一个切片，也就是[]type
func Chain(f http.HandlerFunc, middlewareList ...Middleware) http.HandlerFunc {

	for _, m := range middlewareList {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

//记录所有请求链接 以及 执行时间
func Logging() Middleware {
	//创建中间件
	return func(f http.HandlerFunc) http.HandlerFunc {

		//定义http.HandlerFunc
		return func(writer http.ResponseWriter, request *http.Request) {
			start := time.Now()
			defer func() { log.Println(request.URL.Path, time.Since(start)) }()

			//调用下一个中间件
			f(writer, request)
		}
	}
}

func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(writer http.ResponseWriter, request *http.Request) {
			if request.Method != m {
				//未按照规则请求 报错
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			f(writer, request)
		}
	}
}

/*
$ curl -s http://localhost:8080/
hello world

$ curl -s -XPOST http://localhost:8080/
Bad Request
*/
