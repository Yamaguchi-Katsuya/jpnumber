package jpnumber

import (
	"fmt"
	"testing"
)

func TestToJPNumber(t *testing.T) {
	tests := []struct {
		input    int
		expected string
		err      error
	}{
		{1, "一", nil},
		{10, "十", nil},
		{11, "十一", nil},
		{100, "百", nil},
		{1000, "千", nil},
		{10000, "一万", nil},
		{250000, "二十五万", nil},
		{100000, "十万", nil},
		{33100000, "三千三百十万", nil},
		{100000000, "一億", nil},
		{1000000000, "十億", nil},
		{10000000000, "百億", nil},
		{22232222222, "二百二十二億三千二百二十二万二千二百二十二", nil},
		{1000000000000, "一兆", nil},
		{22222222222222, "二十二兆二千二百二十二億二千二百二十二万二千二百二十二", nil},
		{99_999_999_999_999_999, "九京九千九百九十九兆九千九百九十九億九千九百九十九万九千九百九十九", nil},
		{100_000_000_000_000_000, "", fmt.Errorf("number is too large: 100000000000000000. max is 99999999999999999")},
	}

	for _, test := range tests {
		result, err := ToJPNumber(test.input)
		if err != nil {
			if test.err == nil {
				t.Errorf("ToJPNumber(%d) = %s; expected %s", test.input, err, test.expected)
			}
		} else {
			if test.err != nil {
				t.Errorf("ToJPNumber(%d) = %s; expected %s", test.input, result, test.expected)
			}
		}
		if result != test.expected {
			t.Errorf("ToJPNumber(%d) = %s; expected %s", test.input, result, test.expected)
		}
	}
}
