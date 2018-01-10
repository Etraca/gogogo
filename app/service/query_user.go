package service

import (
	"gogogo/app/common"
	"gogogo/app/model"
)

func QueryAll() (res interface{}) {
	var result common.Result

	var user  =new(model.EtracaUsers)
	result.Data,_=user.FindAll()

	result.ErrorCode=common.Success
	result.ErrorMsg=common.Error[common.Success]
	return result
}

func Total() (res interface{}){
	var result common.Result

	var user  =new(model.EtracaUsers)
	result.Data,_=user.Total()

	result.ErrorCode=common.Success
	result.ErrorMsg=common.Error[common.Success]
	return result
}