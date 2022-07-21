package models

import "time"

type TweetSave struct {
	UserId    string    `bson:"user_id" json:"user_id,omitempty"`
	Message   string    `bson:"message" json:"message,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at,omitempty"`
}

type Tweet struct {
	Message string `bson:"message" json:"message,omitempty"`
}
