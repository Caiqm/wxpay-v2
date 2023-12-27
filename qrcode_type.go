package wxpay

type Qrcode struct {
	AuxParam
}

func (g Qrcode) NeedAppId() bool {
	return false
}

func (g Qrcode) NeedSign() bool {
	return false
}

func (g Qrcode) NeedVerify() bool {
	return false
}

func (g Qrcode) ReturnType() string {
	return "jsonStr"
}

// GetWxACodeUnLimit 获取不限制的小程序码 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getUnlimitedQRCode.html
// 该接口用于获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制。
type GetWxACodeUnLimit struct {
	Qrcode
	Page       string    `json:"page"`                 // 默认是主页，页面 page，例如 pages/index/index，根路径前不要填加 /，不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面。scancode_time为系统保留参数，不允许配置
	Scene      string    `json:"scene"`                // 最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
	CheckPath  bool      `json:"check_path"`           // 默认是true，检查page 是否存在，为 true 时 page 必须是已经发布的小程序存在的页面（否则报错）；为 false 时允许小程序未发布或者 page 不存在， 但page 有数量上限（60000个）请勿滥用。
	EnvVersion string    `json:"env_version"`          // 要打开的小程序版本。正式版为 "release"，体验版为 "trial"，开发版为 "develop"。默认是正式版。
	Width      int       `json:"width"`                // 默认430，二维码的宽度，单位 px，最小 280px，最大 1280px
	AutoColor  bool      `json:"auto_color"`           // 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调，默认 false
	LineColor  LineColor `json:"line_color,omitempty"` // 二维码线条颜色，默认为黑色
	IsHyaline  bool      `json:"is_hyaline"`           // 默认是false，是否需要透明底色，为 true 时，生成透明底色的小程序
}

type LineColor struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

// CreateQRCode 获取小程序二维码 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/createQRCode.html
// 获取小程序二维码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制，详见获取二维码。
type CreateQRCode struct {
	Qrcode
	Path  string `json:"path"`  // 扫码进入的小程序页面路径，最大长度 128 个字符，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。scancode_time为系统保留参数，不允许配置。
	Width int    `json:"width"` // 二维码的宽度，单位 px。最小 280px，最大 1280px;默认是430
}

// GetQRCode 获取小程序码 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getQRCode.html
// 该接口用于获取小程序码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制。
type GetQRCode struct {
	Qrcode
	Path       string    `json:"path"`                 // 扫码进入的小程序页面路径，最大长度 1024 个字符，不能为空，scancode_time为系统保留参数，不允许配置；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
	Width      int       `json:"width"`                // 二维码的宽度，单位 px。默认值为430，最小 280px，最大 1280px
	AutoColor  bool      `json:"auto_color"`           // 默认值false；自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
	LineColor  LineColor `json:"line_color,omitempty"` // 默认值{"r":0,"g":0,"b":0} ；auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
	IsHyaline  bool      `json:"is_hyaline"`           // 默认是false，是否需要透明底色，为 true 时，生成透明底色的小程序
	EnvVersion string    `json:"env_version"`          // 要打开的小程序版本。正式版为 "release"，体验版为 "trial"，开发版为 "develop"。默认是正式版。
}

type QrcodeRsp struct {
	AppletError
	Buffer []byte `json:"buffer"` // 二维码二进制
}
