package main

import (
	"context"
	"github.com/gorilla/websocket"
	"log"
	shiritorisvr "shiritori/gen/http/shiritori/server"
	"shiritori/pkg/msgq"
)

func newWebScocketCofigurer(logger *log.Logger) *shiritorisvr.ConnConfigurer {
	return nil

	return &shiritorisvr.ConnConfigurer{BattleFn: func(conn *websocket.Conn, cancel context.CancelFunc) *websocket.Conn {

		println("connn")
		return conn
		_, b, _ := conn.ReadMessage()
		msgq.Queue.Enqueue("1234", b)
		return conn
	}}
}
