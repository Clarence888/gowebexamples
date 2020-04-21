package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//表单实现
type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {

	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			tmpl.Execute(writer, nil)
			return
		}

		details := ContactDetails{
			Email:   request.FormValue("email"),
			Subject: request.FormValue("subject"),
			Message: request.FormValue("message"),
		}

		// do something with details
		//存入数据库等等

		_ = details

		fmt.Println(details.Email)

		//执行将已解析的模板应用于指定的数据对象， 将输出写入wr。
		//如果在执行模板或写入其输出时发生错误，执行将停止，但部分结果可能已经写入了输出写入器。
		//第二个参数是个接口 理论上任何类型都可以
		tmpl.Execute(writer, struct {
			Success bool
		}{true})
	})

	http.ListenAndServe(":8888", nil)

}
