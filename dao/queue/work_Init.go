package queue

import (
	"encoding/json"
	"fmt"
	"log"
)

// 定义生产者接口
type Producer interface {
	MsgContent() []byte
}

// 定义接收者接口
type Receiver interface {
	Consumer([]byte)    error
}

type InventoryHistory struct {
	ProjectName     string `json:"project_name"`
	Price           string `json:"price"`
	CreateOrderTime string `json:"create_order_time"`
	Ip              string `json:"ip"`
	OutTradeNo      int64  `json:"out_trade_no"`
}

func (u *InventoryHistory) MsgContent() []byte {
	b, err := json.Marshal(u)
	if err != nil {
		log.Fatal("InventoryHistory err :", err)
	}
	return b
}

func (u *InventoryHistory) Consumer(b []byte) error {
	i := InventoryHistory{}
	err := json.Unmarshal(b, &i)
	if err != nil {
		//log.Fatal("Consumer:",err)
		return err
	}
	fmt.Println("receive msg is : ", i)
	return nil
}