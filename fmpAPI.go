package main

import (
    "encoding/json"
    "fmt"
    "os"
    "github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

func loadEnv(){
	err := godotenv.Load("mystuff.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

type FMPCandle struct {
    Date   string  `json:"date"`
    Open   float64 `json:"open"`
    High   float64 `json:"high"`
    Low    float64 `json:"low"`
    Close  float64 `json:"close"`
    Volume float64 `json:"volume"`
}

func callFMPCandles(symbol string) ([]FMPCandle, error) {
    loadEnv()
    
    client := resty.New()

    resp, err := client.R().
        SetQueryParams(map[string]string{
            "apikey": os.Getenv("MY_FMP_API_KEY"),
        }).
        Get(fmt.Sprintf("https://financialmodelingprep.com/api/v3/historical-price-full/%s", symbol))

    if err != nil {
        return nil, fmt.Errorf("request failed: %w", err)
    }

    var result struct {
        Historical []FMPCandle `json:"historical"`
    }

    if err := json.Unmarshal(resp.Body(), &result); err != nil {
        return nil, fmt.Errorf("failed to parse response: %w", err)
    }

    return result.Historical, nil
}
