package pokemon_loader

var _ error = (*PokemonNotFound)(nil)

type PokemonNotFound struct {
	message string
}

func (e *PokemonNotFound) Error() string {
	return e.message
}

type IllegalArgument struct {
	message string
}

var _ error = (*IllegalArgument)(nil)

func (e *IllegalArgument) Error() string {
	return e.message
}
