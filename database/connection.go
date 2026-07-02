package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"eventbe/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectMySQL() *sql.DB {
	if DB != nil {
		return DB
	}

	dsn := config.GetEnv("MYSQL_DSN", "")
	if dsn == "" {
		host := config.GetEnv("DB_HOST", "localhost")
		port := config.GetEnv("DB_PORT", "3306")
		user := config.GetEnv("DB_USER", "root")
		password := config.GetEnv("DB_PASSWORD", config.GetEnv("DB_PASS", ""))
		name := config.GetEnv("DB_NAME", "eventair")

		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", user, password, host, port, name)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open mysql connection: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping mysql: %v", err)
	}

	DB = db
	log.Println("Connected to MySQL")
	return DB
}

func ConnectPostgres() *sql.DB {
	return ConnectMySQL()
}

func ConnectMongoDB() {
	// Placeholder for MongoDB connection
	fmt.Println("Connecting to MongoDB...")
}

func AutoMigrate() {
	// Placeholder for AutoMigrate
	log.Println("Running AutoMigrate...")
}
