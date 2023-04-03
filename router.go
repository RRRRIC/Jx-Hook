package main

import (
	"jx-hook/biz/handler/alert"
	"jx-hook/biz/handler/sender"
	router "jx-hook/biz/router"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// register registers all routers.
func register(r *server.Hertz) {

	router.GeneratedRegister(r)

	senderRegister(r)
	alertRegister(r)

	hlog.Info("Finish router config")

}

func senderRegister(h *server.Hertz) {
	r := h.Group("/sender")

	r.PUT("/save", sender.Save)
	r.DELETE("/del/:id", sender.Del)
	r.GET("/get/:id", sender.Query)
	r.GET("/enable/:id", sender.Enable)
	r.GET("/disable/:id", sender.Disable)

	hlog.Info("Finished sender router register")
}

func alertRegister(h *server.Hertz) {
	r := h.Group("/alert")

	r.PUT("/save", alert.Save)
	r.DELETE("/del/:id", alert.Del)
	r.GET("/get/:id", alert.Query)
	r.GET("/enable/:id", alert.Enable)
	r.GET("/disable/:id", alert.Disable)
	r.GET("/do/:id", alert.Alert)

	hlog.Info("Finished alert router register")
}
