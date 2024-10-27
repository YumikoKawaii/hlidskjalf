package users

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/xerrors"
	"net/http"
	"net/url"
)

const (
	upsertUserInfoEndpoint = "/api/v1/users"
)

type httpClient struct {
	host   string
	client *http.Client
}

func NewHttpClient(config Config) Client {
	return &httpClient{
		host:   config.Host,
		client: &http.Client{},
	}
}

func (c *httpClient) UpsertUser(ctx context.Context, request UpsertUserRequest) (UpsertUserResponse, error) {
	buffer, err := json.Marshal(request)
	if err != nil {
		return UpsertUserResponse{}, err
	}
	requestUrl, _ := url.JoinPath(c.host, upsertUserInfoEndpoint)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestUrl, bytes.NewBuffer(buffer))
	if err != nil {
		return UpsertUserResponse{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return UpsertUserResponse{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return UpsertUserResponse{}, xerrors.Errorf("http status: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	response := UpsertUserResponse{}
	err = decoder.Decode(&resp)
	return response, err
}

func (c *httpClient) GetUsers(ctx context.Context, request GetUsersRequest) (GetUsersResponse, error) {

	requestUrl := fmt.Sprintf("%s%s%s", c.host, upsertUserInfoEndpoint, request.Query())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl, nil)
	if err != nil {
		return GetUsersResponse{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return GetUsersResponse{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return GetUsersResponse{}, xerrors.Errorf("http status: %d", resp.StatusCode)
	}

	response := GetUsersResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
