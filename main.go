package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/huyhoangvp002/Delivery_test/api"
	db "github.com/huyhoangvp002/Delivery_test/db/sqlc"
	"github.com/huyhoangvp002/Delivery_test/util"
	"github.com/joho/godotenv"
)

func main() {
	// Load file .env
	err := godotenv.Load("app.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Load config
	duration := os.Getenv("ACCESS_TOKEN_DURATION")
	if duration == "" {
		log.Fatal("ACCESS_TOKEN_DURATION environment variable not set")
	}
	durationStr, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatal("invalid ACCESS_TOKEN_DURATION:", err)
	}
	config := util.Config{
		TokenSymmetricKey:   os.Getenv("TOKEN_SYMMETRIC_KEY"),
		AccessTokenDuration: durationStr,
		ServerAddress:       os.Getenv("TOKEN_SYMMETRIC_KEY"),
		DBSource:            os.Getenv("DB_SOURCE"),
	}
	//  Lấy biến môi trường DB_SOURCE
	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		log.Fatal("DB_SOURCE environment variable not set")
	}

	//  Mở kết nối database
	conn, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("cannot open database:", err)
	}

	// Test kết nối DB
	err = conn.Ping()
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	fmt.Println("Connected to the database!")

	//Tạo store (Querier)
	store := db.New(conn)

	//Tạo API server
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	//Chạy server trên cổng 8080
	addr := ":9999"
	fmt.Printf("Server is running at %s\n", addr)
	err = server.Start(addr)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
