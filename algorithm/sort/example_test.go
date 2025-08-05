package sort

import "testing"

func TestExample(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"test count-1k", 1000},
		{"test count-1w", 10000},
		{"test count-10w", 100000},
		{"test count-100w", 1000000},
		{"test count-1000w", 10000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Example(tt.n)
		})
	}
}
