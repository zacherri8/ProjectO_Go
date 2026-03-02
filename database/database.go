package database

import (
	"context"
	"log"
	"time"

	"projecto-backend/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	databaseURL := config.GetEnv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("❌ DATABASE_URL not found in .env")
	}

	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatal("❌ Error parsing DATABASE_URL:", err)
	}

	cfg.MaxConns = 10
	cfg.MaxConnIdleTime = 3 * time.Minute

	DB, err = pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Fatal("❌ Failed to connect to PostgreSQL:", err)
	}

	log.Println("✅ Connected to PostgreSQL")
}
