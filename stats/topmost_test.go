package stats

import (
	"reflect"
	"testing"
)

func Test_TopMost(t *testing.T) {
        n:= 3

	tests := map[string]struct {
		data     map[string]int
		expected []string
	}{
		"elements over than configured count": {
			data: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 4,
			},
			expected: []string{"d", "c", "b"},
		},
		"elements has same value": {
			data: map[string]int{
				"a": 1,
				"b": 1,
				"c": 3,
				"d": 4,
			},
			expected: []string{"d", "c", "a"},
		},
		"elements less than configured value": {
			data: map[string]int{
				"b": 1,
				"c": 3,
			},
			expected: []string{"c", "b"},
		},
	}

        for name, test := range tests {
                t.Run(name, func(t *testing.T) {
                        actual := TopMost(test.data, n)

                        if reflect.DeepEqual(test.expected, actual) != true {
                                t.Errorf("Expected %v, got %v", test.expected, actual)
                        }
                })
        }
}
