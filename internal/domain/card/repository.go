package card

import "context"

type Repository interface {
	Create(ctx context.Context, c *Card) error
	FindOne(ctx context.Context, ID string) (*Card, error)
}
