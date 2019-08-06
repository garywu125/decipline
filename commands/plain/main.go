package main

// application source code location : 1.remote repository

import (
	"fmt"

	happy "github.com/garywu125/decipline/api/happy"
	hello "github.com/garywu125/decipline/api/hello"
	util "github.com/garywu125/decipline/util"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	fmt.Println(" where :", util.Where())
	e.GET("/", hello.HelloWorld)
	e.GET("/happy", happy.HeppyWorld)
	e.Logger.Fatal(e.Start(":1324"))
}
