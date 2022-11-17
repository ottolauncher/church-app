package model

type User struct {
	ID        string `json:"id" bson:"_id"`
	Username  string `json:"username" bson:"username"`
	Name      string `json:"name" bson:"name"`
	Email     string `json:"email" bson:"email"`
	Phone     string `json:"phone" bson:"phone"`
	Password  []byte `json:"-" bson:"-"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
