package models

import (
	"errors"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Sight struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
	ContentInfo  string `json:"content_info"`
	BusNumbers   string `json:"bus_numbers"`
	WorkingDays  string `json:"working_days"`
	WorkingHours string `json:"working_hours"`
	Visited      int    `json:"visited"`
	ImageUrl     string `json:"image_url"`
}
