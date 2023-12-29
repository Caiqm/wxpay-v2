package wxpay

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

var (
	ErrWxNullParams         = errors.New("wxpay: param is null")
	ErrWxReturnCodeNotFound = errors.New("wxpay: return_code not found")
	ErrWxPemKeyNotFound     = errors.New("wxpay: wxpay pem or key cert not found")
)

type Client struct {
	appId          string
	secret         string
	mchId          string
	mchSecret      string
	host           string
	signType       string
	pemCert        []byte
	keyCert        []byte
	location       *time.Location
	client         *http.Client
	onReceivedData func(method string, data []byte)
}

type OptionFunc func(c *Client)

// 设置请求链接
func WithApiHost(host string) OptionFunc {
	return func(c *Client) {
		if host != "" {
			c.host = host
		}
	}
}

// 设置小程序登录链接
func WithJsCodeHost() OptionFunc {
	return func(c *Client) {
		c.host = "https://api.weixin.qq.com/sns/jscode2session"
	}
}

// 设置支付请求链接
func WithPayHost() OptionFunc {
	return func(c *Client) {
		c.host = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	}
}

// 设置申请退款请求链接
func WithRefundHost() OptionFunc {
	return func(c *Client) {
		c.host = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	}
}

// 设置商户号信息
func WithMchInformation(id, secret string) OptionFunc {
	return func(c *Client) {
		if id != "" {
			c.mchId = id
		}
		if secret != "" {
			c.mchSecret = secret
		}
		c.signType = "MD5"
	}
}

// 初始化
func New(appId, secret string, opts ...OptionFunc) (nClient *Client, err error) {
	if appId == "" || secret == "" {
		return nil, ErrWxNullParams
	}
	nClient = &Client{}
	nClient.appId = appId
	nClient.secret = secret
	nClient.client = http.DefaultClient
	nClient.location = time.Local
	nClient.LoadOptionFunc(opts...)
	return
}

// 加载接口链接
func (c *Client) LoadOptionFunc(opts ...OptionFunc) {
	for _, opt := range opts {
		if opt != nil {
			opt(c)
		}
	}
}

// 加载证书文件
func (c *Client) LoadAppCertPemKeyFromFile(pemCertPath, keyCertPath string) (err error) {
	p, err := os.ReadFile(pemCertPath)
	if err != nil {
		err = fmt.Errorf("wxpay: read pem cert fail, err = %s", err.Error())
		return
	}
	k, err := os.ReadFile(keyCertPath)
	if err != nil {
		err = fmt.Errorf("wxpay: read key cert fail, %s", err.Error())
		return
	}
	c.pemCert = p
	c.keyCert = k
	return
}

// 加载证书tls配置
func (c *Client) LoadTlsCertConfig() (tlsConfig *tls.Config, err error) {
	if len(c.pemCert) == 0 || len(c.keyCert) == 0 {
		err = ErrWxPemKeyNotFound
		return
	}
	var certificate tls.Certificate
	// 解析证书内容
	if certificate, err = tls.X509KeyPair(c.pemCert, c.keyCert); err != nil {
		err = fmt.Errorf("wxpay: tls.LoadX509KeyPair, %s", err.Error())
		return
	}
	tlsConfig = &tls.Config{
		Certificates:       []tls.Certificate{certificate},
		InsecureSkipVerify: true,
	}
	return
}

// 请求参数
func (c *Client) URLValues(param Param) (value url.Values, err error) {
	var values = url.Values{}
	// 是否需要APPID
	if param.NeedAppId() {
		values.Add(kFieldAppId, c.appId)
	}
	// 是否需要密钥
	if param.NeedSecret() {
		values.Add(kFieldSecret, c.secret)
	}
	var params = c.structToMap(param)
	for k, v := range params {
		if v == "" {
			continue
		}
		values.Add(k, v)
	}
	// 判断是否需要签名
	if param.NeedSign() {
		values.Add(kFieldMchId, c.mchId)
		values.Add(kFieldNonceStr, c.createNonceStr())
		values.Add(kFieldSignType, c.signType)
		signature := c.sign(values)
		// 添加签名
		values.Add(kFieldSign, signature)
	}
	return values, nil
}

