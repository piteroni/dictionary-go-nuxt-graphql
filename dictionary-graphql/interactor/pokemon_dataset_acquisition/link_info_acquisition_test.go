package pokemon_dataset_acquisition

import (
	"piteroni/dictionary-go-nuxt-graphql/database"
	"piteroni/dictionary-go-nuxt-graphql/model"
	itesting "piteroni/dictionary-go-nuxt-graphql/testing"
	"piteroni/dictionary-go-nuxt-graphql/testing/factories"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestLinkInfoAcquisition(t *testing.T) {
	db, err := itesting.ConnnectToInMemoryDatabase()
	if err != nil {
		t.Fatal(err)
	}

	err = database.Migrate(db)
	if err != nil {
		t.Fatal(err)
	}

	factory := factories.NewPokemonFactory(db)

	cleanup := func() {
		err := itesting.RefreshInMemoryDatabase(db)
		if err != nil {
			t.Fatal(err)
		}
	}

	i := linkInfoAcquisition{db: db}

	t.Run("前後にポケモンが登録されているか取得できる", func(t *testing.T) {
		first := &model.Pokemon{
			Model:      gorm.Model{ID: 1},
			NationalNo: 1,
			Name:       "pokemon-1",
		}

		err = factory.CreatePokemon(first)
		if err != nil {
			t.Fatal(err)
		}

		last := &model.Pokemon{
			Model:      gorm.Model{ID: 2},
			NationalNo: 2,
			Name:       "pokemon-2",
		}

		err = factory.CreatePokemon(last)
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		t.Run(`
			渡したオブジェクトの前に位置するレコードが存在しない場合、HasPrevはfalseに設定される
			また渡したオブジェクトの次に位置するレコードが存在する場合、HasNextはtrueに設定される
		`, func(t *testing.T) {
			linkInfo, err := i.getLinkInfo(first)

			assert.NotNil(t, linkInfo)
			assert.Nil(t, err)

			assert.Equal(t, linkInfo, &LinkInfo{
				PrevID:  0,
				NextID:  2,
				HasPrev: false,
				HasNext: true,
			})
		})

		t.Run(`
			渡したオブジェクトの前に位置するレコードが存在する場合、HasPrevはtrueに設定される
			また渡したオブジェクトの次に位置するレコードが存在しない場合、HasNextはfalseに設定される
		`, func(t *testing.T) {
			linkInfo, err := i.getLinkInfo(last)

			assert.NotNil(t, linkInfo)
			assert.Nil(t, err)

			assert.Equal(t, linkInfo, &LinkInfo{
				PrevID:  1,
				NextID:  3,
				HasPrev: true,
				HasNext: false,
			})
		})
	})
}
