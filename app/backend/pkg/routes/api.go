package routes

import (
	"piteroni/dictionary-go-nuxt-graphql/pkg/controllers"
	"piteroni/dictionary-go-nuxt-graphql/pkg/usecases/pokemon"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAPIRouting(e *gin.Engine, db *gorm.DB) {
	pokemonController := &controllers.PokemonController{
		PokemonDetailsAcquisition: pokemon.NewPokemonDetailsAcquisition(db),
	}

	r := e.Group("/api/i")
	{
		r.GET("/pokemons/:id", pokemonController.PokemonDetailsAcquisitionHandler())
	}
}
