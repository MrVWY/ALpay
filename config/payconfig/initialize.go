package payconfig

import (
	"github.com/iGoogle-ink/gopay/alipay"
)

func Inits()  {
	Client = alipay.NewClient(AppId, privateKey, false)
	Client.SetLocation(""). // 设置时区，不设置或出错均为默认服务器时间
		SetAliPayRootCertSN("").                  // 设置支付宝根证书SN，通过 alipay.GetRootCertSN() 获取
		SetAppCertSN("").                         // 设置应用公钥证书SN，通过 alipay.GetCertSN() 获取
		SetAliPayPublicCertSN("").                // 设置支付宝公钥证书SN，通过 alipay.GetCertSN() 获取
		SetCharset("utf-8").                    // 设置字符编码，不设置默认 utf-8
		SetSignType("RSA2").                    // 设置签名类型，不设置默认 RSA2
		SetReturnUrl(ReturnUrl). // 设置返回URL
		SetNotifyUrl(NotifiyUrl). // 设置异步通知URL NotifiyUrl
		SetAppAuthToken("").                      // 设置第三方应用授权
		SetAuthToken("")                          // 设置个人信息授权
}

func Init_bm(subject string, outTradeNo int64, quit_url string, total_amount string, product_code string)  {
	BodyMap.Set("subject", subject)           //商品的标题/交易标题/订单标题/订单关键字等。
	BodyMap.Set("out_trade_no", outTradeNo)   // 商户网站唯一订单号
	BodyMap.Set("quit_url", quit_url)         //用户付款中途退出返回商户网站的地址
	BodyMap.Set("total_amount", total_amount) //	订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	BodyMap.Set("product_code", product_code) //销售产品码，商家和支付宝签约的产品码
}