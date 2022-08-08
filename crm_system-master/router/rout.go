package router

import (
	"../api/pay"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

var routs *gin.Engine

func pay_rout(name string) {
	rout := routs.Group(name)
	//rout.GET("pay_state", pay.PayOrderState)  // 检查订单支付是否成功
	rout.GET("nowpay_order", pay.OrderNowPay) // 聚合现在支付扫码订单支付
}

func init() {
	routs = gin.Default()
	Init()
	// routs = gin.New()
	routs.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	pay_rout("pay")

	routs.Run(Run.RunPort)
}
