package database

import (
	"bioskop/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"embed"
	migrate "github.com/rubenv/sql-migrate"
)

var dbMigrations embed.FS
var DB *sql.DB

func InitDB() {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PSWD, config.DB_NAME)

	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	// DBMigrate(DB)

	fmt.Println("Berhasil terhubung ke database")
}

func DBMigrate(dbParam *sql.DB) {
    migrations := &migrate.EmbedFileSystemMigrationSource{
       FileSystem: dbMigrations,
       Root:       "sql_migrations",
    }

    n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
    if errs != nil {
       panic(errs)
    }

    DB = dbParam

    fmt.Println("Migration success, applied", n, "migrations!")
}