package acoustics

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	EntryEndpoint = "api/v1/entry"
)

type httpClient struct {
	endpoint string
	client   *http.Client
}

func NewHTTPClient(endpoint string) (Client, error) {
	return &httpClient{
		endpoint: endpoint,
		client:   &http.Client{},
	}, nil
}

func (c *httpClient) Entry(ctx context.Context, request EntryRequest) (EntryResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return EntryResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/%s", c.endpoint, EntryEndpoint), bytes.NewReader(body))
	if err != nil {
		return EntryResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return EntryResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return EntryResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response EntryResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return EntryResponse{}, err
	}

	return response, nil
}
