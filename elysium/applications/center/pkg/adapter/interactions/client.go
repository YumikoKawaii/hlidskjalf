package interactions

import (
	"context"
	"elysium.com/applications/utils"
	"fmt"
	"net/url"
	"strconv"
)

type Client interface {
	UpsertInteraction(ctx context.Context, request UpsertInteractionRequest) (UpsertInteractionResponse, error)
	GetInteractions(ctx context.Context, request GetInteractionRequest) (GetInteractionResponse, error)
}

func NewClient(config Config, isUseGrpc bool) Client {
	if isUseGrpc {
		return NewRpcClient(config)
	}

	return NewHttpClient(config)
}

type Config struct {
	Host string `env:"INTERACTIONS_SERVICE_HOST"`
}

type UpsertInteractionRequest struct {
	Id      *uint32 `json:"id,omitempty"`
	PostId  uint32  `json:"postId,omitempty"`
	Author  string  `json:"author,omitempty"`
	Type    string  `json:"type,omitempty"`
	Content string  `json:"content,omitempty"`
}

type UpsertInteractionResponse struct {
	Code          uint32 `json:"code,omitempty"`
	Message       string `json:"message,omitempty"`
	InteractionId uint32 `json:"interactionId,omitempty"`
}

type GetInteractionRequest struct {
	PostId    uint32          `json:"postId,omitempty"`
	SortOrder utils.SortOrder `json:"order,omitempty"`
	Page      uint32          `json:"page,omitempty"`
	PageSize  uint32          `json:"pageSize,omitempty"`
}

func (r *GetInteractionRequest) Query() string {
	query := url.Values{}
	query.Add("postId", strconv.Itoa(int(r.PostId)))

	if r.Page != 0 {
		query.Add("page", strconv.Itoa(int(r.Page)))
	}

	if r.PageSize != 0 {
		query.Add("pageSize", strconv.Itoa(int(r.PageSize)))
	}

	query.Add("sortOrder", r.SortOrder.Value())
	return fmt.Sprintf("?%s", query.Encode())
}

type GetInteractionResponse struct {
	Code    uint32 `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Interactions []Interaction `json:"interactions,omitempty"`
		Page         uint32        `json:"page,omitempty"`
		PageSize     uint32        `json:"pageSize,omitempty"`
	} `json:"data"`
}

type Interaction struct {
	Id        uint32 `json:"id,omitempty"`
	PostId    uint32 `json:"postId,omitempty"`
	Author    string `json:"author,omitempty"`
	Type      string `json:"type,omitempty"`
	Content   string `json:"content,omitempty"`
	CreatedAt uint32 `json:"createdAt,omitempty"`
	UpdatedAt uint32 `json:"updatedAt,omitempty"`
}
