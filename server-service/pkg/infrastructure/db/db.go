package db_server_service

import (
	"context"
	"database/sql"
	"fmt"

	config_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/config"
	domain_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/domain"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbInit(DbConfig config_server_service.DataBasePostgres, mongodb config_server_service.MongodDb) (*gorm.DB, *mongo.Collection, error) {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s sslmode=disable", DbConfig.User, DbConfig.UserPassword, DbConfig.Host)
	sql, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, nil, err
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
		return nil, nil, err
	}

	err = DB.AutoMigrate(
		domain_server_service.Server{},
		domain_server_service.ChannelCategory{},
		domain_server_service.ServerMembers{},
		domain_server_service.Channels{},
	)
	if err != nil {
		return nil, nil, err
	}

	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongodb.MongodbURL).SetServerAPIOptions(serverApi)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, nil, err
	}

	collection := client.Database(mongodb.DataBase).Collection(mongodb.ServerChatCollection)
	return DB, collection, nil
}
