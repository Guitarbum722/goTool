package tool

import "testing"

var isToolCases = []struct {
	input    string
	expected bool
}{
	{
		"Juan",
		false,
	},
	{
		"Sean A",
		true,
	},
}

func TestIsTool(t *testing.T) {
	for _, tt := range isToolCases {
		got := IsTool([]byte(tt.input))

		if got != tt.expected {
			t.Fatalf("IsTool(%v) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}
