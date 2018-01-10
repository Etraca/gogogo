package param

type ParamsResetPwd struct {
	PassWord    string `json:"pass_word"`
	NewPassWord string `json:"new_pass_word""`
	LogonName   string `json:"logon_name"`
}
