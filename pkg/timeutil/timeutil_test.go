package timeutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsTimeRangeOverlap(t *testing.T) {
	tests := []struct {
		name        string
		firstRange  string
		secondRange string
		want        bool
		wantErr     error
	}{
		{
			name:        "overlap in middle",
			firstRange:  "09:00-12:00",
			secondRange: "11:00-14:00",
			want:        true,
			wantErr:     nil,
		},
		{
			name:        "exact same time range",
			firstRange:  "09:00-12:00",
			secondRange: "09:00-12:00",
			want:        true,
			wantErr:     nil,
		},
		{
			name:        "no overlap - gap between ranges",
			firstRange:  "09:00-11:00",
			secondRange: "12:00-14:00",
			want:        false,
			wantErr:     nil,
		},
		{
			name:        "touching times - considered as overlap",
			firstRange:  "09:00-11:00",
			secondRange: "11:00-13:00",
			want:        true,
			wantErr:     nil,
		},
		{
			name:        "first range contains second",
			firstRange:  "09:00-14:00",
			secondRange: "10:00-13:00",
			want:        true,
			wantErr:     nil,
		},
		{
			name:        "second range contains first",
			firstRange:  "10:00-13:00",
			secondRange: "09:00-14:00",
			want:        true,
			wantErr:     nil,
		},
		{
			name:        "missing hyphen in first range",
			firstRange:  "09:0012:00",
			secondRange: "11:00-13:00",
			want:        false,
			wantErr:     ErrInvalidTimeRangeFormat,
		},
		{
			name:        "missing hyphen in second range",
			firstRange:  "09:00-12:00",
			secondRange: "11:0013:00",
			want:        false,
			wantErr:     ErrInvalidTimeRangeFormat,
		},
		{
			name:        "invalid time segments in first range",
			firstRange:  "09:00-12:00-15:00",
			secondRange: "11:00-13:00",
			want:        false,
			wantErr:     ErrInvalidTimeRangeSegments,
		},
		{
			name:        "invalid time segments in second range",
			firstRange:  "09:00-12:00",
			secondRange: "11:00",
			want:        false,
			wantErr:     ErrInvalidTimeRangeFormat,
		},
		{
			name:        "invalid start time format in first range",
			firstRange:  "25:00-12:00",
			secondRange: "11:00-13:00",
			want:        false,
			wantErr:     ErrInvalidStartTime,
		},
		{
			name:        "invalid end time format in first range",
			firstRange:  "09:00-25:00",
			secondRange: "11:00-13:00",
			want:        false,
			wantErr:     ErrInvalidEndTime,
		},
		{
			name:        "invalid start time format in second range",
			firstRange:  "09:00-12:00",
			secondRange: "25:00-13:00",
			want:        false,
			wantErr:     ErrInvalidStartTime,
		},
		{
			name:        "invalid end time format in second range",
			firstRange:  "09:00-12:00",
			secondRange: "11:00-25:00",
			want:        false,
			wantErr:     ErrInvalidEndTime,
		},
		{
			name:        "end before start in first range",
			firstRange:  "12:00-09:00",
			secondRange: "10:00-11:00",
			want:        false,
			wantErr:     ErrEndTimeBeforeStart,
		},
		{
			name:        "end before start in second range",
			firstRange:  "09:00-12:00",
			secondRange: "13:00-11:00",
			want:        false,
			wantErr:     ErrEndTimeBeforeStart,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsTimeRangeOverlap(tt.firstRange, tt.secondRange)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.False(t, got)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
