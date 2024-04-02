package db_friend_server

import (
	"database/sql"
	"fmt"

	config_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/config"
	domain_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(db *config_friend_server.DataBase) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("name= %s password=%s host= %s sslmode=disable", db.DBName, db.UserPassword, db.Host)
	sql, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	rows, err := sql.Query("SELECT 1 FROM pg_database WHERE  datname= '" + db.DBName + "'")
	if err != nil {
		fmt.Println("Error checking database existence:", err)
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("Database" + db.DBName + " already exists.")
	} else {
		// If the database does not exist, create it
		_, err = sql.Exec("CREATE DATABASE " + db.DBName)
		if err != nil {
			fmt.Println("Error creating database:", err)
		}
	}

	DB, err := gorm.Open(postgres.Open(db.DBConeectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err= DB.AutoMigrate(&domain_friend_server.Friends{})
	if err!=nil{
		return nil, err
	}

	return DB, nil
}
