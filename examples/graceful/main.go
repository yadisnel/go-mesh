package main

import (
	"fmt"

	"github.com/yadisnel/go-ms/v2"
	"github.com/yadisnel/go-ms/v2/server"
)

func main() {
	service := micro.NewService()

	service.Server().Init(
		server.Wait(nil),
	)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
