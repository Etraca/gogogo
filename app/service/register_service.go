package service

import (
	"gogogo/app/utils"
	"time"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"gogogo/app/domain"
)

func Register(register domain.User) {
	log.Println("register start,user:", register)

	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/etraca?charset=utf8")
	utils.CkeckError(err, "open db conn error: %v")
	defer db.Close()

	sql := "INSERT etraca_users(user_name, pass_word, logon_name, last_logon_time, status, create_time, update_time) VALUES (?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(sql)
	utils.CkeckError(err, "prepare stmt error: %v")
	defer stmt.Close()

	res, err := stmt.Exec(register.UserName, register.PassWord, register.LogonName, time.Now(), 1, time.Now(), time.Now())
	utils.CkeckError(err, "exec stmt error: %v")
	id, err := res.LastInsertId()
	utils.CkeckError(err, "get id error : %v")

	log.Println("register end,user :", register, ",new user id:", id)
}
