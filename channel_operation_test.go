package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type StockPrice struct {
	Symbol string
	Price  float64
	Time   time.Time
}

func TestChannelOperation(t *testing.T) {
	priceCh := make(chan StockPrice, 100)

	go PriceProcessor(priceCh)

	// start fetching stock prices for multiple symbols
	go fetchStockPrices(priceCh, "AAPL")
	go fetchStockPrices(priceCh, "GOOGL")
	go fetchStockPrices(priceCh, "MSFT")

	// simulate some delay to allow fetching and processing
	time.Sleep(10 * time.Second)
}

// this channel is to send only
//
//	<-chan untuk mengirim data
func PriceProcessor(priceCh <-chan StockPrice) {
	for price := range priceCh {
		fmt.Printf("Processing stock price: %s = %.2f at %s\n", price.Symbol, price.Price, price.Time)
		// Simulate processing time
		time.Sleep(500 * time.Millisecond)
	}
}

// fetchStockPrices will be callled by a scheduler
// this channel is to receive only
// chan <- untuk mengirim data
func fetchStockPrices(priceCh chan<- StockPrice, symbol string) {
	for {
		price := StockPrice{
			Symbol: symbol,
			Price:  rand.Float64() * 1000,
			Time:   time.Now(),
		}
		priceCh <- price
		// simulate delay between prices updates
		time.Sleep(time.Second)
	}
}
