package users

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Client interface {
	UpsertUser(ctx context.Context, request UpsertUserRequest) (UpsertUserResponse, error)
	GetUsers(ctx context.Context, request GetUsersRequest) (GetUsersResponse, error)
}

func NewClient(config Config, isUseHttp bool) Client {
	if isUseHttp {
		return NewHttpClient(config)
	}

	return NewRpcClient(config)
}

type Config struct {
	Host string `env:"USER_SERVICE_HOST"`
}

type UserInfo struct {
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Alias        string `json:"alias,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Workplace    string `json:"workplace,omitempty"`
	Hometown     string `json:"hometown,omitempty"`
}

type UpsertUserRequest struct {
	Name         string `json:"name,omitempty"`
	Alias        string `json:"alias,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Workplace    string `json:"workplace,omitempty"`
	Hometown     string `json:"hometown,omitempty"`
}

type UpsertUserResponse struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Id      string `json:"id,omitempty"`
}

type GetUsersRequest struct {
	Ids      []string `json:"ids,omitempty"`
	Page     int32    `json:"page,omitempty"`
	PageSize int32    `json:"pageSize,omitempty"`
}

func (r *GetUsersRequest) Query() string {
	query := url.Values{}
	if len(r.Ids) != 0 {
		query.Add("ids", strings.Join(r.Ids, ","))
	}

	if r.Page != 0 {
		query.Add("page", strconv.Itoa(int(r.Page)))
	}

	if r.PageSize != 0 {
		query.Add("pageSize", strconv.Itoa(int(r.PageSize)))
	}

	return fmt.Sprintf("?%s", query.Encode())
}

type GetUsersResponse struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		UsersInfo []UserInfo `json:"usersInfo,omitempty"`
		Page      int32      `json:"page,omitempty"`
		PageSize  int32      `json:"pageSize,omitempty"`
	} `json:"data"`
}
