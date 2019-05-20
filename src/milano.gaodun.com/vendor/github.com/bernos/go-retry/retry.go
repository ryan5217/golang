// Package retry implements a simple mechanism for retrying functions, with support for
// specifying backoff strategies, timeout and retry limits
package retry

import (
	"fmt"
	"time"
)

const (
	DefaultMaxRetries = 10
	DefaultBaseDelay  = time.Millisecond
	DefaultMaxDelay   = time.Minute
	Infinity          = -1
)

var (
	DefaultBackoffFunc = BinaryExponentialBackoff()
)

// Options holds options for retrying
type Options struct {
	MaxRetries     int
	BaseDelay      time.Duration
	MaxDelay       time.Duration
	ShouldRetry    func(error) bool
	CalculateDelay BackoffFunc
	Log            func(string, ...interface{})
}

// Retry wraps a func, returning a new func that will retry the wrapped func until
// it succeeds, or until a max retry count or timeout is exceeded. options funcs
// can be used to configure properties of the wrapper func, such as max retry
// attempts, which backoff algrorithm to use and so forth.
func Retry(fn func() (interface{}, error), options ...func(*Options)) func() (interface{}, error) {
	r := Options{
		BaseDelay:   DefaultBaseDelay,
		MaxDelay:    DefaultMaxDelay,
		MaxRetries:  DefaultMaxRetries,
		ShouldRetry: func(err error) bool { return true },
		Log:         func(format string, v ...interface{}) {},
	}

	DefaultBackoffFunc(&r)

	for _, o := range options {
		o(&r)
	}

	return func() (interface{}, error) {
		var count int

		for {
			if count > 0 {
				r.Log("Retrying attempt %d", count)
			}

			value, err := fn()

			if err == nil {
				return value, err
			}

			if !r.ShouldRetry(err) {
				return nil, fmt.Errorf("Retrier aborted due to user supplied ShouldRetry func. Cause: %s", err.Error())
			}

			if count == r.MaxRetries {
				return nil, fmt.Errorf("Retrier exceeded max retry count of %d. Cause: %s", r.MaxRetries, err.Error())
			}

			d := r.CalculateDelay(uint(count), r.BaseDelay, r.MaxDelay)

			r.Log("Will retry in %s", d)

			time.Sleep(d)
			count++
		}
	}
}
