package model

type Filters struct {
	Gender                string
	AgeGroup              string
	CountryID             string
	MinAge                *int
	MaxAge                *int
	MinGenderProbability  *float64
	MinCountryProbability *float64
	SortBy                string
	Order                 string
	Page                  int
	Limit                 int
}
