package internal

type Product struct {
	ImageUrl string   `bson:"image_url" json:"image_url"`
	Tags     []string `bson:"tags" json:"tags"`
}
