package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//get请求示例
func (this *MainController) Get() {
	this.Ctx.WriteString("hello world,I'm beego.")
}

//golang程序的运行入口
func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
