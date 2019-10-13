package equal

import (
	"testing"
)

func TestStringsNoCase(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expequal bool
	}{
		{"t1", "this foo isn't any bar", "THIS FOO ISN'T anY Bar", true},
		{"t2", "this foo isn't any bars", "THIS FOO ISN'T anY Bar's", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := StringsNoCase(test.s1, test.s2)
			if res != test.expequal {
				t.Errorf("expect (%t) got (%t)", test.expequal, res)
			}
		})
	}
}
