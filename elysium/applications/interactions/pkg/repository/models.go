package repository

import (
	"elysium.com/applications/utils"
	"time"
)

type Interaction struct {
	Id        *uint32   `json:"id,omitempty"`
	PostId    uint32    `json:"postId,omitempty"`
	Author    string    `json:"author,omitempty"`
	Type      string    `json:"type,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAT time.Time `json:"updatedAT"`
}

type GetInteractionsParams struct {
	PostId   uint32
	Order    utils.SortOrder
	Page     uint32
	PageSize uint32
}
