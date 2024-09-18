package times

import (
	"github.com/Ccmuyu/gopkg/test"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {

	now := time.Now()

	t.Log(now)
	t.Log(now.Unix())
	// 2024-09-19 00:05:18
	// 1726675518
	unix := time.Unix(1726675518, 0)
	test.AssertEqual(t, Format(unix), "2024-09-19 00:05:18")
	test.AssertNotEqual(t, Format(unix), "=2024-09-19 00:05:18")

	test.AssertEqual(t, FormatF(unix, layoutUTC), "2024-09-19T00:05:18Z")
	test.AssertNotEqual(t, FormatF(unix, layoutUTC), "~2024-09-19T00:05:18Z")
}
