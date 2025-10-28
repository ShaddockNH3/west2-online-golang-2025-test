package websocket

import (
	"context"
	"net/http"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/constants"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
)

// 定义 WebSocket 升级器
var upgrader = websocket.HertzUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(ctx *app.RequestContext) bool {
		return true
	},
}

func serveHome(_ context.Context, c *app.RequestContext) {
	if string(c.URI().Path()) != "/" {
		hlog.Error("Not found", http.StatusNotFound)
		return
	}
	if !c.IsGet() {
		hlog.Error("Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	c.HTML(http.StatusOK, "index.html", nil)
}

func Init() {
	go hub.Run()
	h := server.Default(server.WithHostPorts(constants.WebsocketAddr))
	h.LoadHTMLGlob("index.html")

	h.GET("/", serveHome)
	h.GET("/ws", func(c context.Context, ctx *app.RequestContext) {
		serveWs(ctx)
	})
	h.Spin()
}
