package main

import (
	_ "Test/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("hello world")
	beego.Run()
}

