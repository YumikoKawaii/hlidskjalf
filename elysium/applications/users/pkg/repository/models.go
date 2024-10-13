package repository

import "time"

type User struct {
	Id           string    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Alias        string    `json:"alias,omitempty"`
	Avatar       string    `json:"avatar,omitempty"`
	Introduction string    `json:"introduction,omitempty"`
	Workplace    string    `json:"workplace,omitempty"`
	Hometown     string    `json:"hometown,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"cpdatedAt"`
}
