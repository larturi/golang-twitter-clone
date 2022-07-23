package models

type Relation struct {
	UserID         string `bson:"user_id,omitempty" json:"userId"`
	UserRelationID string `bson:"user_relation_id,omitempty" json:"userRelationId"`
}

type ResponseRelation struct {
	Status bool `json:"status"`
}
