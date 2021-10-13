package controllers

import (
	"piteroni/dictionary-go-nuxt-graphql/pkg/usecases/pokemon"

	"github.com/gin-gonic/gin"
)

type PokemonController struct {
	PokemonDetailsAcquisition *pokemon.PokemonDetailsAcquisition
}

func (ctr *PokemonController) PokemonDetailsAcquisitionHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		// ctr.PokemonDetailsAcquisition.GetPokemonDetails(1)
	}
}
