package hello

import (
	quote "rsc.io/quote/v3"
)

func Hello() string {
	return quote.HelloV3()
}
