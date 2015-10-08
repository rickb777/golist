package main

import "testing"

func TestMinByOther(t *testing.T) {
	things := OtherSlice{50, 100, -20, 7, 100, 99}

	min1, err1 := things.MinBy(func(a, b Other) bool {
		return a < b
	})

	if err1 != nil {
		t.Errorf("MinBy Number should succeed")
	}

	if min1 != -20 {
		t.Errorf("MinBy Number should return %v, got %v", -20, min1)
	}

	_, err2 := OtherSlice{}.MinBy(func(a, b Other) bool {
		return true
	})

	if err2 == nil {
		t.Errorf("MinBy Number should fail on empty slice")
	}
}

func TestMinByThing(t *testing.T) {
	things := ThingSlice{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	min1, err1 := things.MinBy(func(a, b Thing) bool {
		return a.Number < b.Number
	})

	if err1 != nil {
		t.Errorf("MinBy Number should succeed")
	}

	expected1 := Thing{"Second", -20}
	if min1 != expected1 {
		t.Errorf("MinBy Number should return %v, got %v", expected1, min1)
	}

	_, err2 := ThingSlice{}.MinBy(func(a, b Thing) bool {
		return true
	})

	if err2 == nil {
		t.Errorf("MinBy Number should fail on empty slice")
	}
}
