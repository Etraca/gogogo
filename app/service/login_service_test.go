package service

import (
	"testing"
	"gogogo/app/domain"
	"gogogo/app/common"
	"gogogo/app/model"
	"encoding/json"
	"log"
)

func TestLogon(t *testing.T) {
	common.SetConfig()
	model.SetEngine()
	user := domain.User{LogonName: "wpf", PassWord: "123456"}
	if real := Logon(user); !real {
		t.Error()
	}
}

func TestTotal(t *testing.T) {
	common.SetConfig()
	model.SetEngine()

	result := Total()
	str, _ := json.Marshal(result)
	log.Println(str)
	ret := &common.Result{}
	json.Unmarshal([]byte(str), *ret)
	if ret.ErrorCode != 0 {
		t.Fail()
	}
}
