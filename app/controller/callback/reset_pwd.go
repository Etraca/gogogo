package callback

import (
	"gogogo/app/param"
	"gogogo/app/service"
	"gogogo/app/common"
)

func ResetPwd(user param.ParamsResetPwd) (res interface{}) {
	var result common.Result

	if user.LogonName == "" {
		result.ErrorCode = common.ParamIsNotBlank
		result.ErrorMsg = "logon_name" + common.Error[result.ErrorCode]
		return result
	}
	if user.NewPassWord == "" {
		result.ErrorCode = common.ParamIsNotBlank
		result.ErrorMsg = "new_pass_word" + common.Error[result.ErrorCode]
		return result
	}
	if user.PassWord == "" {
		result.ErrorCode = common.ParamIsNotBlank
		result.ErrorMsg = "pass_word" + common.Error[result.ErrorCode]
		return result
	}

	result,_=service.ResetPwd(user)
	return  result
}
