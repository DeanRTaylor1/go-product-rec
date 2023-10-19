package recommendations

import (
	"reflect"
	"testing"
)

func TestMapData(t *testing.T) {
	data := []float64{
		5, 3, 0, 1,
		4, 0, 0, 1,
		1, 1, 0, 5,
		1, 0, 0, 4,
		0, 1, 5, 4,
	}

	tests := []struct {
		userIndex int
		expected  []int
	}{
		{0, []int{2}},    // Based on the given data, user 0 hasn't interacted with item 2.
		{1, []int{1, 2}}, // User 1 hasn't interacted with items 1 and 2.
		{2, []int{2}},    // And so on...
		{3, []int{1, 2}},
		{4, []int{0}},
	}

	for _, tt := range tests {
		result := MapData(data, tt.userIndex)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("for userIndex %d, expected %v, but got %v", tt.userIndex, tt.expected, result)
		}
	}
}
