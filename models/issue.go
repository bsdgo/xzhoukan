package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Issue struct {
	Id          int64
	Title       string
	SubTitle    string
	IssueNum    int16
	PublishDate time.Time `orm:"type(datetime);index"`
	Desc        string
}

func (m *Issue) TableName() string {
	return "issues"
}

func (m *Issue) Add() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
