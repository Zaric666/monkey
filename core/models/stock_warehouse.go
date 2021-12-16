package models

import (
	"gorm.io/gorm"
)

type StockWarehouse struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"size:255"`
}

func (StockWarehouse) TableName() string {
	return "stock_warehouse"
}

func (e *StockWarehouse) GetId() interface{} {
	return e.Id
}

func (e *StockWarehouse) GetList(tx *gorm.DB, list interface{}) (err error) {
	return tx.Table(e.TableName()).Where("status = ?", 1).Find(list).Error
}
