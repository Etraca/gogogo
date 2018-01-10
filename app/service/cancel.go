package service

import (
	"gogogo/app/domain"
	"log"
	"database/sql"
	"gogogo/app/utils"
	"gogogo/app/common"
)

func Cancel(user domain.User) (res interface{}) {
	var result common.Result

	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/etraca?charset=utf8")
	utils.CkeckError(err, "open db conn error: %v")
	defer db.Close()

	/**
		validate old pwd
	 */
	rows, err := db.Query("SELECT pass_word FROM etraca_users WHERE logon_name=?", user.LogonName)
	defer rows.Close()
	utils.CkeckError(err, "db query error :%v")
	var dbPassWord string
	for rows.Next() {
		err := rows.Scan(&dbPassWord)
		utils.CkeckError(err, "db scan error: %v")
	}
	log.Println("password: ", dbPassWord)
	if user.PassWord != dbPassWord {
		result.ErrorCode = common.PassWordError
		result.ErrorMsg = common.Error[result.ErrorCode]
		return result
	}

	sql := "DELETE FROM etraca_users WHERE logon_name=?"
	stmt, err := db.Prepare(sql)
	utils.CkeckError(err, "prepare stmt error: %v")
	defer stmt.Close()

	_, err = stmt.Exec(user.LogonName)
	utils.CkeckError(err, "exec stmt error: %v")

	result.ErrorCode = common.Success
	result.ErrorMsg = common.Error[result.ErrorCode]
	return result
}
