package main

import (
    "github.com/astaxie/beego"
)

type MainController struct {
    beego.Controller
}

func (this *MainController) Get() {
    this.Ctx.WriteString("hello world,I'm beego.")
}

func main() {
    beego.Router("/", &MainController{})
    beego.Run()
}