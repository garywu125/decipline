package main

// application source code location : 1.remote repository

import (
	"fmt"

	util "github.com/garywu125/decipline/internal/util"

	happy "github.com/garywu125/decipline/pkg/happy"
	hello "github.com/garywu125/decipline/pkg/hello"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	fmt.Println(" where :", util.Where())
	e.GET("/", hello.HelloWorld)
	e.GET("/happy", happy.HeppyWorld)
	e.Logger.Fatal(e.Start(":1324"))
}
