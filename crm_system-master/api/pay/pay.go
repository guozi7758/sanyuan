package pay

import (
	"fmt"
	"net/http"
	"strconv"

	"../../lib"
	"github.com/gin-gonic/gin"
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
	money, _ := strconv.ParseInt(c.Query("price"), 10, 64)
	if money <= 0 {
		c.JSON(http.StatusOK, gin.H{"err": 3, "msg": "余额不能为小于0"})
		return
	}
	nopay.MhtOrderAmt = money
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
