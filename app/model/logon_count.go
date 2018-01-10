package model

import (
	"time"
)

type EtracaLogonCount struct {
	Id         int64     `xorm:"int(11) pk not null autoincr"`
	Count      int64     `xorm:"varchar(20) not null"`
	CreateTime time.Time `xorm:"datetime created"`
	UpdateTime time.Time `xorm:"datetime updated"`
	Version    int       `xorm:"version"`
}

func (this *EtracaLogonCount) Insert() (error) {
	_, err := engine.InsertOne(this)
	return err
}

func (this *EtracaLogonCount) Update() (error) {
	_, err := engine.Update(this)
	return err
}

func (this *EtracaLogonCount) Get() (*EtracaLogonCount, error) {
	_, err := engine.ID(1).Get(this)
	return this, err
}
