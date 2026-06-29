package infra

import (
	"ecommerce/config"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cnf *config.Config)(*gorm.DB,error) {



}