package card_type

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	typedm "integration-test/internal/domain/card_type"
	"testing"
)

func TestRepository(t *testing.T) {
	repositories := createRepositories(t)

	for _, r := range repositories {
		t.Run(r.Name, func(t *testing.T) {
			testCreateCardType(t, r.Repository)
		})
	}
}

// Set up
type Repository struct {
	Name       string
	Repository typedm.Repository
}

func createRepositories(t *testing.T) []Repository {
	return []Repository{
		{
			Name:       "postgres",
			Repository: newPostgresRepository(t),
		},
	}
}

func newPostgresRepository(t *testing.T) typedm.Repository {
	db, err := NewPostgresConnection()
	require.NoError(t, err)

	return NewPostgresRepository(db)
}

// const variables
const (
	INVALID_ID     = ""
	INVALID_NUMBER = ""
)

// Test functions
func testCreateCardType(t *testing.T, r typedm.Repository) {
	t.Helper()
	cases := []struct {
		Name string
		In   *typedm.CardType
		Err  bool
	}{
		{Name: "valid card type", In: newCardType(), Err: false},
		{Name: "with invalid ID", In: newCardTypeWithID(INVALID_ID), Err: true},
		{Name: "with invalid number", In: newCardTypeWithNumber(INVALID_NUMBER), Err: true},
	}

	for _, tc := range cases {
		err := r.Create(context.Background(), tc.In)
		if tc.Err {
			require.Error(t, err, tc.Name)
			continue
		}
		require.NoError(t, err, tc.Name)

		found, err := r.FindOne(context.Background(), tc.In.Guid)
		require.NoError(t, err, tc.Name)

		assert.Equal(t, tc.In, found, tc.Name)
	}
}

func newCardType() *typedm.CardType {
	return &typedm.CardType{Guid: gofakeit.UUID(), Provider: gofakeit.CreditCardType(), Number: gofakeit.CreditCard().Number}
}

func newCardTypeWithID(id string) *typedm.CardType {
	ct := newCardType()
	ct.Guid = id
	return ct
}

func newCardTypeWithNumber(number string) *typedm.CardType {
	ct := newCardType()
	ct.Number = number
	return ct
}
