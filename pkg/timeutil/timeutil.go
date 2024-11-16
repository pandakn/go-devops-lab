package timeutil

import (
	"strings"
	"time"
)

// IsTimeRangeOverlap checks if two time ranges overlap. Time ranges must be in 24-hour format.
//
// Parameters:
//   - timeRange1: First time range in format "HH:MM-HH:MM" (e.g., "09:00-12:00")
//   - timeRange2: Second time range in format "HH:MM-HH:MM" (e.g., "11:00-14:00")
//
// Returns:
//   - bool: true if ranges overlap, false otherwise
//   - error: nil if valid input, error otherwise
//
// Examples:
//
//	IsTimeRangeOverlap("09:00-12:00", "11:00-14:00") // returns true, nil
//	IsTimeRangeOverlap("09:00-11:00", "12:00-14:00") // returns false, nil
//	IsTimeRangeOverlap("09:00-08:00", "12:00-14:00") // returns false, ErrEndBeforeStart
//	IsTimeRangeOverlap("25:00-26:00", "12:00-14:00") // returns false, ErrInvalidFirstStart
func IsTimeRangeOverlap(firstRange, secondRange string) (bool, error) {
	if !strings.Contains(firstRange, "-") || !strings.Contains(secondRange, "-") {
		return false, ErrInvalidTimeRangeFormat
	}

	// Parse first time range
	firstRangeParts := strings.Split(firstRange, "-")
	if len(firstRangeParts) != 2 {
		return false, ErrInvalidTimeRangeSegments
	}
	firstStart, err := time.Parse("15:04", firstRangeParts[0])
	if err != nil {
		return false, ErrInvalidStartTime
	}
	firstEnd, err := time.Parse("15:04", firstRangeParts[1])
	if err != nil {
		return false, ErrInvalidEndTime
	}

	// Parse second time range
	secondRangeParts := strings.Split(secondRange, "-")
	if len(secondRangeParts) != 2 {
		return false, ErrInvalidTimeRangeSegments
	}
	secondStart, err := time.Parse("15:04", secondRangeParts[0])
	if err != nil {
		return false, ErrInvalidStartTime
	}
	secondEnd, err := time.Parse("15:04", secondRangeParts[1])
	if err != nil {
		return false, ErrInvalidEndTime
	}

	// Check if end times are after start times
	if firstEnd.Before(firstStart) || secondEnd.Before(secondStart) {
		return false, ErrEndTimeBeforeStart
	}

	return !firstStart.After(secondEnd) && !secondStart.After(firstEnd), nil
}
