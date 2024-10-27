package main

import (
	"context"
	pb "elysium.com/pb/authenticator"
	"encoding/csv"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

type RegisterInformation struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func initAuthenticator() pb.AuthenticatorClient {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return pb.NewAuthenticatorClient(conn)
}

func main() {
	cwd, _ := os.Getwd()
	file, err := os.Open(fmt.Sprintf("%s%s", cwd, "/etl/faker/data/accounts.csv"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	registerInfos := make([]RegisterInformation, 0)
	for idx, record := range records {
		if idx == 0 {
			continue
		}
		registerInfos = append(registerInfos, RegisterInformation{
			Email:    record[0],
			Password: record[1],
		})
	}

	ctx := context.Background()
	authenticator := initAuthenticator()
	for _, registerInfo := range registerInfos {
		_, err := authenticator.Signup(ctx, &pb.SignupRequest{
			Email:    registerInfo.Email,
			Password: registerInfo.Password,
		})
		if err != nil {
			logrus.Error(err.Error())
		}
	}

}
