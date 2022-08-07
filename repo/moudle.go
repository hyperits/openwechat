package repo

import "time"

type Msg struct {
	Id       int       `gorm:"column:id" db:"id" json:"id" form:"id"` //主键
	Sender   string    `gorm:"column:sender" db:"sender" json:"sender" form:"sender"`
	IsDone   int       `gorm:"column:is_done" db:"is_done" json:"is_done" form:"is_done"`
	Msg      string    `gorm:"column:msg" db:"msg" json:"msg" form:"msg"`
	CreateAt time.Time `gorm:"column:create_at" db:"create_at" json:"create_at" form:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at" db:"update_at" json:"update_at" form:"update_at"`
}

type Nic struct {
	Nic string `gorm:"column:nic" db:"nic" json:"nic" form:"nic"`
}
