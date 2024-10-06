package main

import (
	_ "gin_study/dao"
	_ "gin_study/models"
	"gin_study/router"
)

func main() {
	r := router.Router()
	r.Run() // default port 8080
}
