package main

import (
	"fmt"
	"github.com/aminsaki/golang-clean-web-api/api"
	"github.com/aminsaki/golang-clean-web-api/config"
	"github.com/aminsaki/golang-clean-web-api/data"
)

func main() {
	cfg := config.GetConfig()
	data.InitRedis(cfg) //cache redis
	defer data.CloseRedis()
	api.InitServer(cfg)
	fmt.Println("Hello World")
}
