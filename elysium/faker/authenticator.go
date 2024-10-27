package main

import (
	"elysium.com/applications/utils"
	pb "elysium.com/pb/authenticator"
	"encoding/csv"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

type RegisterInformation struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Account struct {
	Id             string `json:"id,omitempty"`
	HashedEmail    string `json:"hashed_email,omitempty"`
	HashedPassword string `json:"hashed_password,omitempty"`
}

func (a *Account) ToSQLValue() string {
	return fmt.Sprintf("(\"%s\",\"%s\",\"%s\")", a.Id, a.HashedEmail, a.HashedPassword)
}

func (a *Account) ToCSV() string {
	return fmt.Sprintf("%s,%s,%s", a.Id, a.HashedEmail, a.HashedPassword)
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
			Id:       record[0],
			Email:    record[1],
			Password: record[2],
		})
	}

	hashKey := "123"

	accounts := make([]Account, 0)
	for _, registerInfo := range registerInfos {
		accounts = append(accounts, Account{
			Id:             registerInfo.Id,
			HashedEmail:    utils.HashString(registerInfo.Email, hashKey),
			HashedPassword: utils.HashString(registerInfo.Password, hashKey),
		})
	}

	csvFile, err := os.Create(fmt.Sprintf("%s%s", cwd, "/faker/migrations/accounts.csv"))
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	header := []string{"id", "hashed_email", "hashed_password"}
	if err := writer.Write(header); err != nil {
		fmt.Println("Error writing header:", err)
		return
	}

	for _, record := range accounts {
		row := []string{record.Id, record.HashedEmail, record.HashedPassword}
		if err := writer.Write(row); err != nil {
			fmt.Println("Error writing record:", err)
			return
		}
	}
	//stmts := make([]string, 0)
	//// divide by batches
	//batchSize := 1000
	//for {
	//	batch := accounts[:batchSize]
	//	accounts = accounts[batchSize:]
	//
	//	values := make([]string, 0)
	//	for _, account := range batch {
	//		values = append(values, account.ToSQLValue())
	//	}
	//
	//	stmt := fmt.Sprintf(""+
	//		"insert into `accounts`(id, hashed_email, encrypted_email, hashed_password)"+
	//		"values %s;",
	//		strings.Join(values, ","),
	//	)
	//
	//	stmts = append(stmts, stmt)
	//	if len(accounts) == 0 {
	//		break
	//	}
	//}
	//
	//migrationFile, err := os.Create(fmt.Sprintf("%s%s", cwd, "/faker/migrations/accounts.txt"))
	//if err != nil {
	//	panic(err)
	//}
	//for _, stmt := range stmts {
	//	_, err := migrationFile.WriteString(stmt + "\n")
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	fmt.Println("Done")
}
