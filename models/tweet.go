package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TweetSave struct {
	UserId    string    `bson:"user_id" json:"user_id,omitempty"`
	Message   string    `bson:"message" json:"message,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at,omitempty"`
}

type Tweet struct {
	Message string `bson:"message" json:"message,omitempty"`
}

type Tweets struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId    string             `bson:"user_id" json:"user_id,omitempty"`
	Message   string             `bson:"message" json:"message,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
}

type TweetsFollowing struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId       string             `bson:"user_id" json:"user_id,omitempty"`
	UserRelation string             `bson:"user_relation_id" json:"user_relation_id,omitempty"`
	Tweet        struct {
		ID        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
		Message   string             `bson:"message" json:"message,omitempty"`
		CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
	}
}
