package domain

import "time"

type User struct {
	Id            uint64    `json:"id"`
	UserName      string    `json:"user_name"`
	PassWord      string    `json:"pass_word"`
	LogonName     string    `json:"logon_name"`
	LastLogonTime time.Time `json:"last_logon_time"`
	Status        int       `json:"status"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
}
