package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {

	//curl -s -XPOST -d'{"firstname":"Elon","lastname":"Musk","age":48}' http://localhost:9980/decode
	//Elon Musk is 48 years old!
	http.HandleFunc("/decode", func(writer http.ResponseWriter, request *http.Request) {
		var user User

		{
			//var user2 User
			//jsonStr := `{"firstname":"Peter","lastname":"Doe","age":25}`
			//json.Unmarshal([]byte(jsonStr), &user2)
			//fmt.Printf("%+v", user2)

			jsonStr := `{"firstname":"Peter","lastname":"Doe","age":25}`
			var f interface{}
			json.Unmarshal([]byte(jsonStr), &f)
			fmt.Println(f)
			//name := f.(map[string]interface{})["Name"].(string)
			//或者定义 var f map[string]interface{}
		}

		json.NewDecoder(request.Body).Decode(&user)

		fmt.Fprintf(writer, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "Peter",
			Lastname:  "Doe",
			Age:       25,
		}

		{

			jsonStr, _ := json.Marshal(peter)
			fmt.Printf("%v\n", string(jsonStr))
		}

		json.NewEncoder(w).Encode(peter)
	})

	http.ListenAndServe(":9980", nil)
}

/*
json.Valid(byte[]) //校验json字符串是否合法
json.Indent  //按照一定的格式缩进
json.Compact  //压缩
json.MarshalIndent //转换回带缩进的json字符串
*/
