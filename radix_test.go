package cmux

import (
	"testing"
)

func TestRadixTree(t *testing.T) {
	strs := []string{"foo", "far", "farther", "boo", "ba", "bar"}
	pt := newRadixTreeString(strs...)
	for _, s := range strs {
		if !pt.match([]byte(s), false) {
			t.Errorf("%s is not matched by %s", s, s)
		}

		if !pt.match([]byte(s+s), true) {
			t.Errorf("%s is not matched as a prefix by %s", s+s, s)
		}

		if pt.match([]byte(s+s), false) {
			t.Errorf("%s matches %s", s+s, s)
		}

		// The following tests are just to catch index out of
		// range and off-by-one errors and not the functionality.
		pt.match([]byte(s[:len(s)-1]), true)
		pt.match([]byte(s[:len(s)-1]), false)
		pt.match([]byte(s+"$"), true)
		pt.match([]byte(s+"$"), false)
	}
}

func TestRadixTreeMatch(t *testing.T) {
	return
	tests := []struct {
		prefixes []string
		val      string
		prefix   bool
		exp      bool
	}{
		{nil, "a", false, false}, // 0
		{nil, "", false, true},
		{nil, "a", true, true},
		{[]string{}, "", false, true},
		{[]string{}, "", true, true},

		{[]string{""}, "", false, true}, // 5
		{[]string{""}, "", true, true},

		{[]string{"a"}, "", false, false}, // 7
		{[]string{"a"}, "", true, false},
		{[]string{"a", "aa"}, "", false, false},
		{[]string{"a", "aa"}, "", true, false},

		{[]string{"a"}, "a", false, true}, // 11
		{[]string{"a"}, "a", true, true},
		{[]string{"b", "aa"}, "a", false, false},
		{[]string{"b", "aa"}, "a", true, false},
		{[]string{"aa", "b"}, "a", false, false},
		{[]string{"aa", "b"}, "a", true, false},

		{[]string{"foo", "bar"}, "foo", true, true}, // 17
		{[]string{"foo", "bar"}, "foo", false, true},
		{[]string{"foo", "bar"}, "bar", true, true},
		{[]string{"foo", "bar"}, "bar", false, true},

		{[]string{"foo ", "bar"}, "foo", true, false}, // 21
		{[]string{"foo ", "bar"}, "foo", false, false},
		{[]string{"foo ", "bar"}, "bar", true, true},
		{[]string{"foo ", "bar"}, "bar", false, true},

		{[]string{" foo ", "bar"}, "foo", true, false}, // 25
		{[]string{" foo ", "bar"}, "foo", false, false},
		{[]string{" foo ", "bar"}, "bar", true, true},
		{[]string{" foo ", "bar"}, "bar", false, true},

		{[]string{"foo", "food"}, "foo", true, true}, // 29
		{[]string{"foo", "food"}, "food", true, true},
		{[]string{"foo", "food"}, "foo", false, true},
		{[]string{"foo", "food"}, "food", false, true},

		{[]string{"food", "foo"}, "foo", true, true}, // 33
		{[]string{"food", "foo"}, "food", true, true},
		{[]string{"food", "foo"}, "foo", false, true},
		{[]string{"food", "foo"}, "food", false, true},

		{[]string{"foo", "foobar"}, "foo", true, true}, // 37
		{[]string{"foo", "foobar"}, "food", true, true},
		{[]string{"foo", "foobar"}, "bar", true, false},
		{[]string{"foo", "foobar"}, "foo", false, true},
		{[]string{"foo", "foobar"}, "food", false, false},
		{[]string{"foo", "foobar"}, "bar", true, false}, // 42
	}
	for i, test := range tests {
		rt := newRadixTreeString(test.prefixes...)
		res := rt.match([]byte(test.val), test.prefix)
		if res != test.exp {
			t.Errorf("test %d for %v radixTree.match(%q, %t) should be %t", i, test.prefixes, test.val, test.prefix, test.exp)
		}
	}
}
