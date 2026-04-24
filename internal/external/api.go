package external

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Genderize struct {
	Gender string  `json:"gender"`
	Prob   float64 `json:"probability"`
	Count  int     `json:"count"`
}

type Agify struct {
	Age int `json:"age"`
}

type Country struct {
	CountryID string  `json:"country_id"`
	Prob      float64 `json:"probability"`
}

type Nationalize struct {
	Country []Country `json:"country"`
}

func FetchAll(name string) (*Genderize, *Agify, *Country, error) {

	gRes, _ := http.Get("https://api.genderize.io?name=" + name)
	var g Genderize
	_ = json.NewDecoder(gRes.Body).Decode(&g)

	if g.Gender == "" {
		return nil, nil, nil, fmt.Errorf("Genderize")
	}

	aRes, _ := http.Get("https://api.agify.io?name=" + name)
	var a Agify
	_ = json.NewDecoder(aRes.Body).Decode(&a)

	if a.Age == 0 {
		return nil, nil, nil, fmt.Errorf("Agify")
	}

	nRes, _ := http.Get("https://api.nationalize.io?name=" + name)
	var n Nationalize
	_ = json.NewDecoder(nRes.Body).Decode(&n)

	if len(n.Country) == 0 {
		return nil, nil, nil, fmt.Errorf("Nationalize")
	}

	best := n.Country[0]
	for _, c := range n.Country {
		if c.Prob > best.Prob {
			best = c
		}
	}

	return &g, &a, &best, nil
}
