package dao

import (
	"ALpay/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Initdb() error{
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/gopay?charset=utf8")
	Db.SetMaxIdleConns(100)
	Db.SetMaxOpenConns(200)
	if err != nil {
		log.Fatal("[Init] 初始化数据库失败")
		return err
	}

	if err := Db.Ping(); err != nil {
		return err
	}
	return nil
}

//表设计
//ID ip ProjectName out_trade_no Price CreateTime UpdataTime State
func InsertinventoryHistory(ProjectName, Price, CreateOrderTime,ip string, out_trade_no int64)  {
	insertSQL := `INSERT inventory_history (IP, ProjectName, Out_trade_no, Price, CreateTime, UpdataTime, State) VALUE (?, ?, ?, ?, ?, ?, ?) `
	database, err := GetDB().Begin()
	if err != nil {
		log.Fatal("[Sell]事务开启失败",err.Error())
		return
	}

	defer func() {
		if err != nil {
			_ = database.Rollback() //回滚
		}
	}()

	_, err = database.Exec(insertSQL,ip, ProjectName, out_trade_no, Price, CreateOrderTime, "nil", config.ConfirmStateIsflase)
	if err != nil {
		log.Fatalf("[Sell]增加库存记录失败：err %s ",err)
		return
	}

	_ = database.Commit()
}

func UpdateinventoryHistory(UpdataOrderTime string, out_trade_no int64) {
	updateSQL := `UPDATE inventory_history SET UpdataTime = ? AND State = ? WHERE Out_trade_no = ?;`
	database, err := GetDB().Begin()
	if err != nil {
		log.Fatal("[Sell]事务开启失败",err.Error())
		return
	}

	defer func() {
		if err != nil {
			_ = database.Rollback() //回滚
		}
	}()

	_, err = database.Exec(updateSQL, UpdataOrderTime, config.ConfirmStateIsflase, out_trade_no)
	if err != nil {
		log.Fatalf("[Sell]增加库存记录失败：err %s ",err)
		return
	}

	_ = database.Commit()
}

func Checkinventory_history() []InventoryHistory {
	P := InventoryHistory{}
	var datas []InventoryHistory
	database := GetDB()
	defer database.Close()
	rows, _ := database.Query("SELECT * FROM inventory_history")
	defer  rows.Close()
	//_ = data.Scan(&P.ID, &P.IP, &P.ProjectName, &P.OutTradeNo, &P.Price, &P.CreateTime, &P.UpdateTime, &P.State)
	for rows.Next() {
		err := rows.Scan(&P.ID, &P.IP, &P.ProjectName, &P.OutTradeNo, &P.Price, &P.CreateTime, &P.UpdateTime, &P.State)
		if err != nil {
			log.Fatal(err)
		}
		datas = append(datas,P)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return datas
}

//表设计
//Use Pwd CreateTime
func InsertUser(Use, Pwd, CreateTime string) {
	insertSQL := `INSERT User (Use, Pwd, CreateTime) VALUE (?, ?, ?) `
	database := GetDB()
	defer database.Close()
	_, err := database.Exec(insertSQL, Use, Pwd, CreateTime)
	if err != nil {
		log.Fatalf("[Sell]创建用户失败：err %s ",err)
		return
	}
}

func CheckUser(Use string) User {
	P := User{}
	database := GetDB()
	defer database.Close()
	data := database.QueryRow("SELECT * FROM User WHERE Use = ?", Use)
	_ = data.Scan(&P.Use, &P.Pwd, &P.CreateTime)
	return P
}
