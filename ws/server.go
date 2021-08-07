// handler/ws/server.go
package ws

import (
	"Fusu/client"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var HistoryMsg []client.Message
var AllMsg = make(chan client.Message)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func EchoMessage(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // 实际应用时记得做错误处理
	go UserHandler(conn)

}

//用于接收消息发送给管理者
func UserHandler(conn *websocket.Conn) {
	defer func(conn *websocket.Conn) {
		delete(client.OnlineList, conn.RemoteAddr().String())
		err := conn.Close()
		if err != nil {
			fmt.Println("conn Close error", err)
			return
		}
	}(conn)

	_, theName, err := conn.ReadMessage()
	if err != nil {
		return
	}
	cli := client.Client{
		C:    make(chan client.Message),
		Conn: conn,
		Name: string(theName),
	}

	client.OnlineList[conn.RemoteAddr().String()] = cli
	go SendToClient(conn, cli)
	go func() {
		for _, message := range HistoryMsg {
			err := conn.WriteJSON(message)
			if err != nil {
				return
			}
		}
		for {
			message := client.Message{}
			err := conn.ReadJSON(&message)
			if err != nil {
				return
			}
			fmt.Println(message)
			if len(HistoryMsg) <= 60 {
				HistoryMsg = append(HistoryMsg, message)
			} else {
				HistoryMsg = HistoryMsg[1:]
				HistoryMsg = append(HistoryMsg, message)
			}

			AllMsg <- message
		}
	}()
	for {

	}

}

func Manager() {
	client.OnlineList = make(map[string]client.Client)
	for {
		msg := <-AllMsg
		for _, cli := range client.OnlineList {
			if cli.Name != msg.Name {
				cli.C <- msg
			}
		}
	}
}

func SendToClient(conn *websocket.Conn, cli client.Client) {
	for msg := range cli.C {
		err := conn.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}

func HistoryMessage() {

}
