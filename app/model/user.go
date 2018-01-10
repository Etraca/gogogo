package model

import (
	"time"
	"errors"
	"log"
)

type EtracaUsers struct {
	Id            int64      `xorm:"int(11) pk not null autoincr"`
	UserName      string     `xorm:"varchar(20) not null"`
	PassWord      string     `xorm:"varchar(20) not null"`
	LogonName     string     `xorm:"varchar(20) not null unique"`
	LastLogonTime SimpleTime `xorm:"datetime updated"`
	Status        int        `xorm:"int"`
	CreateTime    time.Time  `xorm:"datetime created"`
	UpdateTime    time.Time  `xorm:"datetime updated"`
	Page                     `xorm:"-"`
}

func (this *EtracaUsers) Insert() error {
	_, err := engine.InsertOne(this)
	return err
}

func (this *EtracaUsers) Exist() (bool, error) {
	return engine.Get(this)
}

func (this *EtracaUsers) FindUserById(id int64) (*EtracaUsers, error) {
	user := &EtracaUsers{Id: id}
	_, err := engine.Get(user)
	return user, err
}

func (this *EtracaUsers) FindUserByLogonName(logonName string) (*EtracaUsers, error) {
	user := &EtracaUsers{LogonName: logonName}
	_, err := engine.Get(user)
	return user, err
}

func (this *EtracaUsers) Delete() error {
	_, err := engine.Delete(this)
	return err
}

func (this *EtracaUsers) Update() error {
	_, err := engine.ID(this.Id).Update(this)
	return err
}

func (this *EtracaUsers) FindAll() ([]EtracaUsers, error) {
	var users []EtracaUsers
	err := engine.Desc("id").Find(&users)
	return users, err
}

func (this *EtracaUsers) Find() ([]EtracaUsers, error) {
	var users []EtracaUsers
	err := engine.Find(users, this)
	return users, err
}
func (this *EtracaUsers) FindPassWordByLogonName(logonName string) (string, error) {
	this.LogonName = logonName
	var passWord string
	has, err := engine.Get(this)
	if has && err == nil {
		passWord = this.PassWord
		return passWord, err
	}
	err = errors.New("cannot find The user")
	return passWord, err
}

func (this *EtracaUsers) Total() (int64, error) {
	return engine.Count(this)
}

func (this *EtracaUsers) Logon() (error) {
	session := engine.NewSession()
	defer session.Close()
	session.Begin()
	_, err := session.ID(this.Id).Update(this)
	if err != nil {
		log.Println(err)
		session.Rollback()
		panic(err)
	}
	count := &EtracaLogonCount{Id: 1}
	count, err = count.Get()
	count.Count++
	_, err = session.ID(count.Id).Update(count)
	if err != nil {
		session.Rollback()
		panic(err)
	}
	session.Commit()
	return err
}
