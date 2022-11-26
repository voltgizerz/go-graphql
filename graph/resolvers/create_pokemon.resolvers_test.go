package resolvers

import (
	"testing"

	"github.com/go-graphql/mocks"
)

type MockObjectData struct {
	Mutation        mutationResolver
	MockPokeService *mocks.PokemonServiceInterface
	MocTypeService  *mocks.TypeServiceInterface
}

func initMockObject(t *testing.T) MockObjectData {
	resolver := &Resolver{
		PokemonService: &mocks.PokemonServiceInterface{},
		TypeService:    &mocks.TypeServiceInterface{},
	}

	mockObj := MockObjectData{
		Mutation:        mutationResolver{resolver},
		MockPokeService: &mocks.PokemonServiceInterface{},
		MocTypeService:  &mocks.TypeServiceInterface{},
	}

	return mockObj
}

// func TestPokemonResolver_Create(t *testing.T) {
// 	mockObj := initMockObject(t)
// 	mockObj.MockPokeService.On("Create", mock.Anything, mock.Anything).Return(&models.Pokemon{}, nil).Once()

// 	res, err := mockObj.Mutation.CreatePokemon(context.Background(), models.CreatePokemonInput{})
// 	assert.NotNil(t, res)
// 	assert.Nil(t, err)
// }
