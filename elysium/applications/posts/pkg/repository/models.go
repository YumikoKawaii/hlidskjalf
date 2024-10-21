package repository

import (
	"elysium.com/applications/utils"
	"time"
)

type Post struct {
	Id        *uint32   `json:"id,omitempty"`
	Author    string    `json:"author,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type GetPostsParams struct {
	Ids      []uint32
	Author   string
	Page     uint32
	PageSize uint32
	Order    utils.SortOrder
}
