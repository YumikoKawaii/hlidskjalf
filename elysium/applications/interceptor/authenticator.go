package interceptor

import (
	"bytes"
	"context"
	pb "elysium.com/pb/authenticator"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"net/url"
)

const (
	getPermissionsEndpoint = "/api/v1/permissions/"
	verifyEndpoint         = "/api/v1/verify"
	authorizationHeader    = "authorization"
)

type Authenticator interface {
	GetPermissions(ctx context.Context, id string) ([]string, error)
	Verify(ctx context.Context, token, route string) (string, bool, error)
}

type grpcImpl struct {
	client pb.AuthenticatorClient
}

func NewGRPCAuthenticator(host string) Authenticator {

	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := pb.NewAuthenticatorClient(conn)
	return &grpcImpl{
		client: client,
	}
}

func (i *grpcImpl) GetPermissions(ctx context.Context, id string) ([]string, error) {

	resp, err := i.client.GetPermissions(ctx, &pb.GetPermissionsRequest{
		UserId: id,
	})

	if err != nil {
		return nil, err
	}

	return resp.Data.Permissions, nil
}

type httpImpl struct {
	client *http.Client
	host   string
}

func (i *grpcImpl) Verify(ctx context.Context, token, route string) (string, bool, error) {

	resp, err := i.client.Verify(ctx, &pb.VerifyRequest{
		Token: token,
		Route: route,
	})
	if err != nil {
		return "", false, err
	}

	if resp.Code != http.StatusOK {
		return resp.Id, false, nil
	}

	return resp.Id, true, nil
}

func NewHttpAuthenticator(host string) Authenticator {
	return &httpImpl{
		client: &http.Client{},
		host:   host,
	}
}

func (i *httpImpl) GetPermissions(ctx context.Context, id string) ([]string, error) {

	requestURL, _ := url.JoinPath(i.host, getPermissionsEndpoint, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	respData, err := i.client.Do(req)
	defer respData.Body.Close()
	if err != nil {
		return nil, err
	}

	type Resp struct {
		Code    int32  `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
		Data    struct {
			Permissions []string `json:"permissions,omitempty"`
		} `json:"data"`
	}

	respDecoder := json.NewDecoder(respData.Body)
	resp := new(Resp)
	if err := respDecoder.Decode(resp); err != nil {
		return nil, err
	}

	if resp.Code != http.StatusOK {
		logrus.Errorf("error authenticator: %s", resp.Message)
	}

	return resp.Data.Permissions, nil
}

func (i *httpImpl) Verify(ctx context.Context, token, route string) (string, bool, error) {

	type Request struct {
		Token string `json:"token"`
		Route string `json:"route,omitempty"`
	}

	buffer, _ := json.Marshal(&Request{Token: token, Route: route})

	requestURL, _ := url.JoinPath(i.host, verifyEndpoint)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(buffer))
	respData, err := i.client.Do(req)
	if err != nil {
		return "", false, err
	}
	defer respData.Body.Close()

	type Response struct {
		Code int32  `json:"code,omitempty"`
		Id   string `json:"id,omitempty"`
	}
	respDecoder := json.NewDecoder(respData.Body)
	resp := new(Response)
	if err := respDecoder.Decode(resp); err != nil {
		return "", false, err
	}

	return resp.Id, resp.Code == http.StatusOK, nil
}
