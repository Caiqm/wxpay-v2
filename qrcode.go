package wxpay

// GetWxACodeUnLimit 获取不限制的小程序码 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getUnlimitedQRCode.html
// POST https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN
func (c *Client) GetWxACodeUnLimit(param GetWxACodeUnLimit) (result *QrcodeRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// CreateQRCode 获取小程序二维码 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/createQRCode.html
// POST https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=ACCESS_TOKEN
func (c *Client) CreateQRCode(param CreateQRCode) (result *QrcodeRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// GetQRCode 获取小程序码 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getQRCode.html
// POST https://api.weixin.qq.com/wxa/getwxacode?access_token=ACCESS_TOKEN
func (c *Client) GetQRCode(param GetQRCode) (result *QrcodeRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}
