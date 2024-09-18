package times

import "time"

const (
	layoutDefault = "2006-01-02 15:04:05"
	layoutUTC     = "2006-01-02T15:04:05Z"
)

func Format(t time.Time) string {
	return t.Format(layoutDefault)
}

func FormatF(t time.Time, layout string) string {
	return t.Format(layout)
}

func ParseUTCTime(t string) time.Time {
	parse, err := time.Parse(layoutUTC, t)
	if err != nil {
		return time.Time{}
	}
	return parse
}
func ParseUTCTimePtr(t *string) time.Time {
	if t == nil {
		return time.Time{}
	}
	return ParseUTCTime(*t)
}
