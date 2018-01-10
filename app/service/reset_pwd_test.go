package service

import (
	"testing"
	"gogogo/app/param"
	"gogogo/app/common"
)

func TestResetPwd(t *testing.T) {
	user := &param.ParamsResetPwd{LogonName: "wpf", PassWord: "123", NewPassWord: "123456"}
	ret, err := ResetPwd(*user)
	if err != nil || ret.ErrorCode != common.Success {
		t.Error()
	}
}
