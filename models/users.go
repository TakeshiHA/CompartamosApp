package models

type User struct {
	ID    string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string `json:"fullname,omitempty" bson:"fullname,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Phone string `json:"phone,omitempty" bson:"phone,omitempty"`
}
