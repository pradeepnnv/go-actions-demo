package hello

import (
	quoteV3 "rsc.io/quote/v3"
	"fmt"
)

func Proverb() string {
	fmt.Println("hello sravs")
	return quoteV3.Concurrency()
}
