package wxpay

import "fmt"

const (
	kContentType = "application/x-www-form-urlencoded;charset=utf-8"
	kTimeFormat  = "2006-01-02 15:04:05"
)

const (
	kFieldAppId      = "appid"
	kFieldSecret     = "secret"
	kFieldMchId      = "mch_id"
	kFieldNonceStr   = "nonce_str"
	kFieldSign       = "sign"
	kFieldSignType   = "sign_type"
	kFieldErrCode    = "errcode"
	kFieldReturnCode = "return_code"
)

type Param interface {
	// NeedAppId 是否需要APPID，有的借口不需要APPID，比如：小程序二维码接口
	NeedAppId() bool

	// NeedSecret 是否需要密钥，有的借口不需要密钥，比如：支付接口，小程序获取手机号接口
	NeedSecret() bool

	// NeedSign 是否需要签名，有的接口不需要签名，比如：小程序登录与获取手机号接口
	NeedSign() bool

	// NeedVerify 是否对微信接口返回的数据进行签名验证， 为了安全建议都需要对签名进行验证，本方法存在是因为部分接口不支持签名验证。
	NeedVerify() bool

	// NeedTlsCert 是否需要证书，有的接口需要，比如：申请退款接口、提现到零钱用户接口
	NeedTlsCert() bool

	// ReturnType 返回类型，v2版本的接口都是xml的，为兼容小程序接口需要切换json
	ReturnType() string
}

type AuxParam struct {
}

func (aux AuxParam) NeedAppId() bool {
	return true
}

func (aux AuxParam) NeedSecret() bool {
	return false
}

func (aux AuxParam) NeedSign() bool {
	return true
}

func (aux AuxParam) NeedVerify() bool {
	return true
}

func (aux AuxParam) NeedTlsCert() bool {
	return false
}

func (aux AuxParam) ReturnType() string {
	return "json"
}

// ReturnCode 微信支付接口响应错误码
type ReturnCode string

func (c ReturnCode) IsSuccess() bool {
	return c == ReturnCodeSuccess
}

func (c ReturnCode) IsFailure() bool {
	return c != ReturnCodeSuccess
}

// ErrCode 微信小程序接口响应错误码
type ErrCode int32

func (c ErrCode) IsSuccess() bool {
	return c == ErrCodeSuccess
}

func (c ErrCode) IsFailure() bool {
	return c != ErrCodeSuccess
}

const (
	ReturnCodeSuccess ReturnCode = "SUCCESS" // 支付接口调用成功
	ErrCodeSuccess    ErrCode    = 0         // 小程序接口调用成功
)

// PayError 支付错误类
type PayError struct {
	ReturnCode ReturnCode `json:"return_code" xml:"return_code"`   // 此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnMsg  string     `json:"return_msg" xml:"return_msg"`     // 返回信息，如非空，为错误原因
	ErrCode    string     `json:"err_code" xml:"err_code"`         // 错误代码
	ErrCodeDes string     `json:"err_code_des" xml:"err_code_des"` // 错误信息描述
	ResultCode string     `json:"result_code" xml:"result_code"`   // 业务结果，SUCCESS/FAIL
}

func (e PayError) Error() string {
	errMsg := fmt.Sprintf("%s - %s", e.ReturnCode, e.ReturnMsg)
	if e.ErrCode != "" {
		errMsg = fmt.Sprintf("%s，%s - %s", errMsg, e.ErrCode, e.ErrCodeDes)
	}
	if e.ResultCode != "" {
		errMsg = fmt.Sprintf("%s，%s", errMsg, e.ResultCode)
	}
	return errMsg
}

func (e PayError) IsSuccess() bool {
	return e.ReturnCode.IsSuccess()
}

func (e PayError) IsFailure() bool {
	return e.ReturnCode.IsFailure()
}

// AppletError 小程序错误
type AppletError struct {
	Errcode ErrCode `json:"errcode"` // 错误信息
	Errmsg  string  `json:"errmsg"`  // 错误码
}

func (e AppletError) Error() string {
	return fmt.Sprintf("%d - %s", e.Errcode, e.Errmsg)
}

func (e AppletError) IsSuccess() bool {
	return e.Errcode.IsSuccess()
}

func (e AppletError) IsFailure() bool {
	return e.Errcode.IsFailure()
}
