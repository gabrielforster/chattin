package database

import (
  "os"

  "database/sql"
  _ "github.com/lib/pq"
)

type DatabaseConnection struct {
  db *sql.DB
}

func NewDatabase() (*DatabaseConnection, error) {
  db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

  if err != nil {
    return nil, err
  }

  return &DatabaseConnection{db}, nil
}

func (conn *DatabaseConnection) Close() error {
  return conn.db.Close()
}

func (conn *DatabaseConnection) GetDB() *sql.DB {
  return conn.db
}
