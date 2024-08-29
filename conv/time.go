package conv

import "time"

func ParseUTCTime(t string) time.Time {
	parse, err := time.Parse("2006-01-02T15:04:05Z", t)
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
