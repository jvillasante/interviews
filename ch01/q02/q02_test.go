package q02

import "testing"

var cases = []struct {
	str1, str2 string
	want       bool
}{
	{"", "", true},
	{"a", "a", true},
	{"abcde", "ecdab", true},
	{"dog", "god", true},
	{"dog   ", "god", false},
	{"  dog", "god", false},
	{"abc", "def", false},
}

var tests = []func(str1, str2 string) bool{
	Permutation1,
	Permutation2,
}

func TestPermutation(t *testing.T) {
	for i, test := range tests {
		for _, c := range cases {
			got := test(c.str1, c.str2)
			if got != c.want {
				t.Errorf("Permutation%d(%s, %s) == %t, want %t", i+1, c.str1, c.str2, got, c.want)
			}
		}
	}
}
