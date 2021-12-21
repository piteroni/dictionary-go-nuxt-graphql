package document

import (
	"piteroni/dictionary-go-nuxt-graphql/mongo/collection"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pokemon struct {
	ID          primitive.ObjectID  `bson:"_id"`
	NationalNo  int                 `bson:"national_no"`
	EvolutionID *primitive.ObjectID `bson:"evolution_id"`
	Name        string              `bson:"name"`
	ImageURL    string              `bson:"image_url"`
	Species     string              `bson:"species"`
	Height      string              `bson:"height"`
	Weight      string              `bson:"weight"`

	// @see: https://pokemondb.net/pokedex
	HeartPoint          int `bson:"heart_point"`
	AttackPoint         int `bson:"attack_point"`
	DefensePoint        int `bson:"defense_point"`
	SpecialAttackPoint  int `bson:"special_attack_point"`
	SpecialDefensePoint int `bson:"special_defense_point"`
	SpeedPoint          int `bson:"speed_point"`

	Descriptions []Description `bson:"descriptions"`

	Types           *[]Type           `bson:",omitempty"`
	Genders         *[]Gender         `bson:",omitempty"`
	Characteristics *[]Characteristic `bson:",omitempty"`

	References PokemonReferences `bson:"references"`
}

type PokemonReferences struct {
	Types           []primitive.ObjectID `bson:"types"`
	Genders         []primitive.ObjectID `bson:"genders"`
	Characteristics []primitive.ObjectID `bson:"characteristics"`
}

type PokemonAggregate struct{}

func (c PokemonAggregate) StagesOfLookUp() []bson.D {
	return []bson.D{
		{{
			Key: "$lookup", Value: bson.M{
				"from":         collection.Types,
				"localField":   "references.types",
				"foreignField": "_id",
				"as":           "types",
			},
		}},
		{{
			Key: "$lookup", Value: bson.M{
				"from":         collection.Genders,
				"localField":   "references.genders",
				"foreignField": "_id",
				"as":           "genders",
			},
		}},
		{{
			Key: "$lookup", Value: bson.M{
				"from":         collection.Characteristics,
				"localField":   "references.characteristics",
				"foreignField": "_id",
				"as":           "characteristics",
			},
		}},
	}
}
