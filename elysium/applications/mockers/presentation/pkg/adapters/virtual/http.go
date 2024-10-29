package virtual

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type httpImpl struct {
	client      *http.Client
	host        string
	skipMarshal bool
}

func NewHttpClient(host string, skipMarshal bool) Client {
	return &httpImpl{
		client:      &http.Client{},
		host:        host,
		skipMarshal: skipMarshal,
	}
}

func (c *httpImpl) Virtual(ctx context.Context) (Response, error) {

	start := time.Now()
	requestUrl := fmt.Sprintf("%s/api/v1/virtual", c.host)
	req, _ := http.NewRequest(http.MethodGet, requestUrl, nil)
	respData, err := c.client.Do(req)
	logrus.Infof("[Virtual] - Request: %d ms", time.Since(start).Milliseconds())
	if err != nil {
		return Response{}, err
	}

	defer respData.Body.Close()
	response := Response{}
	if !c.skipMarshal {
		err = json.NewDecoder(respData.Body).Decode(&response)
		logrus.Infof("[Echo] - Marshal: %d ms", time.Since(start).Milliseconds())
	}

	return response, err
}
