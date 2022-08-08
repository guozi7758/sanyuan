package pay

import (
	"../../lib"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 * 描述: 聚合扫码支付订单
 * 路由: pay/nowpay_order
 * 类型: GET
 *
 ***********************************************************************/
func OrderNowPay(c *gin.Context) {
	fmt.Println(c.Query("text"))
	var nopay lib.NowPayServer
	//if c.Query("order_id") == "" {
	//	c.JSON(http.StatusOK, gin.H{"err": 3, "msg": "参数错误"})
	//	return
	//}
	payUrl, err := nopay.NowPay()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": 1, "msg": err.Error()})
		return
	}
	if payUrl == "" {
		c.JSON(http.StatusOK, gin.H{"err": 1, "msg": "支付已超时,换其他平台支付"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"err": 0, "msg": "base64", "qrpay": payUrl})
}
