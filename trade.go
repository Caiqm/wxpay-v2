package wxpay

import (
	"fmt"
	"time"
)

// TradeApplet 小程序统一下单接口 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_1
// POST https://api.mch.weixin.qq.com/pay/unifiedorder
func (c *Client) TradeApplet(param TradeApplet) (result TradeAppletPayRsp, err error) {
	if param.TradeType == "" {
		param.TradeType = "JSAPI"
	}
	tradeAppletRst := new(TradeAppletRsp)
	if err = c.doRequest("POST", param, &tradeAppletRst); err != nil {
		return
	}
	result.AppID = c.appId
	result.Timestamp = fmt.Sprintf("%d", time.Now().Unix())
	result.Package = fmt.Sprintf("prepay_id=%s", tradeAppletRst.PrepayId)
	result.NonceStr = c.createNonceStr()
	result.SignType = "MD5"
	result.PaySign = c.createAppletPaySign(result.Timestamp, tradeAppletRst.PrepayId, result.NonceStr)
	return
}

// TradeApp APP统一下单接口 https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_1
// POST https://api.mch.weixin.qq.com/pay/unifiedorder
func (c *Client) TradeApp(param TradeApp) (result *TradeAppRsp, err error) {
	if param.TradeType == "" {
		param.TradeType = "APP"
	}
	err = c.doRequest("POST", param, &result)
	return
}

// TradeJSAPI 微信内H5统一下单 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
// POST https://api.mch.weixin.qq.com/pay/unifiedorder
func (c *Client) TradeJSAPI(param TradeJSAPI) (result *TradeJSAPIRsp, err error) {
	if param.TradeType == "" {
		param.TradeType = "JSAPI"
	}
	err = c.doRequest("POST", param, &result)
	return
}

// TradeNative Native统一下单接口 https://pay.weixin.qq.com/wiki/doc/api/native.php?chapter=9_1
// POST https://api.mch.weixin.qq.com/pay/unifiedorder
func (c *Client) TradeNative(param TradeNative) (result *TradeNativeRsp, err error) {
	if param.TradeType == "" {
		param.TradeType = "NATIVE"
	}
	err = c.doRequest("POST", param, &result)
	return
}

// TradeWap H5支付 https://pay.weixin.qq.com/wiki/doc/api/H5.php?chapter=9_20&index=1
// POST https://api.mch.weixin.qq.com/pay/unifiedorder
func (c *Client) TradeWap(param TradeWap) (result *TradeWapRsp, err error) {
	if param.TradeType == "" {
		param.TradeType = "MWEB"
	}
	err = c.doRequest("POST", param, &result)
	return
}

// TradeOrderQuery 查询订单 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_2
// POST https://api.mch.weixin.qq.com/pay/orderquery
func (c *Client) TradeOrderQuery(param TradeOrderQuery) (result *TradeOrderQueryRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// TradeCloseOrder 关闭订单 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_3
// POST https://api.mch.weixin.qq.com/pay/closeorder
func (c *Client) TradeCloseOrder(param TradeCloseOrder) (result *TradeCloseOrderRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// TradeRefund 申请退款 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_4
// POST https://api.mch.weixin.qq.com/secapi/pay/refund
func (c *Client) TradeRefund(param TradeRefund) (result *TradeRefundRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// TradeRefundQuery 查询退款 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_5
// POST https://api.mch.weixin.qq.com/pay/refundquery
func (c *Client) TradeRefundQuery(param TradeRefundQuery) (result *TradeRefundQueryRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}
