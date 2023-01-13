package main

import (
	"fmt"
	"time"
)

func CheckTick(ticker1 *time.Ticker, ticker2 *time.Ticker) {

	select {
	case t := <-ticker1.C:
		fmt.Println("[11111]t comes from ticker1:", t)

	case t := <-ticker2.C:
		fmt.Println("[22222]t comes from ticker2:", t)

	}

}

func main() {
	fmt.Println("Ticker")

	ticker_1 := time.NewTicker(500 * time.Millisecond)
	ticker_2 := time.NewTicker(2000 * time.Millisecond)

	for {
		go CheckTick(ticker_1, ticker_2)
	}

}
