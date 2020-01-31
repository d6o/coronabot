package model

type (
	Country string
)

func NewCountry(name string) Country {
	return Country(name)
}