// 结构体转map
func (c *Client) structToMap(stu interface{}) map[string]string {
	// 结构体转map
	m, _ := json.Marshal(&stu)
	var parameters map[string]string
	_ = json.Unmarshal(m, &parameters)
	return parameters
}

// 生成签名
func (c *Client) sign(parameters url.Values) string {
	signStr := c.formatBizQueryParaMap(parameters)
	signStr = fmt.Sprintf("%s&key=%s", signStr, c.mchSecret)
	h := md5.New()
	h.Write([]byte(signStr))
	sign := hex.EncodeToString(h.Sum(nil))
	return strings.ToUpper(sign)
}

// 生成小程序签名
func (c *Client) createAppletPaySign(timestamp, prepayId, nonceStr string) string {
	wxPayInfo := make(map[string]string, 5)
	wxPayInfo["appId"] = c.appId
	wxPayInfo["timeStamp"] = timestamp
	wxPayInfo["nonceStr"] = nonceStr
	wxPayInfo["package"] = fmt.Sprintf("prepay_id=%s", prepayId)
	wxPayInfo["signType"] = "MD5"
	signStr := c.formatQueryParaMap(wxPayInfo)
	signStr = fmt.Sprintf("%s&key=%s", signStr, c.mchSecret)
	h := md5.New()
	h.Write([]byte(signStr))
	sign := hex.EncodeToString(h.Sum(nil))
	return strings.ToUpper(sign)
}

// 格式化参数，签名过程需要使用
func (c *Client) formatQueryParaMap(parameters map[string]string) string {
	// 将key值提取出来
	var strs []string
	for k := range parameters {
		strs = append(strs, k)
	}
	// 排序
	sort.Strings(strs)
	// 赋值新map
	var signStr string
	var dot string
	for _, k := range strs {
		if parameters[k] == "" {
			continue
		}
		signStr += dot + k + "=" + parameters[k]
		dot = "&"
	}
	return signStr
}

// 格式化参数，签名过程需要使用
func (c *Client) formatBizQueryParaMap(parameters url.Values) string {
	// 将key值提取出来
	var strs []string
	for k := range parameters {
		strs = append(strs, k)
	}
	// 排序
	sort.Strings(strs)
	// 赋值新map
	var signStr string
	var dot string
	for _, k := range strs {
		if parameters[k][0] == "" || k == kFieldSign {
			continue
		}
		signStr += dot + k + "=" + parameters[k][0]
		dot = "&"
	}
	return signStr
}

// 格式化参数，将url.values转map
func (c *Client) formatUrlValueToMap(parameters url.Values) map[string]string {
	// 将key值提取出来
	var strs []string
	for k := range parameters {
		strs = append(strs, k)
	}
	// 排序
	sort.Strings(strs)
	// 赋值新map
	m := make(map[string]string)
	for _, k := range strs {
		m[k] = parameters[k][0]
	}
	return m
}

// 产生随机字符串，不长于32位
func (c *Client) createNonceStr() string {
	length := 32
	strByte := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	bytes := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		bytes[i] = strByte[r.Intn(len(strByte))]
	}
	return string(bytes)
}

type payXml map[string]string

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

// 重写加密xml
func (m payXml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	for k, v := range m {
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}
	return e.EncodeToken(start.End())
}

// 重写加解密xml
func (m *payXml) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = payXml{}
	for {
		var e xmlMapEntry
		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}

