package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Document struct {
	gorm.Model
	TicketId  string `gorm:"type:varchar(100);unique_index"`
	Content   string `gorm:"type:varchar(100)"`
	Title     string `gorm:"type:varchar(100)"`
	Author    string `gorm:"type:varchar(100)"`
	Topic     string `gorm:"type:varchar(100)"`
	Watermakr string `gorm:"type:varchar(100)"`
}

func Init(dialect, host, port, user, dbname, pass string) (*gorm.DB, error) {

	db, err := gorm.Open(dialect, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbname, pass))

	if err != nil {
		return nil, err
	}

	// if err := db.AutoMigrate(&Document{}).Error; err != nil {
	// 	return nil, err
	// }

	return db, nil
}
