package database

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

func Seed(ctx context.Context, db *mongo.Database) error {
	err := createTypes(ctx, db)
	if err != nil {
		return err
	}

	err = createGenders(ctx, db)
	if err != nil {
		return err
	}

	err = createCharacteristics(ctx, db)
	if err != nil {
		return err
	}

	err = createBulbasaur(ctx, db)
	if err != nil {
		return err
	}

	err = createIvysaur(ctx, db)
	if err != nil {
		return err
	}

	err = createVenusaur(ctx, db)
	if err != nil {
		return err
	}

	err = createDarkrai(ctx, db)
	if err != nil {
		return err
	}

	err = updateEvolutionaryRelation(ctx, db)
	if err != nil {
		return err
	}

	err = createPaginationEdges(ctx, db)
	if err != nil {
		return err
	}

	return nil
}

func createTypes(ctx context.Context, db *mongo.Database) error {
	types := []interface{}{
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "ノーマル",
			IconURL: "/image/icon_type_1_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "ほのお",
			IconURL: "/image/icon_type_2_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "みず",
			IconURL: "/image/icon_type_3_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "でんき",
			IconURL: "/image/icon_type_5_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "くさ",
			IconURL: "/image/icon_type_4_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "こおり",
			IconURL: "/image/icon_type_6_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "かくとう",
			IconURL: "/image/icon_type_7_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "どく",
			IconURL: "/image/icon_type_8_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "じめん",
			IconURL: "/image/icon_type_9_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "ひこう",
			IconURL: "/image/icon_type_10_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "エスパー",
			IconURL: "/image/icon_type_11_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "むし",
			IconURL: "/image/icon_type_12_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "いわ",
			IconURL: "/image/icon_type_13_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "ゴースト",
			IconURL: "/image/icon_type_14_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "ドラゴン",
			IconURL: "/image/icon_type_15_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "あく",
			IconURL: "/image/icon_type_16_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "はがね",
			IconURL: "/image/icon_type_17_on.svg",
		},
		document.Type{
			ID:      primitive.NewObjectID(),
			Name:    "フェアリー",
			IconURL: "/image/icon_type_18_on.svg",
		},
	}

	_, err := db.Collection(collection.Types).InsertMany(ctx, types)
	if err != nil {
		return err
	}

	return nil
}

func createGenders(ctx context.Context, db *mongo.Database) error {
	genders := []interface{}{
		document.Gender{
			ID:      primitive.NewObjectID(),
			Name:    "male",
			IconURL: "/image/icon_male.svg",
		},
		document.Gender{
			ID:      primitive.NewObjectID(),
			Name:    "female",
			IconURL: "/image/icon_female.svg",
		},
	}

	_, err := db.Collection(collection.Genders).InsertMany(ctx, genders)
	if err != nil {
		return err
	}

	return nil
}

func createCharacteristics(ctx context.Context, db *mongo.Database) error {
	characteristics := []interface{}{
		document.Characteristic{
			ID:          primitive.NewObjectID(),
			Name:        "しんりょく",
			Description: "ＨＰが　へったとき　くさタイプの　わざの　いりょくが　あがる。",
		},
		document.Characteristic{
			ID:          primitive.NewObjectID(),
			Name:        "あついしぼう",
			Description: "あつい　しぼうで　まもられているので　ほのおタイプと　こおりタイプの　わざの　ダメージを　はんげんさせる。",
		},
		document.Characteristic{
			ID:          primitive.NewObjectID(),
			Name:        "ナイトメア",
			Description: "ねむり　じょうたいの　あいてにダメージを　あたえる。",
		},
	}

	_, err := db.Collection(collection.Characteristics).InsertMany(ctx, characteristics)
	if err != nil {
		return err
	}

	return nil
}

