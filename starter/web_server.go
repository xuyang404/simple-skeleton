package starter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"simple-skeleton/boot"
)

var router *gin.Engine

func Router() *gin.Engine {
	return router
}

type WebServerStarter struct {
	boot.BaseStarter
}

func (b *WebServerStarter) Start(ctx boot.StaterContext) {
	conf := ctx.Conf()
	address := fmt.Sprintf("%v:%v", conf.App.Host, conf.App.Port)
	router = gin.Default()
	if conf.App.Debug {
		gin.SetMode(gin.DebugMode)
	}

	err := router.Run(address)
	if err != nil {
		log.Fatal(err)
	}
}
