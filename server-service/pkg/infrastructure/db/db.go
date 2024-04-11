package db_server_service

import (
	"database/sql"
	"fmt"

	config_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/config"
	domain_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/domain"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbInit(DbConfig config_server_service.DataBasePostgres) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s sslmode=disable", DbConfig.DBName, DbConfig.UserPassword, DbConfig.Host)
	sql, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	rows, err := sql.Query("SELECT 1 FROM pg_database WHERE datname= '" + DbConfig.DBName + "'")
	if err != nil {
		fmt.Println("Error checking database existence:", err)
	}
	defer rows.Close()

	if rows != nil && rows.Next() {
		fmt.Println("Database" + DbConfig.DBName + " already exists.")
	} else {
		_, err = sql.Exec("CREATE DATABASE " + DbConfig.DBName)
		if err != nil {
			fmt.Println("Error creating database:", err)
		}
	}

	DB, err := gorm.Open(postgres.Open(DbConfig.DBConeectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(domain_server_service.Server{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}
