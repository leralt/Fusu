package client

import (
	"github.com/gorilla/websocket"
)

var OnlineList map[string]Client

type Message struct {
	Name string `json:"name,omitempty"` //用户名
	Date string `json:"time"`           //发送消息的时间
	Msg  string `json:"msg,omitempty"`  //发送的消息
}

type Client struct {
	C    chan Message    //传递消息的通道
	Conn *websocket.Conn //websocket连接
	Name string          //用户名
}
