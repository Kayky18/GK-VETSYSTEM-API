package config

import (
	"os"

	"github.com/Kayky18/GK_API/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating..")

		//Create the database path
		err = os.Mkdir("./db", os.ModePerm)
		if err != nil {
			logger.Error("ERRO CREATING DATABASE PATH")
			return nil, err
		}

		//Creating the database file
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("ERRO CREATING DATABASE FILE: %v", err)
			return nil, err
		}

		//Close the file
		file.Close()
	}

	// Open connection to SQLite database
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("ERRO TO OPEN DATABASE: %v", err)
	}
	//Migrate Data
	err = db.AutoMigrate(&schemas.Patients{})
	if err != nil {
		logger.Errorf("ERROR TO MIGRATE DATABASE: %v", err)
	}

	return db, nil
}
