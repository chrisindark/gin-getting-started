// models/setup.go

package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Print("Hello, ")
		fmt.Print("World!")
		fmt.Print(err)
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
