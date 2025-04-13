package model

type Product struct {
	ID    string  `json:"id" bson:"_id,omitempty"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
