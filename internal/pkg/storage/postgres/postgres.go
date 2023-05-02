package postgres

import (
	"context"
	"fmt"
	"log"

	"link_shorter/internal/config"
	"link_shorter/internal/pkg/storage"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // nolint
)

const (
	linksTable = "links"
)

type postgres struct {
	db *sqlx.DB
}

func connect(cfg config.Config) *sqlx.DB {
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBName, cfg.DBPassword, cfg.DBSSLMode)

	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		log.Fatalf("sqlx.Open: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("db.Ping: %v", err)
	}

	return db
}

func (p *postgres) Shutdown(ctx context.Context) {
	p.db.Close()
}

func NewPostgres(cfg config.Config) storage.Store {
	return &postgres{
		db: connect(cfg),
	}
}
