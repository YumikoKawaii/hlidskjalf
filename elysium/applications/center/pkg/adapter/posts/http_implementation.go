package posts

import (
	"bytes"
	"context"
	"encoding/json"
	"golang.org/x/xerrors"
	"net/http"
	"net/url"
)

const (
	upsertPostEndpoint = "/api/v1/posts"
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

func (c *httpClient) UpsertPost(ctx context.Context, request UpsertPostRequest) (UpsertPostResponse, error) {

	buffer, err := json.Marshal(request)
	if err != nil {
		return UpsertPostResponse{}, err
	}
	requestUrl, _ := url.JoinPath(c.host, upsertPostEndpoint)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestUrl, bytes.NewBuffer(buffer))
	if err != nil {
		return UpsertPostResponse{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return UpsertPostResponse{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return UpsertPostResponse{}, xerrors.Errorf("http status: %d", resp.StatusCode)
	}

	response := UpsertPostResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err

}

func (c *httpClient) GetPosts(ctx context.Context, request GetPostsRequest) (GetPostsResponse, error) {

	requestUrl, _ := url.JoinPath(c.host, upsertPostEndpoint, request.Query())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl, nil)
	if err != nil {
		return GetPostsResponse{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return GetPostsResponse{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return GetPostsResponse{}, xerrors.Errorf("http status: %d", resp.StatusCode)
	}

	response := GetPostsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
