package main

import (
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/router"
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/config"
)

func main() {
	r := router.InitRouter()
	r.Run(config.Port)
}
