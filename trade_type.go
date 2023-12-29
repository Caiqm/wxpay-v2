package wxpay

// 请求参数
type Trade struct {
	AuxParam
	NotifyUrl string `xml:"notify_url" json:"notify_url"` // 接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数。公网域名必须为https，如果是走专线接入，使用专线NAT IP或者私有回调域名可使用http。
	// 必填，主要参数
	Body           string `xml:"body" json:"body"`                         // 商品描述交易字段格式根据不同的应用场景按照以下格式： APP——需传入应用市场上的APP名字-实际商品名称，天天爱消除-游戏充值。
	OutTradeNo     string `xml:"out_trade_no" json:"out_trade_no"`         // 商户系统内部订单号，要求32个字符内（最少6个字符），只能是数字、大小写字母_-|*且在同一个商户号下唯一。
	TotalFee       string `xml:"total_fee" json:"total_fee"`               // 订单总金额，单位为分
	SpbillCreateIp string `xml:"spbill_create_ip" json:"spbill_create_ip"` // 支持IPV4和IPV6两种格式的IP地址。调用微信支付API的机器IP
	TradeType      string `xml:"trade_type" json:"trade_type"`             // 支付类型
	// 选填，额外参数
	Attach        string `xml:"attach,omitempty" json:"attach,omitempty"`                 // 附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
	DeviceInfo    string `xml:"device_info,omitempty" json:"device_info,omitempty"`       // 自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB"
	Detail        string `xml:"detail,omitempty" json:"detail,omitempty"`                 // 商品详细描述，对于使用单品优惠的商户，该字段必须按照规范上传
	FeeType       string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`             // 标价币种，符合ISO 4217标准的三位字母代码，默认人民币：CNY
	TimeStart     string `xml:"time_start,omitempty" json:"time_start,omitempty"`         // 订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。
	TimeExpire    string `xml:"time_expire,omitempty" json:"time_expire,omitempty"`       // 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。
	GoodsTag      string `xml:"goods_tag,omitempty" json:"goods_tag,omitempty"`           // 订单优惠标记，使用代金券或立减优惠功能时需要的参数
	ProductId     string `xml:"product_id,omitempty" json:"product_id,omitempty"`         // 商品ID，trade_type=NATIVE时，此参数必传。此参数为二维码中包含的商品ID，商户自行定义。
	LimitPay      string `xml:"limit_pay,omitempty" json:"limit_pay,omitempty"`           // 指定支付方式，上传此参数no_credit--可限制用户不能使用信用卡支付
	Receipt       string `xml:"receipt,omitempty" json:"receipt,omitempty"`               // 电子发票入口开放标识，Y，传入Y时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效
	ProfitSharing string `xml:"profit_sharing,omitempty" json:"profit_sharing,omitempty"` // 是否需要分账，Y-是，需要分账 N-否，不分账 字母要求大写，不传默认不分账
	SceneInfo     string `xml:"scene_info,omitempty" json:"scene_info,omitempty"`         // 场景信息，WAP支付必填，该字段常用于线下活动时的场景信息上报，支持上报实际门店信息，商户也可以按需求自己上报相关信息。该字段为JSON对象数据，对象格式为{"store_info":{"id": "门店ID","name": "名称","area_code": "编码","address": "地址" }}
}

// TradeSceneInfo 场景信息
type TradeSceneInfo struct {
	Id       string `json:"id"`        // 门店编号，由商户自定义
	Name     string `json:"name"`      // 门店名称 ，由商户自定义
	AreaCode string `json:"area_code"` // 门店所在地行政区划码
	Address  string `json:"address"`   // 门店详细地址 ，由商户自定义
}

type TradeResponse struct {
	PayError
	AppID      string `xml:"appid" json:"appid"`                       // 调用接口提交的公众账号ID
	MchID      string `xml:"mch_id" json:"mch_id"`                     // 调用接口提交的商户号
	NonceStr   string `xml:"nonce_str" json:"nonce_str"`               // 微信返回的随机字符串
	Sign       string `xml:"sign" json:"sign"`                         // 微信返回的签名
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info"` // 调用接口提交的终端设备号
	TradeType  string `xml:"trade_type" json:"trade_type"`             // 调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，,H5支付固定传MWEB
	PrepayId   string `xml:"prepay_id" json:"prepay_id"`               // 微信生成的预支付会话标识，用于后续接口调用中使用，该值有效期为2小时,针对H5支付此参数无特殊用途
}

