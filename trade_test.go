package wxpay

import (
	"testing"
)

const (
	mchId     = ""
	mchSecret = ""
)

// 小程序支付
func TestClient_TradeApplet(t *testing.T) {
	t.Log("========== TradeApplet ==========")
	client.LoadOptionFunc(WithPayHost(), WithMchInformation(mchId, mchSecret))
	var p TradeApplet
	p.Body = "支付测试"
	p.OutTradeNo = "TEST2023112717521212345678"
	p.TotalFee = "1"
	p.SpbillCreateIp = ""
	p.OpenId = ""
	p.NotifyUrl = "https://www.weixin.qq.com/wxpay/pay.php"
	r, err := client.TradeApplet(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

// app支付
func TestClient_TradeApp(t *testing.T) {
	t.Log("========== TradeApp ==========")
	client.LoadOptionFunc(WithPayHost(), WithMchInformation(mchId, mchSecret))
	var p TradeApp
	p.Body = "支付测试"
	p.OutTradeNo = ""
	p.TotalFee = "1"
	p.SpbillCreateIp = ""
	p.NotifyUrl = "https://www.weixin.qq.com/wxpay/pay.php"
	r, err := client.TradeApp(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

// 微信内H5支付
func TestClient_TradeJSAPI(t *testing.T) {
	t.Log("========== TradeJSAPI ==========")
	client.LoadOptionFunc(WithPayHost(), WithMchInformation(mchId, mchSecret))
	var p TradeJSAPI
	p.Body = "支付测试"
	p.OutTradeNo = "TEST2023112717521212345678"
	p.TotalFee = "1"
	p.SpbillCreateIp = ""
	p.OpenId = ""
	p.NotifyUrl = "https://www.weixin.qq.com/wxpay/pay.php"
	r, err := client.TradeJSAPI(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

// 关闭订单
func TestClient_TradeCloseOrder(t *testing.T) {
	t.Log("========== TradeCloseOrder ==========")
	client.LoadOptionFunc(WithApiHost("https://api.mch.weixin.qq.com/pay/closeorder"), WithMchInformation(mchId, mchSecret))
	var p TradeCloseOrder
	p.OutTradeNo = "TEST2023112717521212345678"
	r, err := client.TradeCloseOrder(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

// 查询订单
func TestClient_TradeOrderQuery(t *testing.T) {
	t.Log("========== TradeOrderQuery ==========")
	client.LoadOptionFunc(WithApiHost("https://api.mch.weixin.qq.com/pay/orderquery"), WithMchInformation(mchId, mchSecret))
	var p TradeOrderQuery
	p.OutTradeNo = "TEST2023112717521212345678"
	r, err := client.TradeOrderQuery(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

// 申请退款
func TestClient_TradeRefund(t *testing.T) {
	t.Log("========== TradeRefund ==========")
	client.LoadOptionFunc(WithApiHost("https://api.mch.weixin.qq.com/secapi/pay/refund"), WithMchInformation(mchId, mchSecret))
	var p TradeRefund
	p.OutTradeNo = "TEST2023112717521212345678"
	p.OutRefundNo = "TEST2023112717521212345678"
	p.TotalFee = "1"
	p.RefundFee = "1"
	r, err := client.TradeRefund(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

// 查询退款
func TestClient_TradeRefundQuery(t *testing.T) {
	t.Log("========== TradeRefundQuery ==========")
	client.LoadOptionFunc(WithApiHost("https://api.mch.weixin.qq.com/pay/refundquery"), WithMchInformation(mchId, mchSecret))
	var p TradeRefundQuery
	p.OutRefundNo = "TEST2023112717521212345678"
	r, err := client.TradeRefundQuery(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
