package document

type Gender struct {
	Record
	Name    string `bson:"name"`
	IconURL string `bson:"icon_url"`
}