func createBulbasaur(ctx context.Context, db *mongo.Database) error {
	// acquire relation types.
	condition := bson.M{"name": bson.D{{Key: "$in", Value: []string{"くさ", "どく"}}}}
	opt := options.FindOptions{Projection: bson.M{"_id": 1}}

	cursor, err := db.Collection(collection.Types).Find(ctx, condition, &opt)
	if err != nil {
		return errors.WithStack(err)
	}

	typeIds := []primitive.ObjectID{}
	for cursor.Next(ctx) {
		typeIds = append(typeIds, cursor.Current.Lookup("_id").ObjectID())
	}

	err = cursor.Close(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	// acquire relation genders.
	condition = bson.M{"name": bson.D{{Key: "$in", Value: []string{"male", "female"}}}}
	opt = options.FindOptions{Projection: bson.M{"_id": 1}}

	cursor, err = db.Collection(collection.Genders).Find(ctx, condition, &opt)
	if err != nil {
		return errors.WithStack(err)
	}

	genderIds := []primitive.ObjectID{}
	for cursor.Next(ctx) {
		genderIds = append(genderIds, cursor.Current.Lookup("_id").ObjectID())
	}

	err = cursor.Close(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	// acquire relation characteristics.
	condition = bson.M{"name": bson.D{{Key: "$in", Value: []string{"しんりょく"}}}}
	opt = options.FindOptions{Projection: bson.M{"_id": 1}}

	cursor, err = db.Collection(collection.Characteristics).Find(ctx, condition, &opt)
	if err != nil {
		return errors.WithStack(err)
	}

	characteristicIds := []primitive.ObjectID{}
	for cursor.Next(ctx) {
		characteristicIds = append(characteristicIds, cursor.Current.Lookup("_id").ObjectID())
	}

	err = cursor.Close(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	descriptions := []document.Description{
		{
			ID:     primitive.NewObjectID(),
			Text:   "うまれたときから　せなかに　しょくぶつの　タネが　あって　すこしずつ　おおきく　そだつ。",
			Series: "ポケモン ソード",
		},
		{
			ID:     primitive.NewObjectID(),
			Text:   "うまれて　しばらくの　あいだ　せなかの　タネに　つまった　えいようを　とって　そだつ。",
			Series: "ポケモン シールド",
		},
	}

	bulbasaur := document.Pokemon{
		ID:                  primitive.NewObjectID(),
		NationalNo:          1,
		Name:                "フシギダネ",
		Species:             "たねポケモン",
		ImageURL:            "/image/afa02eaba4c39820fc57f4e8abaeea80.png",
		Height:              "0.7m",
		Weight:              "6.9kg",
		HeartPoint:          45,
		AttackPoint:         49,
		DefensePoint:        49,
		SpecialAttackPoint:  65,
		SpecialDefensePoint: 65,
		SpeedPoint:          45,
		Descriptions:        descriptions,
		References: document.PokemonReferences{
			Types:           typeIds,
			Genders:         genderIds,
			Characteristics: characteristicIds,
		},
	}

	_, err = db.Collection(collection.Pokemons).InsertOne(ctx, bulbasaur)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func createIvysaur(ctx context.Context, db *mongo.Database) error {
	// acquire relation types.
	condition := bson.M{"name": bson.D{{Key: "$in", Value: []string{"くさ", "どく"}}}}
	opt := options.FindOptions{Projection: bson.M{"_id": 1}}

	cursor, err := db.Collection(collection.Types).Find(ctx, condition, &opt)
	if err != nil {
		return errors.WithStack(err)
	}

	typeIds := []primitive.ObjectID{}
	for cursor.Next(ctx) {
		typeIds = append(typeIds, cursor.Current.Lookup("_id").ObjectID())
	}

	err = cursor.Close(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	// acquire relation genders.
	condition = bson.M{"name": bson.D{{Key: "$in", Value: []string{"male", "female"}}}}
	opt = options.FindOptions{Projection: bson.M{"_id": 1}}

	cursor, err = db.Collection(collection.Genders).Find(ctx, condition, &opt)
	if err != nil {
		return errors.WithStack(err)
	}

	genderIds := []primitive.ObjectID{}
	for cursor.Next(ctx) {
		genderIds = append(genderIds, cursor.Current.Lookup("_id").ObjectID())
	}

	err = cursor.Close(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	// acquire relation characteristics.
	condition = bson.M{"name": bson.D{{Key: "$in", Value: []string{"しんりょく"}}}}
	opt = options.FindOptions{Projection: bson.M{"_id": 1}}

	cursor, err = db.Collection(collection.Characteristics).Find(ctx, condition, &opt)
	if err != nil {
		return errors.WithStack(err)
	}

	characteristicIds := []primitive.ObjectID{}
	for cursor.Next(ctx) {
		characteristicIds = append(characteristicIds, cursor.Current.Lookup("_id").ObjectID())
	}

	err = cursor.Close(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	descriptions := []document.Description{
		{
			ID:     primitive.NewObjectID(),
			Text:   "せなかの　つぼみが　おおきく　そだってくると　２ほんあしで　たつことが　できなくなるらしい。",
			Series: "ポケモン ソード",
		},
		{
			ID:     primitive.NewObjectID(),
			Text:   "たいようの　ひかりを　あびるほど　からだに　ちからが　わいて　せなかの　つぼみが　そだっていく。",
			Series: "ポケモン シールド",
		},
	}

	ivysaur := document.Pokemon{
		ID:                  primitive.NewObjectID(),
		NationalNo:          2,
		Name:                "フシギソウ",
		Species:             "たねポケモン",
		ImageURL:            "/image/6f8144eb4659537733b930d6a299d5a7.png",
		Height:              "1.0m",
		Weight:              "13.0kg",
		HeartPoint:          60,
		AttackPoint:         62,
		DefensePoint:        63,
		SpecialAttackPoint:  80,
		SpecialDefensePoint: 80,
		SpeedPoint:          60,
		Descriptions:        descriptions,
		References: document.PokemonReferences{
			Types:           typeIds,
			Genders:         genderIds,
			Characteristics: characteristicIds,
		},
	}

	_, err = db.Collection(collection.Pokemons).InsertOne(ctx, ivysaur)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func createVenusaur(ctx context.Context, db *mongo.Database) error {
	// acquire relation types.
	condition := bson.M{"name": bson.D{{Key: "$in", Value: []string{"くさ", "どく"}}}}
	opt := options.FindOptions{Projection: bson.M{"_id": 1}}

	cursor, err := db.Collection(collection.Types).Find(ctx, condition, &opt)
	if err != nil {
		return errors.WithStack(err)
	}

	typeIds := []primitive.ObjectID{}
	for cursor.Next(ctx) {
		typeIds = append(typeIds, cursor.Current.Lookup("_id").ObjectID())
	}

	err = cursor.Close(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	// acquire relation genders.
	condition = bson.M{"name": bson.D{{Key: "$in", Value: []string{"male", "female"}}}}
	opt = options.FindOptions{Projection: bson.M{"_id": 1}}

	cursor, err = db.Collection(collection.Genders).Find(ctx, condition, &opt)
	if err != nil {
		return errors.WithStack(err)
	}

	genderIds := []primitive.ObjectID{}
	for cursor.Next(ctx) {
		genderIds = append(genderIds, cursor.Current.Lookup("_id").ObjectID())
	}

	err = cursor.Close(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	// acquire relation characteristics.
	condition = bson.M{"name": bson.D{{Key: "$in", Value: []string{"しんりょく"}}}}
	opt = options.FindOptions{Projection: bson.M{"_id": 1}}

	cursor, err = db.Collection(collection.Characteristics).Find(ctx, condition, &opt)
	if err != nil {
		return errors.WithStack(err)
	}

	characteristicIds := []primitive.ObjectID{}
	for cursor.Next(ctx) {
		characteristicIds = append(characteristicIds, cursor.Current.Lookup("_id").ObjectID())
	}

	err = cursor.Close(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	descriptions := []document.Description{
		{
			ID:     primitive.NewObjectID(),
			Text:   "たいようエネルギーを　えいようにして　おおきなハナが　ひらく。　ひなたに　ひきよせられるように　いどうする。",
			Series: "ポケモン ソード",
		},
		{
			ID:     primitive.NewObjectID(),
			Text:   "はなから　うっとりする　かおりが　ただよい　たたかうものの　きもちを　なだめてしまう。",
			Series: "ポケモン シールド",
		},
	}

	ivysaur := document.Pokemon{
		ID:                  primitive.NewObjectID(),
		NationalNo:          3,
		Name:                "フシギバナ",
		Species:             "たねポケモン",
		ImageURL:            "/image/ebccfe6f2ccfe2e851fd29739bf6220c.png",
		Height:              "2.0m",
		Weight:              "100.0kg",
		HeartPoint:          80,
		AttackPoint:         82,
		DefensePoint:        83,
		SpecialAttackPoint:  100,
		SpecialDefensePoint: 100,
		SpeedPoint:          80,
		Descriptions:        descriptions,
		References: document.PokemonReferences{
			Types:           typeIds,
			Genders:         genderIds,
			Characteristics: characteristicIds,
		},
	}

	_, err = db.Collection(collection.Pokemons).InsertOne(ctx, ivysaur)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func createDarkrai(ctx context.Context, db *mongo.Database) error {
	// acquire relation characteristics.
	condition := bson.M{"name": bson.D{{Key: "$in", Value: []string{"ナイトメア"}}}}
	opt := options.FindOptions{Projection: bson.M{"_id": 1}}

	cursor, err := db.Collection(collection.Characteristics).Find(ctx, condition, &opt)
	if err != nil {
		return errors.WithStack(err)
	}

	err = cursor.Close(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	characteristicIds := []primitive.ObjectID{}
	for cursor.Next(ctx) {
		characteristicIds = append(characteristicIds, cursor.Current.Lookup("_id").ObjectID())
	}

	descriptions := []document.Description{
		{
			ID:     primitive.NewObjectID(),
			Text:   "ひとびとを　ふかい　ねむりに　さそいゆめを　みせる　のうりょくを　もつ。　しんげつの　よるに　かつどうする。",
			Series: "ポケモン Y",
		},
	}

	darkrai := document.Pokemon{
		ID:                  primitive.NewObjectID(),
		NationalNo:          491,
		Name:                "ダークライ",
		Species:             "あんこくポケモン",
		ImageURL:            "/image/58d8af4d8c18b2aad5c524ec90e3dbe7.png",
		Height:              "1.5m",
		Weight:              "50.5kg",
		HeartPoint:          70,
		AttackPoint:         90,
		DefensePoint:        90,
		SpecialAttackPoint:  135,
		SpecialDefensePoint: 90,
		SpeedPoint:          125,
		Descriptions:        descriptions,
		References: document.PokemonReferences{
			Characteristics: characteristicIds,
		},
	}

	_, err = db.Collection(collection.Pokemons).InsertOne(ctx, darkrai)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func updateEvolutionaryRelation(ctx context.Context, db *mongo.Database) error {
	// bulbasaur -> ivysaur.
	condition := bson.M{"name": "フシギソウ"}
	eopt := options.FindOneOptions{Projection: bson.M{"_id": 1}}
	pokemon := map[string]interface{}{}

	err := db.Collection(collection.Pokemons).FindOne(ctx, condition, &eopt).Decode(&pokemon)
	if err != nil {
		return errors.WithStack(err)
	}

	condition = bson.M{"name": "フシギダネ"}
	update := bson.M{"$set": bson.M{"evolution_id": pokemon["_id"].(primitive.ObjectID)}}

	r := db.Collection(collection.Pokemons).FindOneAndUpdate(ctx, condition, update)
	if r.Err() != nil {
		return errors.WithStack(r.Err())
	}

	// ivysaur -> venusaur.
	condition = bson.M{"name": "フシギバナ"}
	eopt = options.FindOneOptions{Projection: bson.M{"_id": 1}}
	pokemon = map[string]interface{}{}

	err = db.Collection(collection.Pokemons).FindOne(ctx, condition, &eopt).Decode(&pokemon)
	if err != nil {
		return errors.WithStack(err)
	}

	condition = bson.M{"name": "フシギソウ"}
	update = bson.M{"$set": bson.M{"evolution_id": pokemon["_id"].(primitive.ObjectID)}}

	r = db.Collection(collection.Pokemons).FindOneAndUpdate(ctx, condition, update)
	if r.Err() != nil {
		return errors.WithStack(r.Err())
	}

	return nil
}

func createPaginationEdges(ctx context.Context, db *mongo.Database) error {
	pokemons := []interface{}{}
	max := 192 // 64 * 3

	for i := 0; i < max; i++ {
		pokemon := document.Pokemon{
			ID:                  primitive.NewObjectID(),
			NationalNo:          90000 + i,
			Name:                fmt.Sprintf("pagenation edge %d", i),
			Species:             "",
			ImageURL:            "/image/question.svg",
			Height:              "??",
			Weight:              "??",
			HeartPoint:          40,
			AttackPoint:         40,
			DefensePoint:        40,
			SpecialAttackPoint:  40,
			SpecialDefensePoint: 40,
			SpeedPoint:          40,
		}

		pokemons = append(pokemons, pokemon)
	}

	_, err := db.Collection(collection.Pokemons).InsertMany(ctx, pokemons)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
