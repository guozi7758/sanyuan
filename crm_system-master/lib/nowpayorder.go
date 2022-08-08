package lib

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const (
	TRADE_FUNCODE        = "WP001"
	QUERY_FUNCODE        = "MQ002"
	NOTIFY_FUNCODE       = "N001"
	FRONT_NOTIFY_FUNCODE = "N002"
	TRADE_TYPE           = "05"
	TRADE_CURRENCYTYPE   = "156"
	TRADE_CHARSET        = "UTF-8"
	TRADE_DEVICE_TYPE    = "20"
	TRADE_SIGN_TYPE      = "MD5"
	TRADE_QSTRING_EQUAL  = "="
	TRADE_QSTRING_SPLIT  = "&"
	TRADE_FUNCODE_KEY    = "funcode"
	TRADE_DEVICETYPE_KEY = "deviceType"
	TRADE_SIGNTYPE_KEY   = "mhtSignType"
	TRADE_SIGNATURE_KEY  = "mhtSignature"
	SIGNATURE_KEY        = "signature"
	SIGNTYPE_KEY         = "signType"
	VERIFY_HTTPS_CERT    = false
)

type NowPayServer struct {
	AppId        string `json:"appId"`
	Key          string `json:"key"`
	TradeTimeOut string `json:"tradeTimeOut"`
	ApiUrl       string `json:"apiUrl"`
	NotifyUrl    string `json:"notifyUrl"`
	TimeZone     string `json:"timeZone"`
	MhtOrderNo   string `json:"mhtOrderNo"`
	MhtOrderAmt  int64  `json:"mhtOrderAmt"`
	OutPutType   int    `json:"outputType"`
	SchoolName   string `json:"schoolName"`
}

var NowPay NowPayServer

var Key = "9945021714dfd33d"

/*
 * 描述: 获取聚合扫码接口
 *
 *********************************************************************************/
func (this *NowPayServer) NowPay() (string, error) {

	//拼接请求参数
	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	m["shopNo"] = 304491082990109
	m["orderNo"] = "G2022208008232543"
	m["money"] = 0.01

	//算出md5加密签名
	sign := PayLassSign(m, Key)
	m["sign"] = sign
	mjSon, _ := json.Marshal(m)
	mString := string(mjSon)
	fmt.Println("cet", mString)
	data, err := PayJsonPost("http://hxpay.longshunha.com/hxpay/hxpay/getPayUrl", mString)
	fmt.Println(string(data), err)
	//var val ResponseJson
	//json.Unmarshal(data, &val)
	//if val.ReturnCode != "01" && val.ResultCode != "01" {
	//	return "", err
	//}
	return "val.QrURL", nil
}

//下单签名
func PayLassSign(mReq map[string]interface{}, key string) string {

	//fmt.Println("========STEP 1, 对key进行升序排序.========")
	//fmt.Println("微信支付签名计算, API KEY:", key)
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}

	//sort.Strings(sorted_keys)

	//fmt.Println("========STEP2, 对key=value的键值对用&连接起来，略过空值========")
	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}
	//fmt.Println("========STEP3, 在键值对的最后加上key=API_KEY========")
	//STEP3, 在键值对的最后加上key=API_KEY

	if key != "" {
		signStrings = signStrings + "key=" + key
	}
	fmt.Println(signStrings)
	fmt.Println("sing", StrMd5Str(strings.ToLower(signStrings)))
	return StrMd5Str(signStrings)
}

/*
 * 描述: 初始化配置
 *
 *********************************************************************************/
func init() {
	jsonFile, err := os.Open("./config/nowpay.json")
	if err != nil {
		panic("打开文件错误，请查看:" + err.Error())
	}
	defer jsonFile.Close()
	data, err := ioutil.ReadAll(jsonFile)
	json.Unmarshal(data, &NowPay)
	fmt.Println("NowPay", NowPay)
}

func StrMd5Str(strPass string) string {
	w := md5.New()
	io.WriteString(w, strPass)
	return fmt.Sprintf("%x", w.Sum(nil))
}
