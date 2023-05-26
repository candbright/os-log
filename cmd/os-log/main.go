package main

import (
	"github.com/candbright/go-core/core"
	"github.com/candbright/os-log/config"
	"github.com/gin-gonic/gin"
)

func main() {
	core.AppName(config.AppConfig.Get("application.name"))
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(core.LogHandler())
}
