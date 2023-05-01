package postgres

import (
	"fmt"
	"golang.org/x/net/context"
	"link_shorter/internal/pkg/model"
)

func (p *postgres) Create(ctx context.Context, url, token string) (*model.Link, error) {
	var record model.Link

	query := fmt.Sprintf("INSERT INTO %s (token, url) VALUES ($1, $2) RETURNING *", linksTable)
	row := p.db.QueryRowx(query, token, url)
	if err := row.StructScan(&record); err != nil {
		return nil, fmt.Errorf("failed to scan in create link: %v", err)
	}

	return &record, nil
}
