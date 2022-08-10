package config

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/ichtrojan/thoth"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func Database() *sql.DB {
	logger, _ := thoth.Init("log")

	user, exist := os.LookupEnv("DB_USER")

	if !exist {
		logger.Log(errors.New("DB_USER not set in .env"))
		log.Fatal("DB_USER not set in .env")
	}

	pass, exist := os.LookupEnv("DB_PASS")

	if !exist {
		logger.Log(errors.New("DB_PASS not set in .env"))
		log.Fatal("DB_PASS not set in .env")
	}

	credentials := fmt.Sprintf("user=%s dbname=todoapp password=%s sslmode=disable", user, pass)
	// dsn := flag.String("dsn", credentials, "Psql data source name")
	flag.Parse()

	// database, err := sql.Open("mysql", credentials)
	database, err := sql.Open("postgres", credentials)
	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	} else {
		fmt.Println("Database Connection Successful")
	}
	// defer db.Close()

	return database
}

// func openDB(dsn string) (*sql.DB, error) {
// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err = db.Ping(); err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }
