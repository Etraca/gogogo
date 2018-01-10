package callback

import (
	"gogogo/app/common"
	"gogogo/app/service"
	"gogogo/app/domain"
)

/**
	登录
 */
func Logon(logon domain.User) (res interface{}) {
	var result common.Result
	if service.Logon(logon) {
		result.ErrorCode = common.Success
		result.ErrorMsg = common.Error[result.ErrorCode]
		result.Data = "hi," + logon.LogonName + ",欢迎回来！"
		return result
	}
	result.ErrorCode = common.Fail
	result.ErrorMsg = common.Error[result.ErrorCode]
	return result
}
