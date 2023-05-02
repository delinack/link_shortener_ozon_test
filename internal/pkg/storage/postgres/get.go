package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"link_shorter/internal/pkg/model"
	"link_shorter/internal/pkg/storage"
)

func (p *postgres) GetByToken(ctx context.Context, token string) (*model.Link, error) {
	var record model.Link

	query := fmt.Sprintf("SELECT * FROM %s WHERE token = $1", linksTable)
	row := p.db.QueryRowx(query, token)
	if err := row.StructScan(&record); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("failed to scan in get link: %v", err)
	}

	return &record, nil
}

func (p *postgres) GetByURL(ctx context.Context, url string) (*model.Link, error) {
	var record model.Link

	query := fmt.Sprintf("SELECT * FROM %s WHERE url = $1", linksTable)
	row := p.db.QueryRowx(query, url)
	if err := row.StructScan(&record); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("failed to scan in get link: %v", err)
	}

	return &record, nil
}
