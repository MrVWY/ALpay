package payconfig

import (
	"github.com/iGoogle-ink/gopay/alipay"
	"log"
)

func TradePagePay_Send(client *alipay.Client) string {
	payUrl , err := client.TradePagePay(BodyMap)
	if err != nil {
		log.Fatal(err)
	}

	return payUrl
}

func Verity(request *alipay.NotifyRequest) string {
	ok , err := alipay.VerifySign(Aipaykey,request)
	if !ok || err != nil{
		return "error"
	}
	return "ok"
}
