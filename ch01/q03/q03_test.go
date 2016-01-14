package q03

import "testing"

var cases = []struct {
	str    []rune
	length int
	want   string
}{
	{[]rune("Mr John Smith    "), 13, "Mr%20John%20Smith"},
}

func TestPermutation(t *testing.T) {
	for _, c := range cases {
		original := string(c.str)
		ReplaceSpaces(c.str, c.length)
		got := string(c.str)
		if got != c.want {
			t.Errorf("ReplaceSpaces(%s, %d) == %s, want %s", original, c.length, got, c.want)
		}
	}
}
