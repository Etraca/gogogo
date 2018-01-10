package model

import (
	"testing"
	"gogogo/app/common"
)

func TestEtracaUsers_Logon(t *testing.T) {
	common.SetConfig()
	SetEngine()
	user := &EtracaUsers{LogonName: "wpf", PassWord: "123456"}
	if user.Logon() != nil {
		t.Error()
	}
}

func TestEtracaUsers_Total(t *testing.T) {
	common.SetConfig()
	SetEngine()
	count, err := new(EtracaUsers).Total()
	if err != nil || count < 0 {
		t.Fail()
	}
	t.Log("注册用户个数：",count)
}
