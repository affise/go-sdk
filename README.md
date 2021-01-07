# go-sdk
API v3.1 SDK (Golang)

The public API documentation is available at [api.affise.com](https://api.affise.com/docs3.1/)

## Example
```go
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/clobucks/go-sdk/affise"
)

func main() {
	// custom HTTP client
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	client, err := affise.NewClient(
		affise.WithAPIKey("key"),
		affise.WithBaseURL("https://base.example.com"),
		affise.WithAdminURL("https://admin.example.com"),
		affise.WithHTTPClient(httpClient),
	)
	if err != nil{
		log.Fatalf("creating client err: %v", err)
	}

	ctx := context.TODO()

	// Get statistic by date
	opts := &affise.StatisticGetByDateOpts{
		StatFilter: affise.StatFilter{
			DateFrom:            "2021-01-01",
			DateTo:              "2021-01-04",
		},
		Timezone: "Europe/Berlin",
	}

	stats, _ , err := client.Statistic.GetByDate(ctx, opts)
	if err != nil{
		log.Fatalf("get statistic by date err: %v", err)
	}

	// Get partner by ID (admin method)
	partner, _, err := client.AdminAffiliate.Get(ctx, 1)
	if err != nil{
		log.Fatalf("get partner by ID err: %v", err)
	}
}
```
