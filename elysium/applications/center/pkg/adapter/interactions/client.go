package interactions

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
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
	PostIds  []uint32 `json:"postIds,omitempty"`
	Page     uint32   `json:"page,omitempty"`
	PageSize uint32   `json:"pageSize,omitempty"`
}

func (r *GetInteractionRequest) Query() string {
	query := url.Values{}
	if len(r.PostIds) != 0 {
		ids := make([]string, 0)
		for _, id := range r.PostIds {
			ids = append(ids, strconv.Itoa(int(id)))
		}
		query.Add("postIds", strings.Join(ids, ","))
	}

	if r.Page != 0 {
		query.Add("page", strconv.Itoa(int(r.Page)))
	}

	if r.PageSize != 0 {
		query.Add("pageSize", strconv.Itoa(int(r.PageSize)))
	}

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
	CreatedAt int32  `json:"createdAt,omitempty"`
	UpdatedAt int32  `json:"updatedAt,omitempty"`
}
