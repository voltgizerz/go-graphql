package resolvers

import (
	"context"
	"testing"

	"github.com/go-graphql/auth"
	"github.com/go-graphql/errors"
	"github.com/go-graphql/graph/generated"
	"github.com/go-graphql/mocks"
	"github.com/go-graphql/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type MockObjectData struct {
	MockMutationResolver generated.MutationResolver
	MockPokeService      *mocks.MockPokemonServiceInterface
	MocTypeService       *mocks.MockTypeServiceInterface
}

func initMockObject(t *testing.T) MockObjectData {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPokeService := mocks.NewMockPokemonServiceInterface(ctrl)
	MockTypeService := mocks.NewMockTypeServiceInterface(ctrl)

	resolverData := ResolverData{
		PokemonService: mockPokeService,
		TypeService:    MockTypeService,
	}

	mockResolver := NewResolver(resolverData)
	mockMutationResolver := mockResolver.Mutation()

	mockObj := MockObjectData{
		MockMutationResolver: mockMutationResolver,
		MockPokeService:      mockPokeService,
		MocTypeService:       MockTypeService,
	}

	return mockObj
}

func TestPokemonResolver_Create(t *testing.T) {
	testObj := initMockObject(t)

	user := &auth.User{UserID: 123, Name: "felix", IsAdmin: true}
	ctx := context.Background()
	ctx = context.WithValue(ctx, auth.CONTEXT_USER, user)

	t.Run("Success", func(t *testing.T) {
		testObj.MockPokeService.EXPECT().Create(ctx, models.CreatePokemonInput{}).Return(&models.Pokemon{}, nil)

		poke, err := testObj.MockMutationResolver.CreatePokemon(ctx, models.CreatePokemonInput{})
		assert.NotNil(t, poke)
		assert.Nil(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		testObj.MockPokeService.EXPECT().Create(ctx, models.CreatePokemonInput{}).Return(nil, errors.ErrorOccur)

		poke, err := testObj.MockMutationResolver.CreatePokemon(ctx, models.CreatePokemonInput{})
		assert.Nil(t, poke)
		assert.NotNil(t, err)
	})
}
