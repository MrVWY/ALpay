package dao

import (
	"database/sql"
	"errors"
	"sync"
)

var  (
	Db      *sql.DB
	inlined bool
	m       sync.RWMutex
)

func Init() error  {
	m.Lock()
	defer m.Unlock()

	if inlined {
		return errors.New("[Init] dao 已经初始化")
	}

	err := Initdb()
	if err != nil {
		return err
	}

	inlined = true
	return nil
}

func GetDB() *sql.DB {
	return Db
}