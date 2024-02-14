package card

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"integration-test/internal/domain/card"
	"log"
	"os"
)

type postgresRepository struct {
	db *pgx.Conn
}

func NewPostgresRepository(db *pgx.Conn) card.Repository {

	return &postgresRepository{db: db}
}

func (p *postgresRepository) Create(ctx context.Context, c *card.Card) error {
	clauses := map[string]any{
		"guid":        c.ID,
		"number":      c.Number,
		"masked_pan":  c.MaskedPan,
		"customer_id": c.CustomerID,
	}
	sql, args, err := squirrel.Insert("customer_card").SetMap(clauses).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}
	_, err = p.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (p *postgresRepository) FindOne(ctx context.Context, ID string) (*card.Card, error) {
	sql, args, err := squirrel.Select("guid", "number", "masked_pan", "customer_id").From("customer_card").Where(squirrel.Eq{"guid": ID}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	row := p.db.QueryRow(ctx, sql, args...)

	var card card.Card
	err = row.Scan(&card.ID, &card.Number, &card.MaskedPan, &card.CustomerID)
	if err != nil {
		return nil, err
	}
	return &card, nil
}

func NewPostgresConnection() (*pgx.Conn, error) {
	val, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Fatal("database url not exists")
	}
	conn, err := pgx.Connect(context.Background(), val)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
