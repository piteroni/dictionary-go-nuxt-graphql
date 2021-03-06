// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type EvolutionsResult interface {
	IsEvolutionsResult()
}

type PageInfoResult interface {
	IsPageInfoResult()
}

type PokemonConnectionResult interface {
	IsPokemonConnectionResult()
}

type PokemonResult interface {
	IsPokemonResult()
}

type Ability struct {
	Heart          int `json:"heart"`
	Attack         int `json:"attack"`
	Defense        int `json:"defense"`
	SpecialAttack  int `json:"specialAttack"`
	SpecialDefense int `json:"specialDefense"`
	Speed          int `json:"speed"`
}

type Characteristic struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Description struct {
	Text   string `json:"text"`
	Series string `json:"series"`
}

type Evolutions struct {
	Pokemons []*Pokemon `json:"pokemons"`
}

func (Evolutions) IsEvolutionsResult() {}

type Gender struct {
	Name    string `json:"name"`
	IconURL string `json:"iconURL"`
}

type IllegalArguments struct {
	Message string `json:"message"`
}

func (IllegalArguments) IsPokemonResult()           {}
func (IllegalArguments) IsEvolutionsResult()        {}
func (IllegalArguments) IsPageInfoResult()          {}
func (IllegalArguments) IsPokemonConnectionResult() {}

type PageInfo struct {
	PrevID  string `json:"prevId"`
	NextID  string `json:"nextId"`
	HasPrev bool   `json:"hasPrev"`
	HasNext bool   `json:"hasNext"`
}

func (PageInfo) IsPageInfoResult() {}

type Pokemon struct {
	ID              string            `json:"id"`
	NationalNo      int               `json:"nationalNo"`
	Name            string            `json:"name"`
	ImageURL        string            `json:"imageURL"`
	Species         string            `json:"species"`
	Height          string            `json:"height"`
	Weight          string            `json:"weight"`
	Genders         []*Gender         `json:"genders"`
	Types           []*Type           `json:"types"`
	Characteristics []*Characteristic `json:"characteristics"`
	Description     *Description      `json:"description"`
	Ability         *Ability          `json:"ability"`
	CanEvolution    bool              `json:"canEvolution"`
}

func (Pokemon) IsPokemonResult() {}

type PokemonConnection struct {
	EndCursor string     `json:"endCursor"`
	HasNext   bool       `json:"hasNext"`
	Items     []*Pokemon `json:"items"`
}

func (PokemonConnection) IsPokemonConnectionResult() {}

type PokemonNotFound struct {
	Message string `json:"message"`
}

func (PokemonNotFound) IsPokemonResult()           {}
func (PokemonNotFound) IsEvolutionsResult()        {}
func (PokemonNotFound) IsPageInfoResult()          {}
func (PokemonNotFound) IsPokemonConnectionResult() {}

type Type struct {
	Name    string `json:"name"`
	IconURL string `json:"iconURL"`
}
