package card_type

import "context"

type Repository interface {
	Create(ctx context.Context, cardType *CardType) error
	FindOne(ctx context.Context, guid string) (*CardType, error)
}
