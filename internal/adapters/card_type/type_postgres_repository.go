package card_type

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"integration-test/internal/domain/card_type"
)

type postgresRepository struct {
	db *pgx.Conn
}

func NewPostgresRepository(db *pgx.Conn) card_type.Repository {

	return &postgresRepository{db: db}
}

func (p *postgresRepository) Create(ctx context.Context, c *card_type.CardType) error {
	clauses := map[string]any{
		"guid":     c.Guid,
		"provider": c.Provider,
		"number":   c.Number,
	}
	sql, args, err := squirrel.Insert("card_type").SetMap(clauses).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = p.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (p postgresRepository) FindOne(ctx context.Context, guid string) (*card_type.CardType, error) {
	sql, args, err := squirrel.Select("guid", "provider", "number").From("card_type").Where(squirrel.Eq{"guid": guid}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	row := p.db.QueryRow(ctx, sql, args...)
	var c card_type.CardType
	err = row.Scan(&c.Guid, &c.Provider, &c.Number)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func NewPostgresConnection() (*pgx.Conn, error) {
	connStr := "host=localhost port=5432 user=postgres password=123 dbname=test sslmode=disable"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
