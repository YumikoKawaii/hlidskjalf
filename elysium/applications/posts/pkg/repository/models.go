package repository

import "time"

type Post struct {
	Id        int32     `json:"id,omitempty"`
	Author    string    `json:"author,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type GetPostsParams struct {
	Ids      []int32
	Page     int
	PageSize int
}
