package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

//websocket 一个简单的服务 回发收到的信息

//go get github.com/gorilla/websocket

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)

		for {
			//读取消息
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			//打印消息到console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			//消息写回给浏览器
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websocketssend.html")
	})

	http.ListenAndServe(":9971", nil)

}
