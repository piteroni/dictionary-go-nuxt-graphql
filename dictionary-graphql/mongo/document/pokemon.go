package document

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pokemon struct {
	Record
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

	Descriptions    *[]Description        `bson:"descriptions"`
	Types           *[]primitive.ObjectID `bson:"types"`
	Genders         *[]primitive.ObjectID `bson:"genders"`
	Characteristics *[]primitive.ObjectID `bson:"characteristics"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
