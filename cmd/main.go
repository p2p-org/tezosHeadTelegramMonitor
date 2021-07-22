package main

import (
	"context"

	controller "tezos_node_exporter/internal/controller"
)

const (
	postfix = "/chains/main/blocks/head"
	local   = "http://localhost:8732" + postfix
	remote  = "https://mainnet-tezos.giganode.io"
	chat_id = 320767500
)

func main() {
	c := controller.NewController(local, remote, chat_id)
	c.Run(context.TODO())
	panic("should not exit")
}
