package timeutil

import "errors"

var (
	ErrInvalidTimeRangeFormat   = errors.New("invalid time range format: expected HH:MM-HH:MM")
	ErrEndTimeBeforeStart       = errors.New("end time cannot be before start time")
	ErrInvalidTimeRangeSegments = errors.New("invalid time range format: must have exactly two time segments")
	ErrInvalidStartTime         = errors.New("invalid start time format")
	ErrInvalidEndTime           = errors.New("invalid end time format")
)
