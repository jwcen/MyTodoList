package main

import (
	"MyTodoList/conf"
	"MyTodoList/routers"
)

func main() {
	conf.Init()
	r := routers.NewRouter()
	_ = r.Run(conf.HttpPort)
}
