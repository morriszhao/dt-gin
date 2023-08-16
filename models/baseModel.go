package models

import "time"

type BaseModel struct {
	Id       int64     `xorm:"pk autoincr bigint(20)" json:"id,omitempty"`
	Createat time.Time `xorm:"datetime" json:"createat"`
}

// GetOneById 定义一些 数据库查询的 通用方法
func GetOneById(id int, obj interface{}) error {
	_, err := DbEngin.Where("id=?", id).Get(obj)
	return err
}
