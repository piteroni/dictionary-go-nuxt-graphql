package migration

import (
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/pkg/models"
	"piteroni/dictionary-go-nuxt-graphql/pkg/persistence"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	if err := createTypes(db); err != nil {
		return err
	}

	if err := createGenders(db); err != nil {
		return err
	}

	if err := createCharacteristics(db); err != nil {
		return err
	}

	bulbasaur, err := createBulbasaur(db)
	if err != nil {
		return err
	}

	ivysaur, err := createIvysaur(db)
	if err != nil {
		return err
	}

	venusaur, err := createVenusaur(db)
	if err != nil {
		return err
	}

	bulbasaur.EvolutionID = &ivysaur.ID

	if err := db.Save(bulbasaur).Error; err != nil {
		return err
	}

	ivysaur.EvolutionID = &venusaur.ID

	if err := db.Save(ivysaur).Error; err != nil {
		return err
	}

	return nil
}

func createTypes(db *gorm.DB) error {
	entries := map[string]string{
		"ノーマル":  "/image/icon_type_1_on.svg",
		"ほのお":   "/image/icon_type_2_on.svg",
		"みず":    "/image/icon_type_3_on.svg",
		"でんき":   "/image/icon_type_5_on.svg",
		"くさ":    "/image/icon_type_4_on.svg",
		"こおり":   "/image/icon_type_6_on.svg",
		"かくとう":  "/image/icon_type_7_on.svg",
		"どく":    "/image/icon_type_8_on.svg",
		"じめん":   "/image/icon_type_9_on.svg",
		"ひこう":   "/image/icon_type_10_on.svg",
		"エスパー":  "/image/icon_type_11_on.svg",
		"むし":    "/image/icon_type_12_on.svg",
		"いわ":    "/image/icon_type_13_on.svg",
		"ゴースト":  "/image/icon_type_14_on.svg",
		"ドラゴン":  "/image/icon_type_15_on.svg",
		"あく":    "/image/icon_type_16_on.svg",
		"はがね":   "/image/icon_type_17_on.svg",
		"フェアリー": "/image/icon_type_18_on.svg",
	}

	for name, icon := range entries {
		t := &models.Type{
			Name:    name,
			IconURL: icon,
		}

		if err := db.Create(t).Error; err != nil {
			return err
		}
	}

	return nil
}

func createGenders(db *gorm.DB) error {
	entries := map[string]string{
		"male":   "/image/icon_male.svg",
		"female": "/image/icon_female.svg",
	}

	for name, icon := range entries {
		g := &models.Gender{
			Name:    name,
			IconURL: icon,
		}

		if err := db.Create(g).Error; err != nil {
			return err
		}
	}

	return nil
}

func createCharacteristics(db *gorm.DB) error {
	entries := map[string]string{
		"しんりょく":  `ＨＰが　へったとき　くさタイプの　わざの　いりょくが　あがる。`,
		"あついしぼう": `あつい　しぼうで　まもられているので　ほのおタイプと　こおりタイプの　わざの　ダメージを　はんげんさせる。`,
		"もうか":    `ＨＰが　へったとき　ほのおタイプの　わざの　いりょくが　あがる。`,
	}

	for name, description := range entries {
		c := &models.Characteristic{
			Name:        name,
			Description: description,
		}

		if err := db.Create(c).Error; err != nil {
			return err
		}
	}

	return nil
}

func createBulbasaur(db *gorm.DB) (*models.Pokemon, error) {
	pokemon := &models.Pokemon{
		NationalNo:          1,
		Name:                "フシギダネ",
		Species:             "たねポケモン",
		ImageURL:            "/image/afa02eaba4c39820fc57f4e8abaeea80.png",
		Height:              "0.7m",
		Weight:              "6.9kg",
		HeartPoint:          45,
		AttackPoint:         49,
		DefensePoint:        49,
		SpecialAttachPoint:  65,
		SpecialDefensePoint: 65,
		SpeedPoint:          45,
	}

	if err := db.Create(pokemon).Error; err != nil {
		return nil, err
	}

	dao := persistence.NewPokemonDAO(db)

	for _, name := range []string{"くさ", "どく"} {
		t := &models.Type{}

		if err := db.Model(&models.Type{}).Where(fmt.Sprintf("name = '%s'", name)).First(t).Error; err != nil {
			return nil, err
		}

		if err := dao.AddType(pokemon, t); err != nil {
			return nil, err
		}
	}

	for _, name := range []string{"female", "male"} {
		g := &models.Gender{}

		if err := db.Model(&models.Gender{}).Where(fmt.Sprintf("name = '%s'", name)).First(g).Error; err != nil {
			return nil, err
		}

		if err := dao.AddGender(pokemon, g); err != nil {
			return nil, err
		}
	}

	descriptions := map[string]string{
		"うまれたときから　せなかに　しょくぶつの　タネが　あって　すこしずつ　おおきく　そだつ。": "ポケモン ソード",
		"うまれて　しばらくの　あいだ　せなかの　タネに　つまった　えいようを　とって　そだつ。":  "ポケモン シールド",
	}

	for description, series := range descriptions {
		d := &models.Description{
			Text:   description,
			Series: series,
		}

		if err := dao.AddDescripton(pokemon, d); err != nil {
			return nil, err
		}
	}

	for _, name := range []string{"しんりょく"} {
		c := &models.Characteristic{}

		if err := db.Model(&models.Characteristic{}).Where(fmt.Sprintf("name = '%s'", name)).First(c).Error; err != nil {
			return nil, err
		}

		if err := dao.AddCharacteristics(pokemon, c); err != nil {
			return nil, err
		}
	}

	return pokemon, nil
}