/* 小程序支付 */

// TradeApplet 小程序统一下单接口 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_1
type TradeApplet struct {
	Trade
	OpenId string `xml:"openid" json:"openid"` // trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识。
}

func (t TradeApplet) ReturnType() string {
	return "xml"
}

// TradeAppletRsp 小程序统一下单响应参数
type TradeAppletRsp struct {
	TradeResponse
	CodeUrl string `xml:"code_url,omitempty" json:"code_url"` // trade_type=NATIVE时有返回，此url用于生成支付二维码，然后提供给用户进行扫码支付。注意：code_url的值并非固定，使用时按照URL格式转成二维码即可。时效性为2小时
}

type TradeAppletPayRsp struct {
	AppID     string `json:"appId"`
	Timestamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

/* APP支付 */

// TradeApp APP统一下单接口 https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_1
type TradeApp struct {
	Trade
}

func (t TradeApp) ReturnType() string {
	return "xml"
}

// TradeAppRsp APP统一下单接口响应参数
type TradeAppRsp struct {
	TradeResponse
}

/* JSAPI支付·微信内H5支付 */

// TradeJSAPI 微信内H5统一下单接口 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
type TradeJSAPI struct {
	Trade
	OpenId string `xml:"openid" json:"openid"` // trade_type=JSAPI时（即JSAPI支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识。openid如何获取
}

func (t TradeJSAPI) ReturnType() string {
	return "xml"
}

// TradeJSAPIRsp TradeJSAPI 微信内H5统一下单响应参数
type TradeJSAPIRsp struct {
	TradeResponse
	CodeUrl string `xml:"code_url,omitempty" json:"code_url"` // trade_type=NATIVE时有返回，此url用于生成支付二维码，然后提供给用户进行扫码支付。注意：code_url的值并非固定，使用时按照URL格式转成二维码即可。时效性为2小时
}

/* Native支付·微信扫码支付 */

// TradeNative Native统一下单接口 https://pay.weixin.qq.com/wiki/doc/api/native.php?chapter=9_1
type TradeNative struct {
	Trade
}

func (t TradeNative) ReturnType() string {
	return "xml"
}

// TradeNativeRsp Native统一下单接口响应参数
type TradeNativeRsp struct {
	TradeResponse
	CodeUrl string `xml:"code_url" json:"code_url"` // trade_type=NATIVE时有返回，此url用于生成支付二维码，然后提供给用户进行扫码支付。注意：code_url的值并非固定，使用时按照URL格式转成二维码即可。时效性为2小时
}

/* MWEB支付，微信外H5支付 */

// TradeWap H5支付 https://pay.weixin.qq.com/wiki/doc/api/H5.php?chapter=9_20&index=1
type TradeWap struct {
	Trade
}

func (t TradeWap) ReturnType() string {
	return "xml"
}

type SceneInfo struct {
	H5Info struct {
		Type        string `json:"type"`
		WapURL      string `json:"wap_url,omitempty"`
		WapName     string `json:"wap_name,omitempty"`
		AppName     string `json:"app_name,omitempty"`
		BundleId    string `json:"bundle_id,omitempty"`
		PackageName string `json:"package_name,omitempty"`
	} `json:"h5_info,omitempty"`
	StoreInfo struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		AreaCode string `json:"area_code"`
		Address  string `json:"address"`
	} `json:"store_info,omitempty"`
}

// TradeWapRsp Native统一下单接口响应参数
type TradeWapRsp struct {
	TradeResponse
	MWebUrl string `xml:"mweb_url" json:"mweb_url"` // mweb_url为拉起微信支付收银台的中间页面，可通过访问该url来拉起微信客户端，完成支付,mweb_url的有效期为5分钟。
}

/* 查询订单 */

