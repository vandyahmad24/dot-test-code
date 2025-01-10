package main

import (
	"database/sql"
	"dot-test-vandy/config"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

// digunakan untuk menjalankan migrate otomatis
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide action!")
		return
	}
	migrateStatus := os.Args[1]
	var migrationName string
	if len(os.Args) >= 3 {
		migrationName = os.Args[2]
	}

	cfg := config.NewConfig()

	username := cfg.MYSQL_USERNAME
	password := cfg.MYSQL_PASSWORD
	host := cfg.MYSQL_HOST
	port := cfg.MYSQL_PORT
	dbname := cfg.MYSQL_DATABASE

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FJakarta", username, password, host, port, dbname)
	log.Println(mysqlInfo)
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	switch migrateStatus {
	case "up":
		applyUpMigrations(db)
	case "down":
		applyDownMigrations(db)
	case "status":
		showMigrationStatus(db)
	case "create":
		createMigration(migrationName)
	default:
		fmt.Println("Invalid action!")
	}
}

func applyUpMigrations(db *sql.DB) {
	log.Println("running up migration")
	migrations := &migrate.FileMigrationSource{
		Dir: "migration",
	}
	m, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Applied %d migrations!\n", m)
}

func applyDownMigrations(db *sql.DB) {
	log.Println("running down migration")
	migrations := &migrate.FileMigrationSource{
		Dir: "migration",
	}
	m, err := migrate.ExecMax(db, "mysql", migrations, migrate.Down, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Applied %d migrations!\n", m)
}

func showMigrationStatus(db *sql.DB) {
	log.Println("running status migration")
	m, err := migrate.GetMigrationRecords(db, "mysql")
	if err != nil {
		panic(err)
	}
	for _, record := range m {
		fmt.Printf("%s - %s\n", record.Id, record.AppliedAt)
	}
}

func createMigration(migrationName string) {
	// Generate a timestamp for the file name
	timestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("migration/%s_%s.sql", timestamp, migrationName)

	// Content of the migration file
	content := `-- +migrate Up
-- SQL statements for migration up goes here

-- +migrate Down
-- SQL statements for migration down goes here
`

	// Create the file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write content to the file
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Migration file created successfully:", fileName)
}
