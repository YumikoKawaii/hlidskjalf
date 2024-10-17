package interactions

import (
	"bytes"
	"context"
	"encoding/json"
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
	defer resp.Body.Close()
	if err != nil {
		return UpsertInteractionResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return UpsertInteractionResponse{}, xerrors.Errorf("http status: %d", resp.StatusCode)
	}

	response := UpsertInteractionResponse{}
	bodyDecoder := json.NewDecoder(resp.Body)
	if err := bodyDecoder.Decode(&response); err != nil {
		return UpsertInteractionResponse{}, err
	}

	return response, nil
}

func (c *httpClient) GetInteractions(ctx context.Context, request GetInteractionRequest) (GetInteractionResponse, error) {
	return GetInteractionResponse{}, nil
}
