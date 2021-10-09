package models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)


type Users struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User_id  string 		  `json:"User_id" bson:"User_id"`
	Name   string             `json:"Name" bson:"Name"`
	Email  string             `json:"Email" bson:"Email"`
	Password string           `json:"Password" bson:"Password"`
}

type Post struct {
	ID     primitive.ObjectID    `json:"_id,omitempty" bson:"_id,omitempty"`
	User_id  string 		  `json:"User_id" bson:"User_id"`
	Caption   string             `json:"Caption" bson:"Caption"`
	Image_url string             `json:"Image" bson:"Image"`
	TimeStamp time.Time 		 `json:"Time" bson:"Time"`
}
