package resolvers

import (
	"context"
	"testing"

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

	t.Run("Success", func(t *testing.T) {
		testObj.MockPokeService.EXPECT().Create(context.Background(), models.CreatePokemonInput{}).Return(&models.Pokemon{}, nil)

		poke, err := testObj.MockMutationResolver.CreatePokemon(context.Background(), models.CreatePokemonInput{})
		assert.NotNil(t, poke)
		assert.Nil(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		testObj.MockPokeService.EXPECT().Create(context.Background(), models.CreatePokemonInput{}).Return(nil, errors.ErrorOccur)

		poke, err := testObj.MockMutationResolver.CreatePokemon(context.Background(), models.CreatePokemonInput{})
		assert.Nil(t, poke)
		assert.NotNil(t, err)
	})
}
