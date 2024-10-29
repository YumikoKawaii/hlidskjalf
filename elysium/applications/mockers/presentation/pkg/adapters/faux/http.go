package faux

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type httpImpl struct {
	client *http.Client
	host   string
}

func NewHttpClient(host string) Client {
	return &httpImpl{
		client: &http.Client{},
		host:   host,
	}
}

func (c *httpImpl) Faux(ctx context.Context) (Response, error) {

	requestUrl := fmt.Sprintf("%s/api/v1/faux", c.host)
	req, _ := http.NewRequest(http.MethodGet, requestUrl, nil)
	respData, err := c.client.Do(req)
	if err != nil {
		return Response{}, err
	}

	defer respData.Body.Close()
	response := Response{}
	err = json.NewDecoder(respData.Body).Decode(&response)
	return Response{}, err
}
