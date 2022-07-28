package beers

type Beer struct {
	Id             int     `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	Name           string  `json:"name,omitempty"`
	Ingredients    string  `json:"ingredients,omitempty"`
	AlcoholContent string  `json:"alcoholContent,omitempty"`
	Price          float64 `json:"price,omitempty"`
	Category       string  `json:"category,omitempty"`
}
