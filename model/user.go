package model

type User struct {
	ID      string `json:"id,omitempty" bson:"_id"`
	UserName    string `json:"userName,omitempty" bson:"userName"`
	Password    string `json:"password,omitempty" bson:"password"`
	Email       string `json:"email,omitempty" bson:"email"`
	PhoneNumber string `json:"phoneNumber,omitempty" bson:"phoneNumber"`
}