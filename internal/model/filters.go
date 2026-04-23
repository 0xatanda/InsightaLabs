package model

type Filters struct {
	Gender    string
	AgeGroup  string
	CountryID string
	MinAge    *int
	MaxAge    *int
	SortBy    string
	Order     string
	Page      int
	Limit     int
}