// TradeOrderQuery 查询订单 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_2
type TradeOrderQuery struct {
	AuxParam
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`     // 商户系统内部订单号，要求32个字符内（最少6个字符），只能是数字、大小写字母_-|*且在同一个商户号下唯一。
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"` // 微信的订单号，建议优先使用
}

func (t TradeOrderQuery) ReturnType() string {
	return "xml"
}

// TradeOrderQueryRsp 查询订单响应参数
type TradeOrderQueryRsp struct {
	PayError
	DeviceInfo         string `xml:"device_info" json:"device_info"`                             // 微信支付分配的终端设备号
	OpenId             string `xml:"openid" json:"openid"`                                       // 用户在商户appid下的唯一标识
	IsSubscribe        string `xml:"is_subscribe" json:"is_subscribe"`                           // 已废弃，默认统一返回N
	TradeType          string `xml:"trade_type" json:"trade_type"`                               // 调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，MICROPAY
	TradeState         string `xml:"trade_state" json:"trade_state"`                             // SUCCESS--支付成功 REFUND--转入退款 NOTPAY--未支付 CLOSED--已关闭 REVOKED--已撤销(刷卡支付) USERPAYING--用户支付中 PAYERROR--支付失败(其他原因，如银行返回失败) ACCEPT--已接收，等待扣款
	BankType           string `xml:"bank_type" json:"bank_type"`                                 // 银行类型，采用字符串类型的银行标识
	TotalFee           int    `xml:"total_fee" json:"total_fee"`                                 // 订单总金额，单位为分
	SettlementTotalFee int    `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee"` // 当订单使用了免充值型优惠券后返回该参数，应结订单金额=订单金额-免充值优惠券金额。
	FeeType            string `xml:"fee_type" json:"fee_type"`                                   // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY
	CashFee            int    `xml:"cash_fee,omitempty" json:"cash_fee"`                         // 现金支付金额订单现金支付金额
	CashFeeType        string `xml:"cash_fee_type,omitempty" json:"cash_fee_type"`               // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY
	CouponFee          int    `xml:"coupon_fee,omitempty" json:"coupon_fee"`                     // “代金券”金额<=订单金额，订单金额-“代金券”金额=现金支付金额
	CouponCount        int    `xml:"coupon_count,omitempty" json:"coupon_count"`                 // 代金券使用数量
	CouponType0        string `xml:"coupon_type_0,omitempty" json:"coupon_type_0"`               // CASH：充值代金券 NO_CASH：非充值优惠券 开通免充值券功能，并且订单使用了优惠券后有返回（取值：CASH、NO_CASH）。$n为下标,从0开始编号，举例：coupon_type_$0
	CouponId0          string `xml:"coupon_id_0,omitempty" json:"coupon_id_0"`                   // 代金券ID, $n为下标，从0开始编号
	CouponFee0         string `xml:"coupon_fee_0,omitempty" json:"coupon_fee_0"`                 // 单个代金券支付金额, $n为下标，从0开始编号
	TransactionId      string `xml:"transaction_id" json:"transaction_id"`                       // 微信支付订单号
	OutTradeNo         string `xml:"out_trade_no" json:"out_trade_no"`                           // 商户系统内部订单号，要求32个字符内（最少6个字符），只能是数字、大小写字母_-|*且在同一个商户号下唯一
	Attach             string `xml:"attach" json:"attach"`                                       // 附加数据，原样返回
	TimeEnd            string `xml:"time_end" json:"time_end"`                                   // 订单支付时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010
	TradeStateDesc     string `xml:"trade_state_desc" json:"trade_state_desc"`                   // 对当前查询订单状态的描述和下一步操作的指引
}

/* 关闭订单 */

// TradeCloseOrder 关闭订单 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_3
type TradeCloseOrder struct {
	AuxParam
	OutTradeNo string `xml:"out_trade_no" json:"out_trade_no"` // 商户系统内部订单号，要求32个字符内（最少6个字符），只能是数字、大小写字母_-|*且在同一个商户号下唯一
}

func (t TradeCloseOrder) ReturnType() string {
	return "xml"
}

