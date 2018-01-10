package model

import (
	"gogogo/app/common"
	"gogogo/app/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"fmt"
)
var engine *xorm.Engine

func SetEngine() *xorm.Engine {
	var err error
	server := common.Cfg.MustValue("db", "server", "127.0.0.1")
	username := common.Cfg.MustValue("db", "username", "root")
	password := common.Cfg.MustValue("db", "password", "12345678")
	dbName := common.Cfg.MustValue("db", "db_name", "etraca")
	engine, err = xorm.NewEngine("mysql", username+":"+password+"@tcp("+server+":3306)/"+dbName+"?charset=utf8")
	utils.CkeckError(err, "engin set error :%v")
	engine.TZLocation = time.Local
	engine.ShowSQL(common.Cfg.MustBool("db", "show_sql", false))
	return engine
}

type SimpleTime time.Time

func (this SimpleTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
