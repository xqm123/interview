package heap

import "testing"

func TestFindForKMax(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FindForKMax()
		})
	}
}
