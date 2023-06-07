package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
)

func Route(app *iris.Application) {

	p := app.Party("/api")
	demo(p)
}

func demo(p iris.Party) {
	p.Get("/", func(ctx *context.Context) {
		var resp app.Response
		resp.Code = "418"
		resp.Message = "demo response"
		tool.ResponseJSON(ctx, resp)
	})
}
