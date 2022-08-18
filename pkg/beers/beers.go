package beers

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Beer struct {
	Id             int     `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	Name           string  `json:"name,omitempty"`
	Ingredients    string  `json:"ingredients,omitempty"`
	AlcoholContent string  `json:"alcoholContent,omitempty"`
	Price          float64 `json:"price,omitempty"`
	Category       string  `json:"category,omitempty"`
}

func (b Beer) Validate() error {
	return validation.ValidateStruct(&b,

		validation.Field(&b.Name, validation.Required, validation.Length(1, 50)),

		validation.Field(&b.Ingredients, validation.Required, validation.Length(1, 50)),

		validation.Field(&b.AlcoholContent, validation.Required, validation.Length(1, 50)),

		validation.Field(&b.Price, validation.Required, validation.Min(float64(1)), validation.Max(float64(100))),

		validation.Field(&b.Category, validation.Required, validation.Length(1, 50)),
	)
}
