package middleware

import (
	"github.com/gin-contrib/static"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/gin"
)

func Static() http.Middleware {
	return func(ctx http.Context) {
		staticPath := facades.Config().Get("static.path", "./public").(string)
		static.Serve("/", static.LocalFile(staticPath, false))(ctx.(*gin.Context).Instance())
		ctx.Request().Next()
	}
}