// TradeCloseOrderRsp 关闭订单响应参数
type TradeCloseOrderRsp struct {
	PayError
	AppID     string `xml:"appid" json:"appid"`           // 微信分配的公众账号ID
	MchID     string `xml:"mch_id" json:"mch_id"`         // 微信支付分配的商户号
	NonceStr  string `xml:"nonce_str" json:"nonce_str"`   // 随机字符串，不长于32位
	Sign      string `xml:"sign" json:"sign"`             // 签名
	ResultMsg string `xml:"result_msg" json:"result_msg"` // 对业务结果的补充说明
}

/* 申请退款 */

// TradeRefund 申请退款 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_4
type TradeRefund struct {
	AuxParam
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`       // 商户系统内部订单号，要求32个字符内（最少6个字符），只能是数字、大小写字母_-|*且在同一个商户号下唯一。transaction_id、out_trade_no二选一，如果同时存在优先级：transaction_id > out_trade_no
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`   // 微信生成的订单号，在支付通知中有返回
	OutRefundNo   string `xml:"out_refund_no" json:"out_refund_no"`                         // 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	TotalFee      string `xml:"total_fee" json:"total_fee"`                                 // 订单总金额，单位为分，只能为整数
	RefundFee     string `xml:"refund_fee" json:"refund_fee"`                               // 退款总金额，订单总金额，单位为分，只能为整数
	RefundFeeType string `xml:"refund_fee_type,omitempty" json:"refund_fee_type,omitempty"` // 退款货币类型，需与支付一致，或者不填。符合ISO 4217标准的三位字母代码，默认人民币：CNY
	RefundDesc    string `xml:"refund_desc" json:"refund_desc"`                             // 退款原因，若商户传入，会在下发给用户的退款消息中体现退款原因
	RefundAccount string `xml:"refund_account,omitempty" json:"refund_account,omitempty"`   // 退款资金来源，仅针对老资金流商户使用 REFUND_SOURCE_UNSETTLED_FUNDS---未结算资金退款（默认使用未结算资金退款） REFUND_SOURCE_RECHARGE_FUNDS---可用余额退款
	NotifyUrl     string `xml:"notify_url,omitempty" json:"notify_url,omitempty"`           // 异步接收微信支付退款结果通知的回调地址，通知URL必须为外网可访问的url，不允许带参数 公网域名必须为https，如果是走专线接入，使用专线NAT IP或者私有回调域名可使用http 如果参数中传了notify_url，则商户平台上配置的回调地址将不会生效。
}

func (t TradeRefund) NeedTlsCert() bool {
	return true
}

func (t TradeRefund) ReturnType() string {
	return "xml"
}

