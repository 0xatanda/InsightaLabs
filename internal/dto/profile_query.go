package dto

// ProfileQuery is the internal structured representation
// of both API filters and parsed NLQ filters.
type ProfileQuery struct {
	Gender   string
	Country  string
	AgeGroup string

	MinAge *int
	MaxAge *int

	MinGenderProbability  *float64
	MinCountryProbability *float64

	SortBy string
	Order  string

	Page  int
	Limit int
}
