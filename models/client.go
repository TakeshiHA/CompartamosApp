package models

type Client struct {
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	DNI       string `json:"dni,omitempty" bson:"dni,omitempty"`
	FirstName string `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Email     string `json:"email,omitempty" bson:"email,omitempty"`
	Phone     string `json:"phone,omitempty" bson:"phone,omitempty"`
	CityID    string `json:"cityId,omitempty" bson:"cityId,omitempty"`
}
