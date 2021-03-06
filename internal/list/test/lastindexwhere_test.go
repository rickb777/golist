package main

import (
	"testing"
)

func TestLastIndexWhereNum(t *testing.T) {
	things := num1Collection(1, 3, 17, 5, 6, 17, 8, 9)

	where1 := things.LastIndexWhere(func(x Num1) bool {
		return x == 17
	})

	if where1 != 5 {
		t.Errorf("LastIndexWhere should be 5, got %#v", where1)
	}

	where2 := things.LastIndexWhere(func(x Num1) bool {
		return false
	})

	if where2 != -1 {
		t.Errorf("LastIndexWhere should be -1, got %#v", where2)
	}

	where3 := num1Collection().LastIndexWhere(func(x Num1) bool {
		return true
	})

	if where3 != -1 {
		t.Errorf("LastIndexWhere should be -1, got %#v", where3)
	}
}

func TestLastIndexWhereThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	where1 := things.LastIndexWhere(func(x Thing) bool {
		return x.Name == "Boo"
	})

	if where1 != 4 {
		t.Errorf("LastIndexWhere should be 4, got %#v", where1)
	}

	where2 := things.LastIndexWhere(func(x Thing) bool {
		return false
	})

	if where2 != -1 {
		t.Errorf("LastIndexWhere should be -1, got %#v", where2)
	}

	where3 := thingCollection().LastIndexWhere(func(x Thing) bool {
		return true
	})

	if where3 != -1 {
		t.Errorf("LastIndexWhere should be -1, got %#v", where3)
	}
}
