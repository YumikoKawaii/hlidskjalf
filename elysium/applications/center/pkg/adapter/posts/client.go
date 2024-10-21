package posts

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Client interface {
	UpsertPost(ctx context.Context, request UpsertPostRequest) (UpsertPostResponse, error)
	GetPosts(ctx context.Context, request GetPostsRequest) (GetPostsResponse, error)
}

func NewClient(config Config, isUseGrpc bool) Client {
	if isUseGrpc {
		return NewRpcClient(config)
	}

	return NewHttpClient(config)
}

type Config struct {
	Host string `env:"POST_SERVICE_HOST"`
}

type UpsertPostRequest struct {
	Id      *uint32  `json:"id,omitempty"`
	Author  string   `json:"author,omitempty"`
	Content string   `json:"content,omitempty"`
	Medias  []string `json:"medias,omitempty"`
}

type UpsertPostResponse struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Id      int32  `json:"id,omitempty"`
}

type GetPostsRequest struct {
	Ids      []int32 `json:"ids,omitempty"`
	Page     int32   `json:"page,omitempty"`
	PageSize int32   `json:"pageSize,omitempty"`
}

func (r *GetPostsRequest) Query() string {
	query := url.Values{}
	if len(r.Ids) != 0 {
		ids := make([]string, 0)
		for _, id := range r.Ids {
			ids = append(ids, strconv.Itoa(int(id)))
		}
		query.Add("ids", strings.Join(ids, ","))
	}

	if r.Page != 0 {
		query.Add("page", strconv.Itoa(int(r.Page)))
	}

	if r.PageSize != 0 {
		query.Add("pageSize", strconv.Itoa(int(r.PageSize)))
	}

	return fmt.Sprintf("?%s", query.Encode())
}

type GetPostsResponse struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Posts    []Post `json:"posts,omitempty"`
		Page     int32  `json:"page,omitempty"`
		PageSize int32  `json:"pageSize,omitempty"`
	} `json:"data"`
}

type Post struct {
	Id        int32    `json:"id,omitempty"`
	Author    string   `json:"author,omitempty"`
	Content   string   `json:"content,omitempty"`
	Medias    []string `json:"medias,omitempty"`
	CreatedAt int32    `json:"createdAt,omitempty"`
	UpdatedAt int32    `json:"updatedAt,omitempty"`
}