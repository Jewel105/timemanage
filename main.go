package main

import (
	"flag"
	"fmt"
	"gin_study/config"
	"gin_study/factory"
	"gin_study/gen/mysql"
	"gin_study/router"
)

var env string

func init() {
	flag.StringVar(&env, "env", "dev", "Specify the environment: dev, pro")
}

func main() {
	flag.Parse()
	fmt.Println("Running in environment:", env)
	config.GetConfig(env)
	mysql.Start()
	factory.RedisStart()
	router.Start()
}
