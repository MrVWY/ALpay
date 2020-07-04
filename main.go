package main

import (
	"ALpay/dao"
	"ALpay/middleware"
	"ALpay/service"
	"ALpay/service/model"
)

func main() {
	var err error
	err = dao.Init()
	if err != nil {
		panic(err)
	}
	middleware.C()
	model.InitModel()
	//queue.Init()
	router := service.RouterInit()
	_ = router.Run(":9000")
}
