package retry

import (
	"math/rand/v2"
	"time"
)

type Option func(*options)

type options struct {
	maxAttempts int
	delay       time.Duration
	backoff     bool
	jitter      bool
}

func WithMaxAttempts(n int) Option {
	return func(o *options) {
		o.maxAttempts = n
	}
}

func WithDelay(d time.Duration) Option {
	return func(o *options) {
		o.delay = d
	}
}

func WithBackoff() Option {
	return func(o *options) {
		o.backoff = true
	}
}

func WithJitter() Option {
	return func(o *options) {
		o.jitter = true
	}
}

func Retry(fn func() error, opts ...Option) error {
	o := &options{
		maxAttempts: 3,
		delay:       100 * time.Millisecond,
	}
	for _, opt := range opts {
		opt(o)
	}

	var err error
	delay := o.delay
	for i := 0; i < o.maxAttempts; i++ {
		if err = fn(); err == nil {
			return nil
		}
		if i < o.maxAttempts-1 {
			d := delay
			if o.jitter {
				d = time.Duration(float64(d) * (0.5 + rand.Float64()))
			}
			time.Sleep(d)
			if o.backoff {
				delay *= 2
			}
		}
	}
	return err
}
