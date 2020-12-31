package main

import (
	"github.com/gorilla/websocket"
	cli "shiritori/gen/http/cli/shiritori"
)

func main() {
	dialer := &websocket.Dialer {
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	endpoint, payload, err := cli.ParseEndpoint(
		scheme,
		host,
		doer,
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		debug,
		dialer,
		connConfigurer,
	)
}