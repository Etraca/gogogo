package common

const (
	Success   = iota
	Fail
	LoginFail
	ParamIsNotBlank
	PassWordError
)

var Error = map[int]string{
	Success:   "成功",
	Fail:      "失败",
	LoginFail: "登录失败，请输入正确的用户名和密码",
	ParamIsNotBlank: " 不能为空",
	PassWordError: " 密码不正确",
}