// 请求主方法
func (c *Client) doRequest(method string, param Param, result interface{}) (err error) {
	// 创建一个请求
	req, _ := http.NewRequest(method, c.host, nil)
	// 判断参数是否为空
	if param != nil {
		var values url.Values
		values, err = c.URLValues(param)
		if err != nil {
			return err
		}
		if method == http.MethodPost {
			// 根据类型转换
			if strings.ToLower(param.ReturnType()) == "json" {
				req.PostForm = values
			} else if param.ReturnType() == "jsonStr" {
				// 结构体转map
				var reqByte []byte
				if reqByte, err = json.Marshal(param); err != nil {
					return
				}
				bodyBuffer := bytes.NewBuffer(reqByte)
				req.Body = io.NopCloser(bodyBuffer)
				req.ContentLength = int64(bodyBuffer.Len())
			} else {
				var reqByte []byte
				mapValues := c.formatUrlValueToMap(values)
				if reqByte, err = xml.Marshal(payXml(mapValues)); err != nil {
					return
				}
				req.Body = io.NopCloser(bytes.NewBuffer(reqByte))
			}
		} else if method == http.MethodGet {
			req.URL, _ = url.Parse(c.host + "?" + values.Encode())
		}
	}
	// 是否需要证书
	if param.NeedTlsCert() {
		tlsConfig, err1 := c.LoadTlsCertConfig()
		if err1 != nil {
			err = err1
			return
		}
		c.client.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	}
	// 添加header头
	if param.ReturnType() == "jsonStr" {
		req.Header.Set("Content-Type", kContentTypeJson)
	} else {
		req.Header.Set("Content-Type", kContentType)
	}
	// 发起请求数据
	rsp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return fmt.Errorf("解析返回数据失败: %v", err)
	}
	err = c.decode(bodyBytes, method, param.ReturnType(), param.NeedVerify(), result)
	return
}

// 解密返回数据
func (c *Client) decode(data []byte, method, returnType string, needVerifySign bool, result interface{}) (err error) {
	// 返回结果
	if c.onReceivedData != nil {
		c.onReceivedData(method, data)
	}
	if strings.ToLower(returnType) == "json" || returnType == "jsonStr" || returnType == "" {
		var raw = make(map[string]json.RawMessage)
		if err = json.Unmarshal(data, &raw); err != nil {
			if returnType == "jsonStr" {
				err = nil
				rsp := result.(*QrcodeRsp)
				rsp.Buffer = data
				return
			}
			return fmt.Errorf("解析返回结构失败，%v", err)
		}
		// 判断是否成功
		var errNBytes = raw[kFieldErrCode]
		if len(errNBytes) > 0 && string(errNBytes) != "0" {
			var aErr *AppletError
			if err = json.Unmarshal(data, &aErr); err != nil {
				return
			}
			return aErr
		}
		if err = json.Unmarshal(data, result); err != nil {
			return
		}
	} else {
		var pErr PayError
		if err = xml.Unmarshal(data, &pErr); err != nil {
			return
		}
		if pErr.IsFailure() {
			return pErr
		}
		resultMap := make(map[string]string)
		if err = xml.Unmarshal(data, (*payXml)(&resultMap)); err != nil {
			return
		}
		// 校验签名
		if needVerifySign {
			params := make(url.Values)
			for key, value := range resultMap {
				strValue := fmt.Sprintf("%v", value)
				params.Add(key, strValue)
			}
			// 验证签名
			if err = c.VerifySign(params); err != nil {
				return
			}
		}
		// 参数绑定
		if err = xml.Unmarshal(data, result); err != nil {
			return
		}
	}
	return
}

// 返回内容
func (c *Client) OnReceivedData(fn func(method string, data []byte)) {
	c.onReceivedData = fn
}

// 验证签名
func (c *Client) VerifySign(values url.Values) (err error) {
	verifier := values.Get(kFieldSign)
	compareSign := c.sign(values)
	if strings.Compare(verifier, compareSign) != 0 {
		err = fmt.Errorf("验证签名失败，接口返回签名：%s，生成签名：%s", compareSign, verifier)
		return
	}
	return
}
