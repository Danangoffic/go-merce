package mongo

type Vouchers struct {
	Title string `bson:"title"`
	Price int    `bson:"price"`
	Type  string `bson:"type"`
}
