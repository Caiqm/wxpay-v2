# 微信相关接口（简易版）
微信V2版本支付（JSAPI支付，Native支付，APP支付，H5支付，小程序支付）、小程序登录、生成小程序二维码

## 安装

#### 启用 Go module

```go
go get github.com/Caiqm/wxpay-v2
```

```go
import wxpay "github.com/Caiqm/wxpay-v2"
```

## 如何使用

```go
// 实例化微信支付
var client, err = wxpay.New(appID, Secret)
```

#### 关于密钥（Secret）

密钥是微信应用的唯一凭证密钥，即 AppSecret，获取方式同 appid

## 加载支持配置

```go
// 加载支付接口链接，商户号ID与支付密钥
client.LoadOptionFunc(WithPayHost(), WithMchInformation(mchId, mchSecret))

// 或者一开始就加载支付接口链接，商户号ID与支付密钥
var client, err = wxpay.New(appID, Secret, WithPayHost(), WithMchInformation(mchId, mchSecret))
```

#### 关于加载支持配置

系统内置了几种可支付的配置

```go
// 设置请求链接，可自定义请求接口，传入host字符串
WithApiHost(HOST)

// 设置小程序登录链接
WithJsCodeHost()

// 设置支付请求链接
WithPayHost()

// 设置申请退款请求链接
WithRefundHost()

// 设置商户号信息，传入商户号ID与支付密钥
WithMchInformation(mchId, mchSecret)

// 也可自定义传入配置，返回以下类型即可
type OptionFunc func(c *Client)
```

## 小程序登录
```go
// 小程序登录code2Session
func TestClient_Code2Session(t *testing.T) {
	t.Log("========== Code2Session ==========")
	client.LoadOptionFunc(WithJsCodeHost())
	var p Code2Session
	p.JsCode = "" // 前端获取的code值
	r, err := client.Code2Session(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
```

## 小程序支付
```go
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
```