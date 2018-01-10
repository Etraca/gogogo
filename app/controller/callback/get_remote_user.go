package callback

import "gogogo/app/remote"

func GetRemoteUser() (res interface{}) {

	user := remote.GetUser()
	user.Insert()
	return user
}

func PostRemoteUser(p string) (res interface{}) {
	return remote.PostUser(p)
}
