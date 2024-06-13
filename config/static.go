package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("static", map[string]any{
		"path": config.Env("STATIC_PATH", "./public"),
	})
}
