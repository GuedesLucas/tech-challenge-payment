package config

import (
	"database/sql"
	"fmt"
	"tech-challenge-payment/config/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(dbConfig DatabaseConfig) (*gorm.DB, error) {

	err := ValidateDatabase(dbConfig)
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)
	fmt.Println("DSN:", dsn) // Adicione esta linha para verificar o DSN

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	migrations.RunMigrations(db)

	return db, nil
}

func ValidateDatabase(dbConfig DatabaseConfig) error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, "postgres")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	query := "SELECT datname FROM pg_database WHERE datname = $1"
	var dbName string
	err = db.QueryRow(query, dbConfig.Name).Scan(&dbName)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbConfig.Name)
		_, err = db.Exec(createDBQuery)
		if err != nil {
			return err
		}
	}

	db.Close()
	return nil
}
