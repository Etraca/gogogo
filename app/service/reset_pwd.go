package service

import (
	"database/sql"
	"gogogo/app/utils"
	"gogogo/app/param"
	"log"
	"gogogo/app/common"
)

func ResetPwd(user param.ParamsResetPwd) (common.Result,error) {
	var result common.Result

	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/etraca?charset=utf8")
	utils.CkeckError(err, "open db conn error: %v")
	defer db.Close()
	//开启事务
	tx,_:=db.Begin()

	/**
		validate old pwd
	 */
	rows, err := tx.Query("SELECT pass_word FROM etraca_users WHERE logon_name=?", user.LogonName)
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
		return result,err
	}

	/**
		modify pwd
	 */
	sql := "UPDATE etraca_users SET pass_word =? WHERE logon_name=?"
	stmt, err := tx.Prepare(sql)
	utils.CkeckError(err, "prepare stmt error: %v")
	defer stmt.Close()

	_, err = stmt.Exec(user.NewPassWord, user.LogonName)
	if err != nil {
		//回滚事务
		tx.Rollback()
		utils.CkeckError(err, "exec stmt error: %v")
	}
	//提交事务
	tx.Commit()

	result.ErrorCode = common.Success
	result.ErrorMsg = common.Error[result.ErrorCode]
	return result,err
}
