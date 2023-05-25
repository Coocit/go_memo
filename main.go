package main

import (
	"go_memo/conf"
	"go_memo/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
