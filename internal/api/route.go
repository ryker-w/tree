package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/lishimeng/go-log"
)

func Route(app *iris.Application) {

	app.Get("/some/{a}/p", func(ctx *context.Context) {
		var a = ctx.Params().Get("a")
		log.Info(a)
	})
}
