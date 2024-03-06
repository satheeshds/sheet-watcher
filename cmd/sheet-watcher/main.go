package main

import (
	"context"
	"fmt"

	"google.golang.org/api/sheets/v4"

	"sheet-watcher/pkg/sheetsService"
)

var spreadsheetId string = "your-spreadsheet-id"

func processData(data *sheets.ValueRange) (*sheetsService.Data, error) {
	// Implement the function to process the data
	return nil, nil
}

func main() {

	sheetService := sheetsService.NewSheetService(context.TODO())
	// Read from Google Sheet
	readRange := "Sheet1!A1:E5"
	resp, err := sheetService.ReadFromSheet(spreadsheetId, readRange)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Process the data
	data, err := processData(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Update the Google Sheet with the status
	writeRange := "Sheet1!F1:F5"
	err = sheetService.UpdateSheet(spreadsheetId, writeRange, data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