func createIvysaur(db *gorm.DB) (*models.Pokemon, error) {
	pokemon := &models.Pokemon{
		NationalNo:          2,
		Name:                "フシギソウ",
		Species:             "たねポケモン",
		ImageURL:            "/image/6f8144eb4659537733b930d6a299d5a7.png",
		Height:              "1.0m",
		Weight:              "13.0kg",
		HeartPoint:          60,
		AttackPoint:         62,
		DefensePoint:        63,
		SpecialAttachPoint:  80,
		SpecialDefensePoint: 80,
		SpeedPoint:          60,
	}

	if err := db.Create(pokemon).Error; err != nil {
		return nil, err
	}

	dao := persistence.NewPokemonDAO(db)

	for _, name := range []string{"くさ", "どく"} {
		t := &models.Type{}

		if err := db.Model(&models.Type{}).Where(fmt.Sprintf("name = '%s'", name)).First(t).Error; err != nil {
			return nil, err
		}

		if err := dao.AddType(pokemon, t); err != nil {
			return nil, err
		}
	}

	for _, name := range []string{"female", "male"} {
		g := &models.Gender{}

		if err := db.Model(&models.Gender{}).Where(fmt.Sprintf("name = '%s'", name)).First(g).Error; err != nil {
			return nil, err
		}

		if err := dao.AddGender(pokemon, g); err != nil {
			return nil, err
		}
	}

	descriptions := map[string]string{
		"せなかの　つぼみが　おおきく　そだってくると　２ほんあしで　たつことが　できなくなるらしい。":   "ポケモン ソード",
		"たいようの　ひかりを　あびるほど　からだに　ちからが　わいて　せなかの　つぼみが　そだっていく。": "ポケモン シールド",
	}

	for description, series := range descriptions {
		d := &models.Description{
			Text:   description,
			Series: series,
		}

		if err := dao.AddDescripton(pokemon, d); err != nil {
			return nil, err
		}
	}

	for _, name := range []string{"しんりょく"} {
		c := &models.Characteristic{}

		if err := db.Model(&models.Characteristic{}).Where(fmt.Sprintf("name = '%s'", name)).First(c).Error; err != nil {
			return nil, err
		}

		if err := dao.AddCharacteristics(pokemon, c); err != nil {
			return nil, err
		}
	}

	return pokemon, nil
}

func createVenusaur(db *gorm.DB) (*models.Pokemon, error) {
	pokemon := &models.Pokemon{
		NationalNo:          3,
		Name:                "フシギバナ",
		Species:             "たねポケモン",
		ImageURL:            "/image/ebccfe6f2ccfe2e851fd29739bf6220c.png",
		Height:              "2.0m",
		Weight:              "100.0kg",
		HeartPoint:          80,
		AttackPoint:         82,
		DefensePoint:        83,
		SpecialAttachPoint:  100,
		SpecialDefensePoint: 100,
		SpeedPoint:          80,
	}

	err := db.Create(pokemon).Error
	if err != nil {
		return nil, err
	}

	dao := persistence.NewPokemonDAO(db)

	for _, name := range []string{"くさ", "どく"} {
		t := &models.Type{}

		if err := db.Model(&models.Type{}).Where(fmt.Sprintf("name = '%s'", name)).First(t).Error; err != nil {
			return nil, err
		}

		if err := dao.AddType(pokemon, t); err != nil {
			return nil, err
		}
	}

	for _, name := range []string{"female", "male"} {
		g := &models.Gender{}

		if err := db.Model(&models.Gender{}).Where(fmt.Sprintf("name = '%s'", name)).First(g).Error; err != nil {
			return nil, err
		}

		if err := dao.AddGender(pokemon, g); err != nil {
			return nil, err
		}
	}

	descriptions := map[string]string{
		"たいようエネルギーを　えいようにして　おおきなハナが　ひらく。　ひなたに　ひきよせられるように　いどうする。": "ポケモン ソード",
		"はなから　うっとりする　かおりが　ただよい　たたかうものの　きもちを　なだめてしまう。":            "ポケモン シールド",
	}

	for description, series := range descriptions {
		d := &models.Description{
			Text:   description,
			Series: series,
		}

		if err := dao.AddDescripton(pokemon, d); err != nil {
			return nil, err
		}
	}

	for _, name := range []string{"しんりょく"} {
		c := &models.Characteristic{}

		if err := db.Model(&models.Characteristic{}).Where(fmt.Sprintf("name = '%s'", name)).First(c).Error; err != nil {
			return nil, err
		}

		if err := dao.AddCharacteristics(pokemon, c); err != nil {
			return nil, err
		}
	}

	return pokemon, nil
}
