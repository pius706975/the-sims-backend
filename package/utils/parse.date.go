package utils

import "time"

func ParseDate(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, nil
	}
	layout := "2006-01-02"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
