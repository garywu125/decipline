package main

// package source : 1. filesystem 2.remote repository

import (
	happy "github.com/garywu125/api/happy"
	hello "github.com/garywu125/api/hello" // 这里使用的是filesystem相对路径

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", hello.HelloWorld)
	e.GET("/happy", happy.HeppyWorld)
	e.Logger.Fatal(e.Start(":1324"))
}
