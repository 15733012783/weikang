package pay

import (
	"context"
	"fmt"
	"github.com/15733012783/weikang/nacos"
	"github.com/smartwalle/alipay/v3"
	"log"
	"net/http"
	"strconv"
	"time"
)

func NewPayClient() *alipay.Client {
	APP_ID := nacos.ApiNac.AlipaySandbox.APPID
	//公钥
	PUBLIC_KEY := nacos.ApiNac.AlipaySandbox.PUBLICKEY
	//私钥
	PRIVATE_KEY := nacos.ApiNac.AlipaySandbox.PRIVATEKEY

	var client, err = alipay.New(APP_ID, PRIVATE_KEY, true)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if err = client.LoadAliPayPublicKey(PUBLIC_KEY); err != nil {
		fmt.Println(err)
		return nil
	}
	return client
}

func Pays(orderSnc string, price string) (string, error) {
	client := NewPayClient()
	var p = alipay.TradeWapPay{}
	p.NotifyURL = nacos.ApiNac.AlipaySandbox.NotifyURL //设置支付宝异步通知的回调URL，当支付状态发生变化时，支付宝会向该URL发送通知。
	p.ReturnURL = nacos.ApiNac.AlipaySandbox.ReturnURL //设置支付宝同步通知的回调URL，支付完成后用户将跳转回该URL
	p.Subject = "支付" + orderSnc                        //标题
	p.OutTradeNo = orderSnc                            //传递一个唯一单号
	p.TotalAmount = price                              //金额
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"           //产品代码
	currentTime := time.Now()
	unixTimestamp := currentTime.Unix()
	p.TimeoutExpress = strconv.FormatInt(unixTimestamp, 10) //设置超时时间，即订单有效期，以Unix时间戳形式表示。

	p.Body = "描述" //设置订单描述信息。
	var str, err = client.TradeWapPay(p)
	if err != nil {
		return "", err
	}
	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	fmt.Println(str.String())
	return str.String(), nil
}

// Refund 退款
func Refund(tradeNo string, refundAmount string) (*alipay.TradeRefundRsp, error) {
	client := NewPayClient()
	var p = alipay.TradeRefund{}
	p.RefundAmount = refundAmount
	p.TradeNo = tradeNo
	refunds, err := client.TradeRefund(context.Background(), p)
	if err != nil {
		return nil, err
	}
	return refunds, nil
}

// Callback 回调
func Callback(writer http.ResponseWriter, request *http.Request) string {
	client := NewPayClient()
	request.ParseForm()
	if err := client.VerifySign(request.Form); err != nil {
		log.Println("回调验证签名发生错误", err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("回调验证签名发生错误"))
		return "0"
	}
	log.Println("回调验证签名通过")
	// 示例一：使用已有接口进行查询
	var outTradeNo = request.Form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo

	rsp, err := client.TradeQuery(context.Background(), p)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf("验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())))
		return "0"
	}

	if rsp.IsFailure() {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf("验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg)))
	}
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write([]byte(fmt.Sprintf("订单 %s 支付成功", outTradeNo)))
	if err != nil {
		return "0"
	}
	return outTradeNo
}
