package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*sql.DB, error) {
	host := "storidb.cbo8uuu4wq2w.us-east-1.rds.amazonaws.com"
	port := "5432"
	user := "postgres"
	password := "storipass"
	dbname := "storidb"

	fmt.Printf("DB_HOST: %s\n", host)
	fmt.Printf("DB_PORT: %s\n", port)
	fmt.Printf("DB_USER: %s\n", user)
	fmt.Printf("DB_PASSWORD: %s\n", password)
	fmt.Printf("DB_NAME: %s\n", dbname)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		host, port, user, password, dbname)

	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := NewPostgresConnection()
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
		return
	}
	defer db.Close()

	fmt.Println("Successfully connected to the database")

	// Realiza una simple consulta de prueba
	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		fmt.Printf("Failed to execute query: %v\n", err)
		return
	}

	fmt.Printf("PostgreSQL version: %s\n", version)
}
