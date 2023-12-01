package wxpay

// GetAccessToken 接口调用凭据 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html
// GET https://api.weixin.qq.com/cgi-bin/token
func (c *Client) GetAccessToken(param GetAccessToken) (result *GetAccessTokenRsp, err error) {
	if param.GrantType == "" {
		param.GrantType = "client_credential"
	}
	err = c.doRequest("GET", param, &result)
	return
}

// Code2Session 小程序登录 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/code2Session.html
// GET https://api.weixin.qq.com/sns/jscode2session
func (c *Client) Code2Session(param Code2Session) (result *Code2SessionRsp, err error) {
	if param.GrantType == "" {
		param.GrantType = "authorization_code"
	}
	err = c.doRequest("GET", param, &result)
	return
}

// GetPhoneNumber 获取手机号 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
// POST https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=ACCESS_TOKEN
func (c *Client) GetPhoneNumber(param GetPhoneNumber) (result *GetPhoneNumberRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}
