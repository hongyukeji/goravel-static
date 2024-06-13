package static

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
	"github.com/hongyukeji/goravel-static/middleware"
)

const Binding = "static"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return nil, nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	// Add HTTP middleware
	facades.Route().GlobalMiddleware(middleware.Static())

	// 将包的配置文件发布到应用程序的 config 目录
	app.Publishes("github.com/hongyukeji/goravel-static", map[string]string{
		"config/static.go": app.ConfigPath("static.go"),
	})

}
