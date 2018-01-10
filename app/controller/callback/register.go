package callback

import (
	"gogogo/app/common"
	"gogogo/app/domain"
	"gogogo/app/model"
	"gogogo/app/utils"
)

func Register(user domain.User) (res interface{}) {
	var result common.Result

	if user.LogonName == "" {
		result.ErrorCode = common.ParamIsNotBlank
		result.ErrorMsg = "logon_name" + common.Error[result.ErrorCode]
		return result
	}
	if user.UserName == "" {
		result.ErrorCode = common.ParamIsNotBlank
		result.ErrorMsg = "user_name" + common.Error[result.ErrorCode]
		return result
	}
	if user.PassWord == "" {
		result.ErrorCode = common.ParamIsNotBlank
		result.ErrorMsg = "pass_word" + common.Error[result.ErrorCode]
		return result
	}

	//service.Register(user)

	var userModel  =new(model.EtracaUsers)
	userModel.UserName=user.UserName
	userModel.LogonName=user.LogonName
	userModel.PassWord=user.PassWord
	userModel.Status=1
	err:=userModel.Insert()
	utils.CkeckError(err,"insert user error")

	result.ErrorCode = common.Success
	result.ErrorMsg = common.Error[result.ErrorCode]
	result.Data = "hi," + user.UserName + ",终于等到你！"
	return result
}
