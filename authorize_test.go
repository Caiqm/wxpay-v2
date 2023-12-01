package wxpay

import (
	"log"
	"testing"
)

var client *Client

func init() {
	var err error
	client, err = New("", "")
	if err != nil {
		log.Fatalln(err)
	}
	client.OnReceivedData(func(method string, data []byte) {
		log.Println(method, string(data))
	})
}

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

func TestClient_GetPhoneNumber(t *testing.T) {
	t.Log("========== GetPhoneNumber ==========")
	client.LoadOptionFunc(WithApiHost("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=ACCESS_TOKEN"))
	var p GetPhoneNumber
	p.Code = "" // 前端获取的code值
	r, err := client.GetPhoneNumber(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
