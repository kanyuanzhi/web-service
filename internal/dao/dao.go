package dao

import (
	"github.com/kanyuanzhi/web-service/global"
	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

func New() *Dao {
	return &Dao{db: global.DB}
}
