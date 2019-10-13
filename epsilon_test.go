package equal

import "testing"

func TestEpsilon(t *testing.T) {
	tests := []struct {
		name     string
		epsilon  float64
		v1       float64
		v2       float64
		expequal bool
	}{
		{"t1", 1e-4, 567.67, 567.671, false},
		{"t2", 1e-4, 567.67123, 567.67124, true},
		{"t3", 10, 11, 22, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			eq := Epsilon(test.epsilon, test.v1, test.v2)
			if eq != test.expequal {
				t.Errorf("expect (%t) got (%t)", test.expequal, eq)
			}
		})
	}
}
