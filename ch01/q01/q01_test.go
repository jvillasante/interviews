package q01

import "testing"

var cases = []struct {
	in   string
	want bool
}{
	{"", true},
	{"a", true},
	{"abc", true},
	{"abcded", false},
	{"aabcde", false},
	{"abcdee", false},
}

var tests = []func(str string) bool{
	IsUnique1,
	IsUnique2,
	IsUnique3,
	IsUnique4,
}

func TestUnique(t *testing.T) {
	for i, test := range tests {
		for _, c := range cases {
			got := test(c.in)
			if got != c.want {
				t.Errorf("IsUnique%d(%s) == %t, want %t", i+1, c.in, got, c.want)
			}
		}
	}
}
