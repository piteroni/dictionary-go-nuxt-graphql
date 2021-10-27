// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

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

type Gender struct {
	Name    string `json:"name"`
	IconURL string `json:"iconURL"`
}

type LinkInfo struct {
	PrevNationalNo int  `json:"prevNationalNo"`
	NextNationalNo int  `json:"nextNationalNo"`
	HasPrev        bool `json:"hasPrev"`
	HasNext        bool `json:"hasNext"`
}

type PageInfo struct {
	HasPreviousPage bool `json:"hasPreviousPage"`
	StartCursor     *int `json:"startCursor"`
	HasNextPage     bool `json:"hasNextPage"`
	EndCursor       *int `json:"endCursor"`
}

type Pokemon struct {
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
	Evolutions      []*Pokemon        `json:"evolutions"`
	LinkInfo        *LinkInfo         `json:"linkInfo"`
}

type PokemonConnection struct {
	PageInfo *PageInfo      `json:"pageInfo"`
	Edges    []*PokemonEdge `json:"edges"`
}

type PokemonEdge struct {
	Cursor string   `json:"cursor"`
	Node   *Pokemon `json:"node"`
}

type Type struct {
	Name    string `json:"name"`
	IconURL string `json:"iconURL"`
}
