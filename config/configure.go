package config

import (
	"database/sql"
	"log"
)

// ServerGen represents server
type ServerGen struct {
	Reader      *sql.DB
	Writer      *sql.DB
	Port        string
}

func ConfigureMySQL() (*sql.DB, *sql.DB) {

	readerConfig := Option{
		Host:     MySQL.ReaderHost,
		Port:     MySQL.ReaderPort,
		Database: MySQL.Database,
		User:     MySQL.ReaderUser,
		Password: MySQL.ReaderPassword,
	}
	writerConfig := Option{
		Host:     MySQL.WriterHost,
		Port:     MySQL.WriterPort,
		Database: MySQL.Database,
		User:     MySQL.WriterUser,
		Password: MySQL.WriterPassword,
	}

	reader, writer, err := SetupDatabase(readerConfig, writerConfig)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect mysql", err)
	}
	log.Println("MySQL connection is successfully established!")
	return reader, writer
}

