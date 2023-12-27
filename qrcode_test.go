package wxpay

import (
	"fmt"
	"testing"
)

// 获取不限制的小程序码
func TestClient_GetWxACodeUnLimit(t *testing.T) {
	t.Log("========== GetWxACodeUnLimit ==========")
	client.LoadOptionFunc(WithApiHost("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESSTOKEN"))
	// 二维码参数
	var p GetWxACodeUnLimit
	p.Page = "pages/card/other-card"
	p.Scene = fmt.Sprintf("id=%d", 1)
	p.EnvVersion = "develop"
	p.CheckPath = false
	r, err := client.GetWxACodeUnLimit(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.Errmsg)
}

// 获取小程序二维码
func TestClient_CreateQRCode(t *testing.T) {
	t.Log("========== CreateQRCode ==========")
	client.LoadOptionFunc(WithApiHost("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESSTOKEN"))
	// 二维码参数
	var p CreateQRCode
	p.Path = "pages/card/other-card"
	r, err := client.CreateQRCode(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.Errmsg)
}
