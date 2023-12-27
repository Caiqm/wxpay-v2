package wxpay

import (
	"fmt"
	"testing"
)

// 生成无限制小程序码
func TestClient_GetWxACodeUnLimit(t *testing.T) {
	t.Log("========== GetWxACodeUnLimit ==========")
	client.LoadOptionFunc(WithApiHost("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESSTOKEN"))
	// 二维码参数
	var p GetWxACodeUnLimit
	p.Page = "page/index/index"
	p.Scene = fmt.Sprintf("id=%d", 1)
	r, err := client.GetWxACodeUnLimit(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