// TradeRefundRsp 申请退款响应参数
type TradeRefundRsp struct {
	PayError
	AppID               string `xml:"appid" json:"appid"`                                   // 微信分配的公众账号ID
	MchID               string `xml:"mch_id" json:"mch_id"`                                 // 微信支付分配的商户号
	NonceStr            string `xml:"nonce_str" json:"nonce_str"`                           // 随机字符串，不长于32位
	Sign                string `xml:"sign" json:"sign"`                                     // 签名
	TransactionId       string `xml:"transaction_id" json:"transaction_id"`                 // 微信订单号
	OutTradeNo          string `xml:"out_trade_no" json:"out_trade_no"`                     // 商户系统内部订单号，要求32个字符内（最少6个字符），只能是数字、大小写字母_-|*且在同一个商户号下唯一。
	OutRefundNo         string `xml:"out_refund_no" json:"out_refund_no"`                   // 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	RefundId            string `xml:"refund_id" json:"refund_id"`                           // 微信退款单号
	RefundFee           int    `xml:"refund_fee" json:"refund_fee"`                         // 退款总金额，单位为分，可以做部分退款
	SettlementRefundFee int    `xml:"settlement_refund_fee" json:"settlement_refund_fee"`   // 应结退款金额，去掉非充值代金券退款金额后的退款金额，退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
	TotalFee            int    `xml:"total_fee" json:"total_fee"`                           // 订单总金额，单位为分，只能为整数
	SettlementTotalFee  int    `xml:"settlement_total_fee" json:"settlement_total_fee"`     // 应结订单金额，去掉非充值代金券金额后的订单总金额，应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	FeeType             string `xml:"fee_type" json:"fee_type"`                             // 标价币种，订单金额货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY
	CashFee             int    `xml:"cash_fee" json:"cash_fee"`                             // 现金支付金额，单位为分，只能为整数
	CashFeeType         string `xml:"cash_fee_type" json:"cash_fee_type"`                   // 现金支付币种，货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY
	CashRefundFee       int    `xml:"cash_refund_fee" json:"cash_refund_fee"`               // 现金退款金额，单位为分，只能为整数
	CouponType0         string `xml:"coupon_type_0,omitempty" json:"coupon_type_0"`         // 代金券类型，CASH--充值代金券 NO_CASH---非充值代金券 订单使用代金券时有返回（取值：CASH、NO_CASH）。$n为下标,从0开始编号，举例：coupon_type_0
	CouponRefundFee     int    `xml:"coupon_refund_fee,omitempty" json:"coupon_refund_fee"` // 代金券退款总金额，代金券退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为现金
	CouponRefundFee0    int    `xml:"coupon_refund_fee_0" json:"coupon_refund_fee_0"`       // 单个代金券退款金额，代金券退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为现金
	CouponRefundCount   int    `xml:"coupon_refund_count" json:"coupon_refund_count"`       // 退款代金券使用数量
	CouponRefundId0     string `xml:"coupon_refund_id_0" json:"coupon_refund_id_0"`         // 退款代金券ID, $n为下标，从0开始编号
}

/* 查询退款 */

// TradeRefundQuery 查询退款 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_5
type TradeRefundQuery struct {
	AuxParam
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"` // 微信订单号查询的优先级是： refund_id > out_refund_no > transaction_id > out_trade_no
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`     // 商户系统内部订单号，要求32个字符内（最少6个字符），只能是数字、大小写字母_-|*且在同一个商户号下唯一。
	OutRefundNo   string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty"`   // 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	RefundId      string `xml:"refund_id,omitempty" json:"refund_id,omitempty"`           // 微信生成的退款单号，在申请退款接口有返回
	Offset        string `xml:"offset,omitempty" json:"offset,omitempty"`                 // 偏移量，当部分退款次数超过10次时可使用，表示返回的查询结果从这个偏移量开始取记录
}

func (t TradeRefundQuery) ReturnType() string {
	return "xml"
}

