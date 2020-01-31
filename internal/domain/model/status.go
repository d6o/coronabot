package model

import (
	"strings"
	"time"
)

type (
	Status struct {
		country    string
		deaths     int
		confirmed  int
		recovered  int
		updateTime time.Time
	}

	StatusReporter interface {
		StatusPerCountry() ([]Status, error)
	}
)

func NewStatus(country string, deaths int, confirmed int, recovered int, epochUpdate int64) Status {
	return Status{
		country:    strings.TrimSpace(country),
		deaths:     deaths,
		confirmed:  confirmed,
		recovered:  recovered,
		updateTime: time.Unix(epochUpdate, 0),
	}
}

func (s Status) Recovered() int {
	return s.recovered
}

func (s Status) Confirmed() int {
	return s.confirmed
}

func (s Status) Deaths() int {
	return s.deaths
}

func (s Status) UpdateTime() time.Time {
	return s.updateTime
}

func (s Status) Country() string {
	return s.country
}
