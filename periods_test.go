package periods_test

import (
	"testing"

	periods "github.com/kotaoue/go-periods"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		name    string
		period  string
		want    []string
		wantErr bool
	}{
		{
			name:   "single day",
			period: "20220101",
			want:   []string{"20220101"},
		},
		{
			name:   "day range",
			period: "20220101-20220105",
			want:   []string{"20220101", "20220102", "20220103", "20220104", "20220105"},
		},
		{
			name:   "single month",
			period: "202201",
			want:   []string{"202201"},
		},
		{
			name:   "month range",
			period: "202201-202203",
			want:   []string{"202201", "202202", "202203"},
		},
		{
			name:    "three periods",
			period:  "20220101-20220102-20220103",
			wantErr: true,
		},
		{
			name:    "unexpected format",
			period:  "2022-2023",
			wantErr: true,
		},
		{
			name:    "invalid day from",
			period:  "20221301-20221305",
			wantErr: true,
		},
		{
			name:    "invalid day to",
			period:  "20220101-20221305",
			wantErr: true,
		},
		{
			name:    "invalid month from",
			period:  "202213-202215",
			wantErr: true,
		},
		{
			name:    "invalid month to",
			period:  "202201-202213",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := periods.Split(tt.period)
			if (err != nil) != tt.wantErr {
				t.Errorf("Split(%q) error = %v, wantErr %v", tt.period, err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("Split(%q) = %v, want %v", tt.period, got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Split(%q)[%d] = %q, want %q", tt.period, i, got[i], tt.want[i])
				}
			}
		})
	}
}
