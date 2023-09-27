package httpclient

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const (
	contentTypeHeaderKey = "Content-Type"
	ContentTypeJSON      = "application/json"

	acceptHeaderKey          = "Accept"
	acceptEncodingHeaderKey  = "Accept-Encoding"
	contentEncodingHeaderKey = "Content-Encoding"
	gzipEncodingHeaderValue  = "gzip"
)

type HttpHandler interface {
	Do(method, path, contentType string, send []byte, acceptHeaderValues ...string) (*http.Response, error)
	ReadJson(response *http.Response, valuePtr interface{}) error
}

type httpHandlerImpl struct {
	baseURL    string
	httpClient http.Client
}

func NewHttpHandler(baseURL string) HttpHandler {
	return &httpHandlerImpl{
		baseURL: baseURL,
	}
}

func (h *httpHandlerImpl) Do(method, path, contentType string, send []byte, acceptHeaderValues ...string) (*http.Response, error) {

	if path[0] == '/' {
		path = path[1:]
	}

	uri := h.baseURL + "/" + path

	req, err := http.NewRequest(method, uri, acquireBuffer(send))
	if err != nil {
		return nil, err
	}

	if contentType != "" {
		req.Header.Set(contentTypeHeaderKey, contentType)
	}

	req.Header.Add(acceptEncodingHeaderKey, gzipEncodingHeaderValue)
	if len(acceptHeaderValues) > 0 {
		req.Header.Add(acceptHeaderKey, strings.Join(acceptHeaderValues, ","))
	}

	response, err := h.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if isRespOk(response) {
		return response, nil
	}

	defer response.Body.Close()
	return nil, err
}

func (h *httpHandlerImpl) ReadJson(response *http.Response, valuePtr interface{}) error {
	body, err := h.readResponseBody(response)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, valuePtr)
}

func (h *httpHandlerImpl) readResponseBody(response *http.Response) ([]byte, error) {
	reader, err := acquireResponseBodyStream(response)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	if err = reader.Close(); err != nil {
		return nil, err
	}
	return body, err
}

type gzipReadCloser struct {
	responseReader io.ReadCloser
	gzipReader     io.ReadCloser
}

func (rc *gzipReadCloser) Close() error {
	if rc.gzipReader != nil {
		defer rc.gzipReader.Close()
	}

	return rc.responseReader.Close()
}

func (rc *gzipReadCloser) Read(p []byte) (n int, err error) {
	if rc.gzipReader != nil {
		return rc.gzipReader.Read(p)
	}

	return rc.responseReader.Read(p)
}

func acquireBuffer(b []byte) *bytes.Buffer {
	if len(b) > 0 {
		return bytes.NewBuffer(b)
	}
	return new(bytes.Buffer)
}

func isRespOk(response *http.Response) bool {
	return !(response.StatusCode < 200 || response.StatusCode >= 300)
}

func acquireResponseBodyStream(response *http.Response) (reader io.ReadCloser, err error) {
	reader = response.Body

	if encoding := response.Header.Get(contentEncodingHeaderKey); encoding == gzipEncodingHeaderValue {
		reader, err = gzip.NewReader(reader)
		if err != nil {
			return nil, err
		}
		return &gzipReadCloser{
			responseReader: response.Body,
			gzipReader:     reader,
		}, nil
	}

	return reader, err
}
