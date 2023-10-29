package handlers_test

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"os"
	"server/data_analysis/api/handlers"
	"server/data_analysis/models"
	"testing"
)

// Load mock data from a file and unmarshal it into a models.Stock object
func loadMockData(t *testing.T) models.Stock {
	// Read the mock data from the file
	body, err := os.ReadFile("../../testdata/stock_data.json")
	if err != nil {
		t.Fatal(err) // Fail the test if reading the file fails
	}

	var stockData models.Stock
	// Parse the JSON data into the Stock struct
	err = json.Unmarshal(body, &stockData)
	if err != nil {
		t.Fatal(err) // Fail the test if unmarshaling fails
	}

	return stockData
}

func TestFetchStockData(t *testing.T) {
	// Load mock data for comparison and server response
	mockData := loadMockData(t)

	// Create a mock HTTP server to simulate the external API
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Convert the mock data to JSON
		body, err := json.Marshal(mockData)
		if err != nil {
			t.Fatal(err) // Fail the test if marshaling fails
		}

		// Send the mock data as a response
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}))
	defer ts.Close() // Ensure the server is closed after the test

	// Call the function to be tested
	resp, err := handlers.FetchStockData(ts.URL)
	if err != nil {
		t.Fatal(err) // Fail if the function returns an error
	}
	defer resp.Body.Close() // Ensure response body is closed

	// Validate the HTTP status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code to be 200, but got %d", resp.StatusCode)
	}

	// Decode the response to a Stock object
	var stockData models.Stock
	err = json.NewDecoder(resp.Body).Decode(&stockData)
	if err != nil {
		t.Fatal(err) // Fail if decoding fails
	}

	// Compare the received data with the mock data
	if stockData.MetaData.Symbol != mockData.MetaData.Symbol {
		t.Errorf("Expected stock symbol to be %s, but got %s", mockData.MetaData.Symbol, stockData.MetaData.Symbol)
	}
}

func TestGetStockData(t *testing.T) {
	// Load mock data for comparison
	mockData := loadMockData(t)

	// Create a mock echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	// Set the mock stock symbol as a parameter in the echo context
	ctx.SetParamNames("symbol")
	ctx.SetParamValues(mockData.MetaData.Symbol)

	// Execute the handler function
	if err := handlers.GetStockData(ctx); err != nil {
		t.Fatal(err) // Fail if the handler returns an error
	}

	// Check the HTTP status code of the handler response
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code to be 200, but got %d", rec.Code)
	}

	// Decode the handler's JSON response
	var stockData models.Stock
	err := json.Unmarshal(rec.Body.Bytes(), &stockData)
	if err != nil {
		t.Fatal(err) // Fail if decoding fails
	}

	// Compare the handler's response with the mock data
	if stockData.MetaData.Symbol != mockData.MetaData.Symbol {
		t.Errorf("Expected stock symbol to be %s, but got %s", mockData.MetaData.Symbol, stockData.MetaData.Symbol)
	}
}
