# go-retry

[![Build Status](https://travis-ci.org/bernos/go-retry.svg)](https://travis-ci.org/bernos/go-retry)&nbsp;[![GoDoc](https://godoc.org/github.com/bernos/go-retry?status.svg)](https://godoc.org/github.com/bernos/go-retry)

`go get github.com/bernos/go-retry`

Simple pkg for creating retryable funcs in go. See simple example below. For more detail, check the [go docs](https://godoc.org/github.com/bernos/go-retry)

```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bernos/go-retry"
)

func main() {
	r := retry.Retry(
		numberOrBust(128, 255),
		retry.MaxRetries(5),
		retry.BaseDelay(time.Millisecond))

	value, err := r()

	if err != nil {
		fmt.Printf("Bad luck! %s\n", err.Error())
	} else {
		fmt.Printf("Jackpot! You rolled a %d\n", value)
	}
}

// numberOrBust creates a func that chooses a random number between 0 and maxNumber
// and returns an error if that random number does not match the value of magicNumber
func numberOrBust(magicNumber int, maxNumber int) func() (interface{}, error) {
	return func() (interface{}, error) {
		guess := rand.Intn(maxNumber)

		if guess == magicNumber {
			return "Got it!", nil
		}

		return nil, fmt.Errorf("Want %d, got %d", magicNumber, guess)
	}
}
```
