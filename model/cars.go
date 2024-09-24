package model

type Car struct {
	ID         string `json:"id"`
	Brand      string `json:"brand"`
	Colour     string `json:"colour"`
	HorsePower int    `json:"horsepower"`
}

var Database = map[string]Car{}
