package db_auth_server

import (
	"database/sql"
	"fmt"

	configl_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/config"
	domainl_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/domain"
	utils_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/utils"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *configl_auth_server.DataBase) (*gorm.DB, error) {
	// connectionString := "user=" + config.User + " password=" + config.UserPassword + " host=" + config.Host
	connectionString := "user=postgres password=8086 host=localhost"
	sql, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("--", err)
		return nil, err
	}

	rows, err := sql.Query("SELECT 1 FROM pg_database WHERE datname = '" + config.DBName + "'")
	if err != nil {
		fmt.Println("Error checking database existence:", err)
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("Database" + config.DBName + " already exists.")
	} else {
		// If the database does not exist, create it
		_, err = sql.Exec("CREATE DATABASE " + config.DBName)
		if err != nil {
			fmt.Println("Error creating database:", err)
		}
	}

	DB, err := gorm.Open(postgres.Open(config.DBConeectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(domainl_auth_server.Users{},
		domainl_auth_server.Admins{},
	)
	if err != nil {
		return nil, err
	}

	CheckAndCreateAdmin(DB)

	return DB, nil
}

func CheckAndCreateAdmin(DB *gorm.DB) {
	var count int
	var (
		Name     = "HyperHive"
		Email    = "hyperhive@gmail.com"
		Password = "TopSecret@878"
	)
	HashedPassword := utils_auth_server.HashPassword(Password)

	query := "SELECT COUNT(*) FROM admins"
	DB.Raw(query).Row().Scan(&count)
	if count <= 0 {
		query = "INSERT INTO admins(name, email, password) VALUES(?, ?, ?)"
		DB.Exec(query, Name, Email, HashedPassword).Row().Err()
	}
}
