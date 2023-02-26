package hello

import (
	quoteV3 "rsc.io/quote/v3"
)

func Proverb() string {
	return quoteV3.Concurrency()
}
