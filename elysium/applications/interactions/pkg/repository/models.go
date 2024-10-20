package repository

import "time"

type Interaction struct {
	Id        *uint32   `json:"id,omitempty"`
	PostId    int32     `json:"postId,omitempty"`
	Author    string    `json:"author,omitempty"`
	Type      string    `json:"type,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAT time.Time `json:"updatedAT"`
}

type GetInteractionsParams struct {
	PostIds  []int32
	Page     int
	PageSize int
}
