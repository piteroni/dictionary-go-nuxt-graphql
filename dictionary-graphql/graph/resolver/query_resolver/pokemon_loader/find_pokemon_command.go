package pokemon_loader

import (
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/model"
	"sort"

	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type findPokemonCommand struct {
	db *gorm.DB
}

func (c *findPokemonCommand) execute(o *int, l *int) ([]*model.Pokemon, error) {
	offset, limit, err := c.decideParameters(o, l)
	if err != nil {
		return nil, err
	}

	pokemons, err := c.findPokemons(offset, limit)
	if err != nil {
		return nil, err
	}

	err = c.resolveRelations(pokemons)
	if err != nil {
		return nil, err
	}

	return *pokemons, nil
}

func (c *findPokemonCommand) decideParameters(o *int, l *int) (offset int, limit int, err error) {
	const (
		limitMin  = 0
		limitMax  = 15
		offsetMin = 0
	)

	if l != nil {
		if *l < limitMin {
			err := errors.Cause(&IllegalArgument{
				message: fmt.Sprintf("limit less then %d: limit = %d", limitMin, *l),
			})

			return 0, 0, err
		}

		if *l > limitMax {
			err := errors.Cause(&IllegalArgument{
				message: fmt.Sprintf("limit graeter then %d: limit = %d", limitMax, *l),
			})

			return 0, 0, err
		}

		limit = *l
	} else {
		limit = limitMax
	}

	if o != nil {
		if *o < offsetMin {
			err := errors.Cause(&IllegalArgument{
				message: fmt.Sprintf("offset less then %d: offset = %d", offsetMin, *o),
			})

			return 0, 0, err
		}

		offset = *o
	} else {
		offset = offsetMin
	}

	return offset, limit, nil
}

func (c *findPokemonCommand) findPokemons(offset int, limit int) (*[]*model.Pokemon, error) {
	fmt.Printf("offset: %v\n", offset)
	pokemons := &[]*model.Pokemon{}

	err := c.db.Model(&model.Pokemon{}).Limit(limit).Offset(offset).Scan(pokemons).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Cause(&PokemonNotFound{})
		}

		return nil, errors.WithStack(err)
	}

	return pokemons, nil
}

func (c *findPokemonCommand) resolveRelations(pokemons *[]*model.Pokemon) error {
	pokemonIDs := []uint{}

	for _, pokemon := range *pokemons {
		pokemonIDs = append(pokemonIDs, pokemon.ID)
	}

	rs := []*result{}

	err := c.db.
		Table("pokemons").
		Select(
			// pokemon
			"pokemons.id as pokemon_id",
			// gender
			"genders.id as gender_id",
			"genders.name as gender_name",
			"genders.icon_url as gender_icon_url",
			// type
			"types.id as type_id",
			"types.name as type_name",
			"types.icon_url as type_icon_url",
			// characteristics
			"characteristics.id as characteristic_id",
			"characteristics.name as characteristic_name",
			"characteristics.description as characteristic_description",
			// description
			"descriptions.id as description_id",
			"descriptions.text as description_text",
			"descriptions.series as description_series",
		).
		Joins(`
			left outer join pokemon_genders on pokemons.id = pokemon_genders.pokemon_id
			left outer join genders on pokemon_genders.gender_id = genders.id
		`).
		Joins(`
			left outer join pokemon_types on pokemons.id = pokemon_types.pokemon_id
			left outer join types on pokemon_types.type_id = types.id
		`).
		Joins(`
			left outer join pokemon_characteristics on pokemons.id = pokemon_characteristics.pokemon_id
			left outer join characteristics on pokemon_characteristics.characteristic_id = characteristics.id
		`).
		Joins("left outer join descriptions on pokemons.id = descriptions.pokemon_id").
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

func (_ *findPokemonCommand) sort(s map[uint]interface{}) *[]interface{} {
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
