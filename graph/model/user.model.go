package model

type User struct {
	ID        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	Email     string `json:"email" bson:"email"`
	Phone     string `json:"phone" bson:"phone"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
}
