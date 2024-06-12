package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "user=username dbname=mydb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	conn, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	filePath := "data.csv"
	tableName := "nombre_de_la_tabla"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	connStr := "user=username dbname=mydb sslmode=disable"
	pgxConn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer pgxConn.Close()

	copyCount, err := pgxConn.CopyFrom(
		context.Background(),
		pgx.Identifier{tableName},
		[]string{"columna1", "columna2", "columna3"},
		pgx.CopyFromReader(file),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Se copiaron %d filas correctamente.\n", copyCount)
}
