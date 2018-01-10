package callback

import "gogogo/app/service"

func QueryAll() (res interface{}) {
	return service.QueryAll()
}

func Total() (res interface{}) {
	return service.Total()
}