// TradeRefundQueryRsp 查询退款响应参数
type TradeRefundQueryRsp struct {
	PayError
	AppID                string `xml:"appid" json:"appid"`                                               // 微信分配的公众账号ID（企业号corpid即为此appid）
	MchID                string `xml:"mch_id" json:"mch_id"`                                             // 微信支付分配的商户号
	NonceStr             string `xml:"nonce_str" json:"nonce_str"`                                       // 随机字符串，不长于32位
	Sign                 string `xml:"sign" json:"sign"`                                                 // 签名
	TotalRefundCount     int    `xml:"total_refund_count" json:"total_refund_count"`                     // 订单总共已发生的部分退款次数，当请求参数传入offset后有返回
	TransactionId        string `xml:"transaction_id" json:"transaction_id"`                             // 微信订单号
	OutTradeNo           string `xml:"out_trade_no" json:"out_trade_no"`                                 // 商户系统内部订单号，要求32个字符内（最少6个字符），只能是数字、大小写字母_-|*且在同一个商户号下唯一。
	TotalFee             int    `xml:"total_fee" json:"total_fee"`                                       // 订单总金额，单位为分，只能为整数
	SettlementTotalFee   int    `xml:"settlement_total_fee" json:"settlement_total_fee"`                 // 应结订单金额，当订单使用了免充值型优惠券后返回该参数，应结订单金额=订单金额-免充值优惠券金额。
	FeeType              string `xml:"fee_type" json:"fee_type"`                                         // 货币种类，订单金额货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY
	CashFee              int    `xml:"cash_fee" json:"cash_fee"`                                         // 现金支付金额，单位为分，只能为整数
	RefundCount          int    `xml:"refund_count" json:"refund_count"`                                 // 当前返回退款笔数
	OutRefundNo0         string `xml:"out_refund_no_0,omitempty" json:"out_refund_no_0"`                 // 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	RefundId0            string `xml:"refund_id_0,omitempty" json:"refund_id_0"`                         // 微信退款单号
	RefundChannel0       string `xml:"refund_channel_0,omitempty" json:"refund_channel_0"`               // 退款渠道 ORIGINAL—原路退款 BALANCE—退回到余额 OTHER_BALANCE—原账户异常退到其他余额账户 OTHER_BANKCARD—原银行卡异常退到其他银行卡
	RefundFee0           int    `xml:"refund_fee_0,omitempty" json:"refund_fee_0"`                       // 申请退款金额，退款总金额，单位为分，可以做部分退款
	RefundFee            int    `xml:"refund_fee" json:"refund_fee"`                                     // 退款总金额，各退款单的退款金额累加
	CouponRefundFee      int    `xml:"coupon_refund_fee" json:"coupon_refund_fee"`                       // 代金券退款总金额，各退款单的代金券退款金额累加
	SettlementRefundFee0 int    `xml:"settlement_refund_fee_0,omitempty" json:"settlement_refund_fee_0"` // 退款金额，退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
	CouponType00         string `xml:"coupon_type_0_0,omitempty" json:"coupon_type_00"`                  // 代金券类型，CASH--充值代金券 NO_CASH---非充值优惠券 开通免充值券功能，并且订单使用了优惠券后有返回（取值：CASH、NO_CASH）。$n为下标,$m为下标,从0开始编号，举例：coupon_type_$0_$1
	CouponRefundFee0     int    `xml:"coupon_refund_fee_0,omitempty" json:"coupon_refund_fee_0"`         // 总代金券退款金额，代金券退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为现金
	CouponRefundCount0   int    `xml:"coupon_refund_count_0,omitempty" json:"coupon_refund_count_0"`     // 退款代金券使用数量 ,$n为下标,从0开始编号
	CouponRefundId00     string `xml:"coupon_refund_id_0_0,omitempty" json:"coupon_refund_id_00"`        // 退款代金券ID, $n为下标，$m为下标，从0开始编号
	CouponRefundFee00    int    `xml:"coupon_refund_fee_0_0,omitempty" json:"coupon_refund_fee_00"`      // 单个退款代金券支付金额, $n为下标，$m为下标，从0开始编号
	RefundStatus0        string `xml:"refund_status_0,omitempty" json:"refund_status_0"`                 // 退款状态： SUCCESS—退款成功 REFUNDCLOSE—退款关闭，指商户发起退款失败的情况。 PROCESSING—退款处理中 CHANGE—退款异常，退款到银行发现用户的卡作废或者冻结了，导致原路退款银行卡失败，可前往商户平台（pay.weixin.qq.com）-交易中心，手动处理此笔退款。$n为下标，从0开始编号。
	RefundAccount0       string `xml:"refund_account_0,omitempty" json:"refund_account_0"`               // 退款资金来源，REFUND_SOURCE_RECHARGE_FUNDS---可用余额退款/基本账户 REFUND_SOURCE_UNSETTLED_FUNDS---未结算资金退款 $n为下标，从0开始编号
	RefundRecvAccout0    string `xml:"refund_recv_accout_0,omitempty" json:"refund_recv_accout_0"`       // 退款入账账户，取当前退款单的退款入账方 1）退回银行卡： {银行名称}{卡类型}{卡尾号} 2）退回支付用户零钱: 支付用户零钱 3）退还商户: 商户基本账户 商户结算银行账户 4）退回支付用户零钱通: 支付用户零钱通
	RefundSuccessTime0   string `xml:"refund_success_time_0,omitempty" json:"refund_success_time_0"`     // 退款成功时间，当退款状态为退款成功时有返回。$n为下标，从0开始编号。
	CashRefundFee        int    `xml:"cash_refund_fee" json:"cash_refund_fee"`                           // 用户退款金额，退款给用户的金额，不包含所有优惠券金额
}
