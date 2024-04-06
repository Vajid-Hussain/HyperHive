package db_friend_server

import (
	"context"
	"database/sql"
	"fmt"

	config_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/config"
	domain_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/domain"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MongoDBCollections struct {
	FriendChatCollection *mongo.Collection
}

func InitDB(db *config_friend_server.DataBase, mongoDB *config_friend_server.MongodDb) (*gorm.DB, *MongoDBCollections, error) {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s sslmode=disable", db.User, db.UserPassword, db.Host)
	sql, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, nil, err
	}

	rows, err := sql.Query("SELECT 1 FROM pg_database WHERE datname = '" + db.DBName + "'")
	if err != nil {
		fmt.Println("Error checking database existence:", err)
	}
	defer rows.Close()

	if rows != nil && rows.Next() {
		fmt.Println("Database" + db.DBName + " already exists.")
	} else {
		_, err = sql.Exec("CREATE DATABASE " + db.DBName)
		if err != nil {
			fmt.Println("Error creating database:", err)
		}
	}

	DB, err := gorm.Open(postgres.Open(db.DBConeectionString), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	err = DB.AutoMigrate(&domain_friend_server.Friends{})
	if err != nil {
		return nil, nil, err
	}

	// ----------mongdob

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoDB.MongodbURL).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, nil, err
	}

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	var result bson.M
	if err := client.Database("message").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return nil, nil, err
	}
	fmt.Println("-mongdob", result)

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("ping")
	} else {
		fmt.Println("connected to mongodb")
	}

	var mongoDBCollection MongoDBCollections
	mongoDBCollection.FriendChatCollection = client.Database("hyperhive").Collection(mongoDB.FriendChatCollection)

	return DB, &mongoDBCollection, nil
}
