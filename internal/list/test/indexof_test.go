package main

import (
	"testing"
)

func TestIndexOfNum1(t *testing.T) {
	things := num1Collection(1, 3, 17, 5, 6, 17, 8, 9)

	where1 := things.IndexOf(17)

	if where1 != 2 {
		t.Errorf("IndexOf should be 2, got %#v", where1)
	}

	where2 := things.IndexOf(19)

	if where2 != -1 {
		t.Errorf("IndexOf should be -1, got %#v", where2)
	}

	where3 := num1Collection().IndexOf(1)

	if where3 != -1 {
		t.Errorf("IndexOf should be -1, got %#v", where3)
	}
}

func TestIndexOfNum2(t *testing.T) {
	things := num2Collection(ip(1), ip(3), ip(17), ip(5), ip(6), ip(17), ip(8), ip(9))

	where1 := things.IndexOf(ip(17))

	if where1 != 2 {
		t.Errorf("IndexOf should be 2, got %#v", where1)
	}

	where2 := things.IndexOf(ip(19))

	if where2 != -1 {
		t.Errorf("IndexOf should be -1, got %#v", where2)
	}

	where3 := num2Collection().IndexOf(ip(1))

	if where3 != -1 {
		t.Errorf("IndexOf should be -1, got %#v", where3)
	}
}
