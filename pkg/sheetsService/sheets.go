package sheetsService

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SheetService struct {
	service *sheets.Service
}

type Data struct{}

func getClient(config *oauth2.Config) *http.Client {
	tokenFile := "token.json"
	tok, err := tokenFromFile(tokenFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokenFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func NewSheetService(ctx context.Context) *SheetService {
	// Load client secret file

	// b, err := os.ReadFile("client_secret.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Get Google Sheets service
	// config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client := getClient(config)
	service, err := sheets.NewService(ctx, option.WithCredentialsFile("client_secret.json"))
	if err != nil {
		log.Fatal(err)
	}
	return &SheetService{service: service}
}

func (s *SheetService) ReadFromSheet(spreadsheetId, readRange string) (*sheets.ValueRange, error) {
	// Implement the function to read data from the Google Sheet
	return nil, nil
}

func (s *SheetService) UpdateSheet(spreadsheetId, writeRange string, data *Data) error {
	// Implement the function to update the Google Sheet with the status
	return nil
}
