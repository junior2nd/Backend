package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	//?parseTime=true จัดการเรื่อง Time.time ใน model ได้ แก้ปัญหา uint8
	dsn := "root:0902708988@tcp(127.0.0.1:3306)/msa_db?parseTime=true&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect Database")
	}
	fmt.Println("conneted database")
}