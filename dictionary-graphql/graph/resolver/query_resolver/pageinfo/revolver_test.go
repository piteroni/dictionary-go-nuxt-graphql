package pageinfo

import (
	"piteroni/dictionary-go-nuxt-graphql/database"
	graph "piteroni/dictionary-go-nuxt-graphql/graph/model"
	"piteroni/dictionary-go-nuxt-graphql/model"
	itesting "piteroni/dictionary-go-nuxt-graphql/testing"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestPageInfoQueryResolver(t *testing.T) {
	db, err := itesting.ConnnectToInMemoryDatabase()
	if err != nil {
		t.Fatal(err)
	}

	err = database.Migrate(db)
	if err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		err := itesting.RefreshInMemoryDatabase(db)
		if err != nil {
			t.Fatal(err)
		}
	}

	r := &PageInfoQueryResolver{DB: db}

	t.Run("前後にポケモンが登録されているか取得できる", func(t *testing.T) {
		data := []*model.Pokemon{
			{
				Model:      gorm.Model{ID: 1},
				NationalNo: 1,
				Name:       "pokemon-1",
			},
			{
				Model:      gorm.Model{ID: 2},
				NationalNo: 2,
				Name:       "pokemon-2",
			},
		}

		err := db.Create(data).Error
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		t.Run(`
			渡したオブジェクトの前に位置するレコードが存在しない場合、HasPrevはfalseに設定される
			また渡したオブジェクトの次に位置するレコードが存在する場合、HasNextはtrueに設定される
		`, func(t *testing.T) {
			info, err := r.PageInfo(1)
			expected := graph.PageInfo{
				PrevID:  0,
				NextID:  2,
				HasPrev: false,
				HasNext: true,
			}

			assert.Equal(t, expected, info)
			assert.Nil(t, err)
		})

		t.Run(`
			渡したオブジェクトの前に位置するレコードが存在する場合、HasPrevはtrueに設定される
			また渡したオブジェクトの次に位置するレコードが存在しない場合、HasNextはfalseに設定される
		`, func(t *testing.T) {
			info, err := r.PageInfo(2)

			expected := graph.PageInfo{
				PrevID:  1,
				NextID:  3,
				HasPrev: true,
				HasNext: false,
			}

			assert.Equal(t, expected, info)
			assert.Nil(t, err)
		})
	})

	t.Run("指定されたIDに一致するポケモンが存在しない場合、例外が送出される", func(t *testing.T) {
		actual, err := r.PageInfo(1)

		expected := graph.PokemonNotFound{}

		assert.IsType(t, expected, actual)
		assert.Nil(t, err)
	})
}
