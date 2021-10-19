package migration

import (
	"fmt"
	"piteroni/dictionary-go-nuxt-graphql/pkg/models"

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

	bulbasaur.EvolutionID = &ivysaur.ID

	if err := db.Save(bulbasaur).Error; err != nil {
		return err
	}

	return nil
}

func createTypes(db *gorm.DB) error {
	entries := map[string]string{
		"ノーマルタイプ":  "/image/icon_type_1_on.svg",
		"ほのおタイプ":   "/image/icon_type_2_on.svg",
		"みずタイプ":    "/image/icon_type_3_on.svg",
		"でんきタイプ":   "/image/icon_type_5_on.svg",
		"くさタイプ":    "/image/icon_type_4_on.svg",
		"こおりタイプ":   "/image/icon_type_6_on.svg",
		"かくとうタイプ":  "/image/icon_type_7_on.svg",
		"どくタイプ":    "/image/icon_type_8_on.svg",
		"じめんタイプ":   "/image/icon_type_9_on.svg",
		"ひこうタイプ":   "/image/icon_type_10_on.svg",
		"エスパータイプ":  "/image/icon_type_11_on.svg",
		"むしタイプ":    "/image/icon_type_12_on.svg",
		"いわタイプ":    "/image/icon_type_13_on.svg",
		"ゴーストタイプ":  "/image/icon_type_14_on.svg",
		"ドラゴンタイプ":  "/image/icon_type_15_on.svg",
		"あくタイプ":    "/image/icon_type_16_on.svg",
		"はがねタイプ":   "/image/icon_type_17_on.svg",
		"フェアリータイプ": "/image/icon_type_18_on.svg",
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

	dao := models.NewPokemonDAO(db)

	for _, name := range []string{"くさタイプ", "どくタイプ"} {
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

	dao := models.NewPokemonDAO(db)

	for _, name := range []string{"くさタイプ", "どくタイプ"} {
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
