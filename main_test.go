/*package main

import "testing"

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{0, "zero"},
	{5, "short"},
	{11, "long"},
	{101, "very long"},
}

func TestLength(t *testing.T) {
	for i, test := range tests {
		size := Length(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}
*/
/*
package main

import "testing"

func TestMultiply(t *testing.T) {
	// Определяем тестовые случаи
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 2},
	}
	for _, test := range tests {
		result := Multiply(test.a, test.b)
		if result != test.a*test.b {
			t.Errorf("Multiply(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}
*/
/*
package main

import "testing"

func TestDeleteVowels(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{"Aa", ""},
		{"Ee", ""},
		{"Oo", ""},
		{"Ii", ""},
		{"Uu", ""},
		{"Hello", "Hll"},
		{"0123", "0123"},
		{"", ""},
	}
	for _, test := range tests {
		result := DeleteVowels(test.s)
		if result != test.expected {
			t.Error("Ошибка")
			//t.Errorf("Multiply(%s) = %s; expected %s", test.s, result, test.expected)
		}
	}
}
*/

package main

import "testing"

func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		input    []byte
		expected int
		err      error
	}{
		{[]byte("jkds"), 4, nil},
		{[]byte(""), 0, ErrInvalidUTF8},
		{[]byte{0xff, 0xfe, 0xfd}, 0, ErrInvalidUTF8},
	}
	for _, test := range tests {
		result, err := GetUTFLength(test.input)
		if result != test.expected || err != nil && err.Error() != test.err.Error() {
			t.Error("Ошибка")
		}
	}
}
