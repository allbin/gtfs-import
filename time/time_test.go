package time

import "testing"

func TestAddHoursToTimeString(t *testing.T) {
	tests := []struct {
		time     string
		sep      string
		expected string
	}{
		{"00:00:00", ":", "24:00:00"},
		{"28:00:00", ":", "28:00:00"},
		{"28,00,00", ",", "28,00,00"},
	}
	for _, test := range tests {
		total := FormatTimeString(test.time, test.sep)
		if total != test.expected {
			t.Errorf("Result to AddHoursToString was incorrect, got: %s, want: %s.", total, test.expected)
		}
	}

}
