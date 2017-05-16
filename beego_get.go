package main

import "github.com/astaxie/beego"
import "github.com/astaxie/beego/context"

type MainController struct {
	beego.Controller
}

func (this *MainController) Hello() {
	this.Ctx.Output.Body([]byte("hello world"))
}
func (this *MainController) HelloId() {
	this.Ctx.Output.Body([]byte("Hello " + this.Ctx.Input.Param(":id")))
}
func main() {
	beego.Get("/test", func(ctx *context.Context) {
		ctx.Output.Body([]byte("Hello world"))
	})
	beego.Get("/test/:name", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello " + ctx.Input.Param(":name")))
	})
	mainController := MainController{}
	beego.Router("/main/hello/", &mainController, "get:Hello")
	beego.Router("/main/hello/:id", &mainController, "get:HelloId")
	beego.Run()
}
