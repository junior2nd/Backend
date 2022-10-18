package migration

import (
	"fmt"
	"log"
	"msa/database"
	"msa/model/entity"
)

func RunMigration(){
	err := database.DB.AutoMigrate(
		&entity.User{},
		&entity.Apartment{},
		&entity.Member{},
		&entity.Room{},
		&entity.Roomtype{},
		&entity.Transaction{},
	)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("data migrated")
}