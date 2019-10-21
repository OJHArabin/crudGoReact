package model

type UserDetails struct {
	Id       string `bson:"id" json:"id"`
	Password string `bson:"password" json:"password"`
	UserName string `bson:"username" json:"username"`
	FullName string `bson:"fullname" json:"fullname"`
	Email    string `bson:"email" json:"email"`
}
