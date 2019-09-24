package main

import (
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/router"
)

func main() {

	r := router.InitRouter()
	r.Run()
}
