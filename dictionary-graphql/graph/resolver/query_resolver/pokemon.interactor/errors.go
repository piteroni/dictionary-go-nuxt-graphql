package pokemon_interactor

var _ error = (*PokemonNotFound)(nil)

type PokemonNotFound struct {
	message string
}

func (e *PokemonNotFound) Error() string {
	return e.message
}

var _ error = (*IllegalArguments)(nil)

type IllegalArguments struct {
	Message string
}

func (e *IllegalArguments) Error() string {
	return e.Message
}
