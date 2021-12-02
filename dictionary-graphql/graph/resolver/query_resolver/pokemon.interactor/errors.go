package pokemon_interactor

var _ error = (*PokemonNotFound)(nil)

type PokemonNotFound struct {
	message string
}

func (e *PokemonNotFound) Error() string {
	return e.message
}

type IllegalArguments struct {
	message string
}

var _ error = (*IllegalArguments)(nil)

func (e *IllegalArguments) Error() string {
	return e.message
}
