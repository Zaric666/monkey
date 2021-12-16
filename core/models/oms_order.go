package models

import (
	"gorm.io/gorm"
)

type OmsOrder struct {
	Id        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderId   int    `json:"orderId" gorm:"int:16"`
	OrderCode string `json:"orderCode" gorm:"size:255;"`
}

func (OmsOrder) TableName() string {
	return "oms_orders"
}

func (e *OmsOrder) GetId() interface{} {
	return e.Id
}

func (e *OmsOrder) GetList(tx *gorm.DB, list interface{}) (err error) {
	return tx.Table(e.TableName()).Where("status = ?", 2).Find(list).Error
}
