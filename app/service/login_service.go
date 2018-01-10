package service

import (
	"log"
	"gogogo/app/domain"
	"gogogo/app/model"
)

func Logon(logon domain.User) bool {
	var user = new(model.EtracaUsers)
	dbPassWord, err := user.FindPassWordByLogonName(logon.LogonName)
	if err != nil {
		return false
	}
	log.Println("password: ", dbPassWord)
	if logon.PassWord == dbPassWord {
		user.Logon()
		return true
	}
	return false
}
