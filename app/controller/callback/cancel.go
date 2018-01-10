package callback

import (
	"gogogo/app/domain"
	"gogogo/app/service"
	"gogogo/app/common"
)

func Cancel(user domain.User) (res interface{}) {
	var result common.Result

	if user.LogonName == "" {
		result.ErrorCode = common.ParamIsNotBlank
		result.ErrorMsg = "logon_name" + common.Error[result.ErrorCode]
		return result
	}
	if user.PassWord == "" {
		result.ErrorCode = common.ParamIsNotBlank
		result.ErrorMsg = "pass_word" + common.Error[result.ErrorCode]
		return result
	}
	return service.Cancel(user)
}
