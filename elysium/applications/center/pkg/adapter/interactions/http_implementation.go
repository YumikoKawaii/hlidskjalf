package interactions

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/xerrors"
	"net/http"
	"net/url"
)

type httpClient struct {
	client *http.Client
	host   string
}

const (
	upsertInteractionEndpoint = "/api/v1/interactions"
)

func NewHttpClient(config Config) Client {
	return &httpClient{
		host:   config.Host,
		client: &http.Client{},
	}
}

func (c *httpClient) UpsertInteraction(ctx context.Context, request UpsertInteractionRequest) (UpsertInteractionResponse, error) {

	buffer, err := json.Marshal(request)
	if err != nil {
		return UpsertInteractionResponse{}, err
	}
	requestUrl, _ := url.JoinPath(c.host, upsertInteractionEndpoint)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestUrl, bytes.NewBuffer(buffer))
	if err != nil {
		return UpsertInteractionResponse{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return UpsertInteractionResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UpsertInteractionResponse{}, xerrors.Errorf("http status: %d", resp.StatusCode)
	}

	response := UpsertInteractionResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func (c *httpClient) GetInteractions(ctx context.Context, request GetInteractionRequest) (GetInteractionResponse, error) {

	requestUrl := fmt.Sprintf("%s%s%s", c.host, upsertInteractionEndpoint, request.Query())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl, nil)
	if err != nil {
		return GetInteractionResponse{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return GetInteractionResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return GetInteractionResponse{}, xerrors.Errorf("http status: %d", resp.StatusCode)
	}

	response := GetInteractionResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
