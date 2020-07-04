package dao

type User struct {
	Use string
	Pwd string
	CreateTime string
}

type InventoryHistory struct {
	ID          int
	IP          string
	ProjectName string
	OutTradeNo  int64
	Price       string
	CreateTime  string
	UpdateTime  string
	State       string
}