package model

type User struct {
	ID      string `json:"id,omitempty" bson:"_id"`
	userName    string `json:"userName,omitempty" bson:"userName"`
	password    string `json:"password,omitempty" bson:"password"`
	email       string `json:"email,omitempty" bson:"email"`
	phoneNumber string `json:"phoneNumber,omitempty" bson:"phoneNumber"`
}