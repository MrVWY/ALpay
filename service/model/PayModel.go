package model

import (
	"ALpay/config/payconfig"
	"ALpay/dao"
	//"Ailpay/dao/queue"
	"github.com/gin-gonic/gin"
	"github.com/iGoogle-ink/gopay/alipay"
	"net/http"
	"strconv"
	"time"
	"log"
)

var timeTemplate1 = "2006-01-02 15:04:05"
func MakeMilkyOrder(money string, c *gin.Context) string {
	outTradeNo := time.Now().Unix()
	times := time.Unix(outTradeNo, 0).Format(timeTemplate1)
	ip := c.Request.RemoteAddr

	//t := &queue.InventoryHistory{
	//	ProjectName:     "请婕纶喝奶茶",
	//	Price:           money,
	//	CreateOrderTime: times,
	//	Ip:              ip,
	//	OutTradeNo:      outTradeNo,
	//}
	//a := &queue.QueueExchange{
	//	"hello",
	//	"info",
	//	"exchange1",
	//	"direct",
	//}
	//rabbit := queue.New(a)
	//rabbit.ListenProducer(t)

	payconfig.Init_bm("请婕纶喝奶茶", outTradeNo, "", money, "QUICK_WAP_WAY")
	payUrl := payconfig.TradePagePay_Send(payconfig.Client)

	dao.InsertinventoryHistory("请婕纶喝奶茶", money, times, ip, outTradeNo)

	return payUrl
}

func MakeOrder1Handler(c *gin.Context)  {
	payUrl := MakeMilkyOrder("1.0", c)
	c.Redirect(http.StatusMovedPermanently,payUrl)
}

func MakeOrder10Handler(c *gin.Context)  {
	payUrl := MakeMilkyOrder("10.0", c)
	c.Redirect(http.StatusMovedPermanently,payUrl)
}

func MakeOrder100Handler(c *gin.Context)  {
	payUrl := MakeMilkyOrder("100.0", c)
	c.Redirect(http.StatusMovedPermanently,payUrl)
}

func Verity(c *gin.Context)  {
	notifyReq, err := alipay.ParseNotifyResult(c.Request)
	log.Println(c.Request.URL)
	if err != nil {
		c.JSON(404,gin.H{"status":"Verity err"})
	}
	// 验签操作
	ok, _ := alipay.VerifySign(payconfig.Aipaykey, notifyReq)
	log.Println("123")
	log.Println(ok)
	if ok && notifyReq != nil  {
		times, _ := strconv.ParseInt(notifyReq.NotifyTime, 10, 64)
		updatatime := time.Unix(times, 0).Format(timeTemplate1)
		out_trade_nos, _ := strconv.ParseInt(notifyReq.OutTradeNo, 10, 64)
		log.Println(notifyReq.NotifyTime,notifyReq.OutTradeNo, updatatime, out_trade_nos)
		dao.UpdateinventoryHistory(notifyReq.NotifyTime,out_trade_nos)
	}
	c.String(http.StatusOK, "success")
}

func CompleteOrder(c *gin.Context)  {
	c.HTML(200,"returns.html",gin.H{})
}
