package card

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	cardm "integration-test/internal/domain/card"
	"testing"
)

// Set up
func TestRepository(t *testing.T) {
	repositories := createRepositories(t)

	for _, r := range repositories {
		t.Run(r.Name, func(t *testing.T) {
			testCreateCard(t, r.Repository)
		})
	}
}

type Repository struct {
	Name       string
	Repository cardm.Repository
}

func newPostgresRepository(t *testing.T) cardm.Repository {
	db, err := NewPostgresConnection()
	require.NoError(t, err)

	return NewPostgresRepository(db)
}

func createRepositories(t *testing.T) []Repository {
	return []Repository{
		{
			Name:       "postgres",
			Repository: newPostgresRepository(t),
		},
	}
}

// Test functions
const (
	INVALID_ID     = ""
	INVALID_NUMBER = ""
)

func testCreateCard(t *testing.T, r cardm.Repository) {
	cases := []struct {
		Name string
		In   *cardm.Card
		Err  bool
	}{
		{Name: "valid card", In: newCard(), Err: false},
		{Name: "invalid card(without ID)", In: newCardWithID(INVALID_ID), Err: true},
		{Name: "invalid card(without number)", In: newCardWithNumber(INVALID_NUMBER), Err: true},
	}

	for _, tc := range cases {
		err := r.Create(context.Background(), tc.In)
		if tc.Err {
			require.Error(t, err, tc.Name)
			continue
		}
		require.NoError(t, err)

		found, err := r.FindOne(context.Background(), tc.In.ID)
		require.NoError(t, err)

		assert.Equal(t, tc.In, found)
	}
}

func newCard() *cardm.Card {

	return &cardm.Card{ID: gofakeit.UUID(), Number: gofakeit.CreditCard().Number, MaskedPan: gofakeit.CreditCard().Number, CustomerID: uuid.NewString()}
}

func newCardWithID(id string) *cardm.Card {
	card := newCard()
	card.ID = id
	return card
}

func newCardWithNumber(number string) *cardm.Card {
	card := newCard()
	card.Number = number
	return card
}
