package domain

type User struct {
	Id   string `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	Sex  int    `bson:"sex" json:"sex"`
}
