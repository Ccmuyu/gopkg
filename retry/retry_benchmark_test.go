package retry

import (
	"errors"
	"testing"
	"time"
)

func BenchmarkRetrySuccess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Retry(func() error {
			return nil
		}, WithMaxAttempts(1))
	}
}

func BenchmarkRetryFailure(b *testing.B) {
	err := errors.New("fail")
	for i := 0; i < b.N; i++ {
		Retry(func() error {
			return err
		}, WithMaxAttempts(3), WithDelay(time.Microsecond))
	}
}
