package repository

import "time"

type Account struct {
	Id             string    `json:"id,omitempty"`
	HashedEmail    string    `json:"hashed_email,omitempty"`
	EncryptEmail   []byte    `json:"encrypt_email,omitempty"`
	HashedPassword string    `json:"hashed_password,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Permission struct {
	Id        int    `json:"id"`
	AccountId string `json:"account_id,omitempty"`
	Route     string `json:"route,omitempty"`
}
