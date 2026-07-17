package retry

import (
	"context"
	"errors"
	"testing"
	"time"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestRetrySuccess(t *testing.T) {
	attempts := 0
	err := Retry(func() error {
		attempts++
		return nil
	}, WithMaxAttempts(3))
	AssertNoError(t, err)
	AssertEqual(t, attempts, 1)
}

func TestRetrySuccessAfterFailure(t *testing.T) {
	attempts := 0
	err := Retry(func() error {
		attempts++
		if attempts < 3 {
			return errors.New("not yet")
		}
		return nil
	}, WithMaxAttempts(5), WithDelay(time.Millisecond))
	AssertNoError(t, err)
	AssertEqual(t, attempts, 3)
}

func TestRetryExhausted(t *testing.T) {
	attempts := 0
	err := Retry(func() error {
		attempts++
		return errors.New("always fails")
	}, WithMaxAttempts(3), WithDelay(time.Millisecond))
	AssertError(t, err)
	AssertEqual(t, attempts, 3)
}

func TestRetryMaxAttempts(t *testing.T) {
	custom := 5
	attempts := 0
	err := Retry(func() error {
		attempts++
		return errors.New("fail")
	}, WithMaxAttempts(custom), WithDelay(time.Millisecond))
	AssertError(t, err)
	AssertEqual(t, attempts, custom)
}

func TestRetryBackoff(t *testing.T) {
	attempts := 0
	err := Retry(func() error {
		attempts++
		return errors.New("fail")
	}, WithMaxAttempts(3), WithDelay(time.Millisecond), WithBackoff())
	AssertError(t, err)
	AssertEqual(t, attempts, 3)
}

func TestRetryJitter(t *testing.T) {
	attempts := 0
	err := Retry(func() error {
		attempts++
		return errors.New("fail")
	}, WithMaxAttempts(3), WithDelay(time.Millisecond), WithJitter())
	AssertError(t, err)
	AssertEqual(t, attempts, 3)
}

func TestRetryDefaultValues(t *testing.T) {
	attempts := 0
	err := Retry(func() error {
		attempts++
		return errors.New("fail")
	}, WithDelay(time.Millisecond))
	AssertError(t, err)
	AssertEqual(t, attempts, 3)
}

func TestRetryNonPositiveAttemptsRunsOnce(t *testing.T) {
	for _, attempts := range []int{0, -1} {
		calls := 0
		err := Retry(func() error {
			calls++
			return errors.New("fail")
		}, WithMaxAttempts(attempts))
		AssertError(t, err)
		AssertEqual(t, calls, 1)
	}
}

func TestRetryCtxDoesNotRunAfterCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	calls := 0

	err := RetryCtx(ctx, func() error {
		calls++
		return nil
	})
	AssertTrue(t, errors.Is(err, context.Canceled))
	AssertEqual(t, calls, 0)
}
