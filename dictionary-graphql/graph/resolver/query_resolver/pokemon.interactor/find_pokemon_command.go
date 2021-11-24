package pokemon_interactor

import (
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/model"
	"sort"

	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type FindPokemonCommand struct {
	DB *gorm.DB
}

func (c *FindPokemonCommand) Execute(first *int, after *int) ([]*model.Pokemon, error) {
	f, a, err := c.decideParameters(first, after)
	if err != nil {
		return nil, err
	}

	pokemons, err := c.findPokemons(f, a)
	if err != nil {
		return nil, err
	}

	err = c.resolveRelations(pokemons)
	if err != nil {
		return nil, err
	}

	return *pokemons, nil
}

func (c *FindPokemonCommand) decideParameters(first *int, after *int) (int, int, error) {
	var (
		f = 0
		a = 0
	)

	const (
		firstMin = 0
		firstMax = 15
		afterMin = 0
	)

	if first != nil {
		if *first < firstMin {
			err := errors.Cause(&IllegalArgument{
				message: fmt.Sprintf("first less then %d: first = %d", firstMin, *first),
			})

			return 0, 0, err
		}

		if *first > firstMax {
			err := errors.Cause(&IllegalArgument{
				message: fmt.Sprintf("first graeter then %d: first = %d", firstMax, *first),
			})

			return 0, 0, err
		}

		f = *first
	} else {
		f = firstMax
	}

	if after != nil {
		if *after < afterMin {
			err := errors.Cause(&IllegalArgument{
				message: fmt.Sprintf("offset less then %d: offset = %d", afterMin, *after),
			})

			return 0, 0, err
		}

		a = *after
	} else {
		a = afterMin
	}

	return f, a, nil
}

func (c *FindPokemonCommand) findPokemons(first int, after int) (*[]*model.Pokemon, error) {
	pokemons := &[]*model.Pokemon{}

	err := c.DB.Model(&model.Pokemon{}).Where("id BETWEEN ? AND ?", after-1, after+first).Scan(pokemons).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Cause(&PokemonNotFound{})
		}

		return nil, errors.WithStack(err)
	}

	return pokemons, nil
}

func (c *FindPokemonCommand) resolveRelations(pokemons *[]*model.Pokemon) error {
	pokemonIDs := []uint{}

	for _, pokemon := range *pokemons {
		pokemonIDs = append(pokemonIDs, pokemon.ID)
	}

	rs := []*result{}

	err := c.DB.
		Table("pokemons").
		Select(
			// pokemon
			"pokemons.id AS pokemon_id",
			// gender
			"genders.id AS gender_id",
			"genders.name AS gender_name",
			"genders.icon_url AS gender_icon_url",
			// type
			"types.id AS type_id",
			"types.name AS type_name",
			"types.icon_url AS type_icon_url",
			// characteristics
			"characteristics.id AS characteristic_id",
			"characteristics.name AS characteristic_name",
			"characteristics.description AS characteristic_description",
			// description
			"descriptions.id AS description_id",
			"descriptions.text AS description_text",
			"descriptions.series AS description_series",
		).
		Joins(`
			LEFT OUTER JOIN pokemon_genders ON pokemons.id = pokemon_genders.pokemon_id
			LEFT OUTER JOIN genders ON pokemon_genders.gender_id = genders.id
		`).
		Joins(`
			LEFT OUTER JOIN pokemon_types ON pokemons.id = pokemon_types.pokemon_id
			LEFT OUTER JOIN types ON pokemon_types.type_id = types.id
		`).
		Joins(`
			LEFT OUTER JOIN pokemon_characteristics ON pokemons.id = pokemon_characteristics.pokemon_id
			LEFT OUTER JOIN characteristics ON pokemon_characteristics.characteristic_id = characteristics.id
		`).
		Joins("LEFT OUTER JOIN descriptions ON pokemons.id = descriptions.pokemon_id").
		Where("pokemons.id in ?", pokemonIDs).
		Find(&rs).Error

	if err != nil {
		return errors.WithStack(err)
	}

	s := map[uint]*[]*result{}

	for _, pokemon := range *pokemons {
		s[pokemon.ID] = &[]*result{}

		for _, r := range rs {
			if pokemon.ID == (*r).PokemonID {
				*s[pokemon.ID] = append(*s[pokemon.ID], r)
			}
		}
	}

	for _, pokemon := range *pokemons {
		gs := map[uint]interface{}{}
		ts := map[uint]interface{}{}
		cs := map[uint]interface{}{}
		ds := map[uint]interface{}{}

		for _, r := range *s[pokemon.ID] {
			if r.GenderID != nil {
				id := *r.GenderID

				if gs[id] == nil {
					gs[id] = &model.Gender{
						Model: gorm.Model{
							ID: id,
						},
						Name:    *r.GenderName,
						IconURL: *r.GenderIconURL,
					}
				}
			}

			if r.TypeID != nil {
				id := *r.TypeID

				if ts[id] == nil {
					ts[id] = &model.Type{
						Model: gorm.Model{
							ID: id,
						},
						Name:    *r.TypeName,
						IconURL: *r.TypeIconURL,
					}
				}
			}

			if r.CharacteristicID != nil {
				id := *r.CharacteristicID

				if cs[id] == nil {
					cs[id] = &model.Characteristic{
						Model: gorm.Model{
							ID: id,
						},
						Name:        *r.CharacteristicName,
						Description: *r.CharacteristicDescription,
					}
				}
			}

			if r.DescriptionID != nil {
				id := *r.DescriptionID

				if ds[id] == nil {
					ds[id] = &model.Description{
						Model: gorm.Model{
							ID: id,
						},
						Text:   *r.DescriptionText,
						Series: *r.DescriptionSeries,
					}
				}
			}
		}

		if len(gs) > 0 {
			l := c.sort(gs)
			for _, g := range *l {
				pokemon.Genders = append(pokemon.Genders, *g.(*model.Gender))
			}
		}

		if len(ts) > 0 {
			l := c.sort(ts)
			for _, t := range *l {
				pokemon.Types = append(pokemon.Types, *t.(*model.Type))
			}
		}

		if len(cs) > 0 {
			l := c.sort(cs)
			for _, c := range *l {
				pokemon.Characteristics = append(pokemon.Characteristics, *c.(*model.Characteristic))
			}
		}

		if len(ds) > 0 {
			l := c.sort(ds)
			for _, d := range *l {
				pokemon.Descriptions = append(pokemon.Descriptions, *d.(*model.Description))
			}
		}
	}

	return nil
}

func (_ *FindPokemonCommand) sort(s map[uint]interface{}) *[]interface{} {
	r := []interface{}{}

	keys := make([]int, 0, len(s))
	for k := range s {
		keys = append(keys, int(k))
	}

	sort.Ints(keys)

	for _, v := range keys {
		r = append(r, s[uint(v)])
	}

	return &r
}

type result struct {
	PokemonID uint

	// gender
	GenderID      *uint
	GenderName    *string
	GenderIconURL *string

	// type
	TypeID      *uint
	TypeName    *string
	TypeIconURL *string

	// characteristics
	CharacteristicID          *uint
	CharacteristicName        *string
	CharacteristicDescription *string

	// descriptions
	DescriptionID     *uint
	DescriptionText   *string
	DescriptionSeries *string
}
