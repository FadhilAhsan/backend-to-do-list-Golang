package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Task struct {
	Id         int64        `json:"id" orm:"pk"`
	Task       string     	`json:"task"`
	CreatedAt  time.Time 	`json:"-" orm:"type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Task))
}

func (this *Task) TableName() string {
	return "task"
}
