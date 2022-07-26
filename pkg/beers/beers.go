package beers

type Beer struct {
	Id             int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string  `json:"name"`
	Ingredients    string  `json:"ingredients"`
	AlcoholContent string  `json:"alcoholContent"`
	Price          float64 `json:"price"`
	Category       string  `json:"category"`
}
