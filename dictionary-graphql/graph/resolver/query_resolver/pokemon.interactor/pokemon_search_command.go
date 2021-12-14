package pokemon_interactor

import (
	"context"
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/mongo/collection"
	"piteroni/dictionary-go-nuxt-graphql/mongo/document"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PokemonSearchCommand interface {
	Execute(first *int, after *string) ([]*document.Pokemon, error)
}

var _ PokemonSearchCommand = (*PokemonSearchCommandImpl)(nil)

type PokemonSearchCommandImpl struct {
	DB      *mongo.Database
	Context context.Context
}

func (c *PokemonSearchCommandImpl) Execute(first *int, after *string) ([]*document.Pokemon, error) {
	f, a, err := c.prepareParams(first, after)
	if err != nil {
		return nil, err
	}

	pokemons, err := c.findPokemons(f, a)
	if err != nil {
		return nil, err
	}

	// 	err = c.resolveRelations(pokemons)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	return *pokemons, nil
}

func (c *PokemonSearchCommandImpl) prepareParams(first *int, after *string) (int, primitive.ObjectID, error) {
	f := 0

	const (
		min = 0
		max = 64
	)

	if first != nil {
		value := *first

		if value < min {
			err := errors.Cause(&IllegalArguments{
				message: fmt.Sprintf("first less then %d: first = %d", min, value),
			})

			return 0, primitive.NilObjectID, err
		}

		if value > max {
			err := errors.Cause(&IllegalArguments{
				message: fmt.Sprintf("first graeter then %d: first = %d", max, value),
			})

			return 0, primitive.NilObjectID, err
		}

		f = value
	} else {
		f = max
	}

	if after == nil {
		return f, primitive.ObjectID{}, nil
	}

	a, err := primitive.ObjectIDFromHex(*after)
	if err != nil {
		return 0, primitive.NilObjectID, errors.WithStack(err)
	}

	return f, a, nil
}

func (c *PokemonSearchCommandImpl) findPokemons(first int, after primitive.ObjectID) (*[]*document.Pokemon, error) {
	pokemons := []*document.Pokemon{}

	condition := interface{}(nil)

	if after.IsZero() {
		condition = bson.D{{}}
	} else {
		condition = bson.M{"_id": bson.D{{Key: "$gt", Value: after}}}
	}

	f := int64(first)
	option := options.FindOptions{Limit: &f}

	cursor, err := c.DB.Collection(collection.Pokemons).Find(c.Context, condition, &option)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = cursor.All(c.Context, &pokemons)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &pokemons, nil
}

// func (c *PokemonSearchCommandImpl) resolveRelations(pokemons *[]*model.Pokemon) error {
// 	pokemonIDs := []uint{}

// 	for _, pokemon := range *pokemons {
// 		pokemonIDs = append(pokemonIDs, pokemon.ID)
// 	}

// 	rs := []*result{}

// 	err := c.DB.
// 		Table("pokemons").
// 		Select(
// 			// pokemon
// 			"pokemons.id AS pokemon_id",
// 			// gender
// 			"genders.id AS gender_id",
// 			"genders.name AS gender_name",
// 			"genders.icon_url AS gender_icon_url",
// 			// type
// 			"types.id AS type_id",
// 			"types.name AS type_name",
// 			"types.icon_url AS type_icon_url",
// 			// characteristics
// 			"characteristics.id AS characteristic_id",
// 			"characteristics.name AS characteristic_name",
// 			"characteristics.description AS characteristic_description",
// 			// description
// 			"descriptions.id AS description_id",
// 			"descriptions.text AS description_text",
// 			"descriptions.series AS description_series",
// 		).
// 		Joins("LEFT OUTER JOIN pokemon_genders ON pokemons.id = pokemon_genders.pokemon_id").
// 		Joins("LEFT OUTER JOIN genders ON pokemon_genders.gender_id = genders.id").
// 		Joins("LEFT OUTER JOIN pokemon_types ON pokemons.id = pokemon_types.pokemon_id").
// 		Joins("LEFT OUTER JOIN types ON pokemon_types.type_id = types.id").
// 		Joins("LEFT OUTER JOIN pokemon_characteristics ON pokemons.id = pokemon_characteristics.pokemon_id").
// 		Joins("LEFT OUTER JOIN characteristics ON pokemon_characteristics.characteristic_id = characteristics.id").
// 		Joins("LEFT OUTER JOIN descriptions ON pokemons.id = descriptions.pokemon_id").
// 		Where("pokemons.id in ?", pokemonIDs).
// 		Find(&rs).Error

// 	if err != nil {
// 		return err
// 	}

// 	s := map[uint]*[]*result{}

// 	for _, pokemon := range *pokemons {
// 		s[pokemon.ID] = &[]*result{}

// 		for _, r := range rs {
// 			if pokemon.ID == (*r).PokemonID {
// 				*s[pokemon.ID] = append(*s[pokemon.ID], r)
// 			}
// 		}
// 	}

// 	for _, pokemon := range *pokemons {
// 		gs := map[uint]interface{}{}
// 		ts := map[uint]interface{}{}
// 		cs := map[uint]interface{}{}
// 		ds := map[uint]interface{}{}

// 		for _, r := range *s[pokemon.ID] {
// 			if r.GenderID != nil {
// 				id := *r.GenderID

// 				if gs[id] == nil {
// 					gs[id] = &model.Gender{
// 						Model: gorm.Model{
// 							ID: id,
// 						},
// 						Name:    *r.GenderName,
// 						IconURL: *r.GenderIconURL,
// 					}
// 				}
// 			}

// 			if r.TypeID != nil {
// 				id := *r.TypeID

// 				if ts[id] == nil {
// 					ts[id] = &model.Type{
// 						Model: gorm.Model{
// 							ID: id,
// 						},
// 						Name:    *r.TypeName,
// 						IconURL: *r.TypeIconURL,
// 					}
// 				}
// 			}

// 			if r.CharacteristicID != nil {
// 				id := *r.CharacteristicID

// 				if cs[id] == nil {
// 					cs[id] = &model.Characteristic{
// 						Model: gorm.Model{
// 							ID: id,
// 						},
// 						Name:        *r.CharacteristicName,
// 						Description: *r.CharacteristicDescription,
// 					}
// 				}
// 			}

// 			if r.DescriptionID != nil {
// 				id := *r.DescriptionID

// 				if ds[id] == nil {
// 					ds[id] = &model.Description{
// 						Model: gorm.Model{
// 							ID: id,
// 						},
// 						Text:   *r.DescriptionText,
// 						Series: *r.DescriptionSeries,
// 					}
// 				}
// 			}
// 		}

// 		if len(gs) > 0 {
// 			l := c.sort(gs)
// 			for _, g := range *l {
// 				pokemon.Genders = append(pokemon.Genders, *g.(*model.Gender))
// 			}
// 		}

// 		if len(ts) > 0 {
// 			l := c.sort(ts)
// 			for _, t := range *l {
// 				pokemon.Types = append(pokemon.Types, *t.(*model.Type))
// 			}
// 		}

// 		if len(cs) > 0 {
// 			l := c.sort(cs)
// 			for _, c := range *l {
// 				pokemon.Characteristics = append(pokemon.Characteristics, *c.(*model.Characteristic))
// 			}
// 		}

// 		if len(ds) > 0 {
// 			l := c.sort(ds)
// 			for _, d := range *l {
// 				pokemon.Descriptions = append(pokemon.Descriptions, *d.(*model.Description))
// 			}
// 		}
// 	}

// 	return nil
// }

// func (_ *PokemonSearchCommandImpl) sort(s map[uint]interface{}) *[]interface{} {
// 	r := []interface{}{}

// 	keys := make([]int, 0, len(s))
// 	for k := range s {
// 		keys = append(keys, int(k))
// 	}

// 	sort.Ints(keys)

// 	for _, v := range keys {
// 		r = append(r, s[uint(v)])
// 	}

// 	return &r
// }

// type result struct {
// 	PokemonID uint

// 	// gender
// 	GenderID      *uint
// 	GenderName    *string
// 	GenderIconURL *string

// 	// type
// 	TypeID      *uint
// 	TypeName    *string
// 	TypeIconURL *string

// 	// characteristics
// 	CharacteristicID          *uint
// 	CharacteristicName        *string
// 	CharacteristicDescription *string

// 	// descriptions
// 	DescriptionID     *uint
// 	DescriptionText   *string
// 	DescriptionSeries *string
// }
