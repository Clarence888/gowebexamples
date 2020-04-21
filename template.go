package main

import (
	"html/template"
	"net/http"
)

//Go的html/template软件包

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	//从文件中解析模板
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//定义相关数据
		data := TodoPageData{
			PageTitle: "my todo list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		//在请求处理程序中执行模板
		tmpl.Execute(writer, data)
	})
	http.ListenAndServe(":9997", nil)
}

/*

控制结构					定义
{{斜杠*号注释}}			定义注释
{{.}}							渲染根元素
{{.Title}}						在嵌套元素中渲染“标题”字段
{{if .Done}} {{else}} {{end}}	定义一个if语句
{{range .Todos}} {{.}} {{end}}	遍历所有“ Todos”，并使用 {{.}}
{{block "content" .}} {{end}}	定义一个名称为“ content”的块
*/
