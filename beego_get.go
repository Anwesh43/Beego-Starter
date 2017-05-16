package main

import "github.com/astaxie/beego"
import "github.com/astaxie/beego/context"

func main() {
	beego.Get("/test", func(ctx *context.Context) {
		ctx.Output.Body([]byte("Hello world"))
	})
	beego.Get("/test/:name", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello " + ctx.Input.Param("name")))
	})
	beego.Run()
}
