package retry

import (
	"math"
	"time"
)

// BackoffFunc is a function that calculates the interval between retry attempts.
// The iteration argument indicates the retry number. baseDelay and maxDelay are
// the minimum and maximum times between retry attempts. A BackoffFunc should return
// a time.Duration that is a multiple of baseDelay, and less than maxDelay
type BackoffFunc func(iteration uint, baseDelay, maxDelay time.Duration) time.Duration

// BaseDelay sets the base delay time between retry attempts. By default the base
// delay is set to the value of DefaultBaseDelay
func BaseDelay(d time.Duration) func(*Options) {
	return func(r *Options) {
		r.BaseDelay = d
	}
}

// Forever configures the Retrier to retry forever
func Forever() func(*Options) {
	return func(r *Options) {
		r.MaxRetries = Infinity
	}
}

// MaxRetries sets the maximum number of times the Retrier will retry. By default the
// maximum number of retries is set to the value of DefaultMaxRetries
func MaxRetries(n int) func(*Options) {
	return func(r *Options) {
		r.MaxRetries = n
	}
}

// MaxDelay sets the maximum time between retry attempts. By default the maximum delay
// time is set to the value of DefaultMaxDelay
func MaxDelay(d time.Duration) func(*Options) {
	return func(r *Options) {
		r.MaxDelay = d
	}
}

// While configures the Retrier to only retry when fn returns true. The error
// passed to fn is the error that was returned by the function that is being
// retried. The function passed to While is checked only if the maximum retry
// count has not already been passed. By default the function will be retried
// if the value of error is not nil
func While(fn func(error) bool) func(*Options) {
	return func(r *Options) {
		r.ShouldRetry = fn
	}
}

// Backoff specifies the BackoffFunc used to calculate the interval between
// retry attempts. By default the BinaryExponentialBackoff function is used
func Backoff(fn BackoffFunc) func(*Options) {
	return func(r *Options) {
		r.CalculateDelay = fn
	}
}

// BinaryExponentialBackoff configures the Retrier to use the binary exponential
// backoff function when calculating the time between retry attempts. This is the
// default func used if no user defined backoff func is supplied
func BinaryExponentialBackoff() func(*Options) {
	return Backoff(calculateBinaryExponentialDelay)
}

// Log specifies a function to use to log retry attempts. Log messages will
// be sent when the a retry is performed, and when the retry interval is
// calculated
func Log(fn func(format string, v ...interface{})) func(*Options) {
	return func(r *Options) {
		r.Log = fn
	}
}

// An int64 binary exponential implementation to save from having to convert
// to and from float64 to use the stdlib math.Pow func
func binaryRaise(exponent uint) int64 {
	// Clamp exponent to avoid 64 bit overflow
	if exponent > 62 {
		exponent = 62
	}

	return 1 << exponent
}

func calculateBinaryExponentialDelay(iteration uint, baseDelay, maxDelay time.Duration) time.Duration {
	m := (binaryRaise(iteration) - 1) >> 1

	// If multiplier is greater than maxDelay, then we we don't need to
	// calculate
	if m > int64(maxDelay) {
		return maxDelay
	}

	d := (time.Duration(m) * baseDelay) + baseDelay

	if d < maxDelay && d > 0 {
		return d
	}

	return maxDelay
}

func calculateDelay(iteration uint, baseDelay, maxDelay time.Duration) time.Duration {
	m := ((math.Pow(2, float64(iteration))) - 1) / 2

	// Check for wrapping, or multiplier greater than max duration in
	// nanosecs
	if int64(m) < 0 || int64(m) > int64(maxDelay) {
		return maxDelay
	}

	d := (time.Duration(m) * baseDelay) + baseDelay

	if d < maxDelay && d > 0 {
		return d
	}

	return maxDelay
}
