package configs

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

type DBConfig struct {
	Host           string
	Port           string
	User           string
	Password       string
	DBName         string
	DataSourceName string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Host:           getEnvOrDefault("DB_HOST", "localhost"),
		Port:           getEnvOrDefault("DB_PORT", "5432"),
		User:           getEnvOrDefault("DB_USER", "postgres"),
		Password:       getEnvOrDefault("DB_PASSWORD", "postgres"),
		DBName:         getEnvOrDefault("DB_NAME", "rockstart"),
		DataSourceName: getEnvOrDefault("DB_DATA_SOURCE_NAME", "./database.db"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *DBConfig) ConnectDB(ctx context.Context) *pgx.Conn {
	conn, err := pgx.Connect(ctx, fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", c.User, c.Password, c.Host, c.Port, c.DBName))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	return conn
}

func (c *DBConfig) ConnectSQLiteDB(ctx context.Context) *sql.DB {
	conn, err := sql.Open("sqlite", c.DataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to the SQLite database: %v", err)
	}
	return conn
}
