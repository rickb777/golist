package main

import "testing"

func TestExistsNum1(t *testing.T) {
	things := num1Collection(60, -20, 100)

	any1 := things.Exists(func(x Num1) bool {
		return x == 10
	})

	if any1 {
		t.Errorf("Exists should not evaluate true for 10")
	}

	any2 := things.Exists(func(x Num1) bool {
		return x > 50
	})

	if !any2 {
		t.Errorf("Exists should evaluate true for Number > 50")
	}

	any3 := num1Collection().Exists(func(x Num1) bool {
		return true
	})

	if any3 {
		t.Errorf("Exists should evaluate false for empty slices")
	}
}

func TestExistsThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	any1 := things.Exists(func(x Thing) bool {
		return x.Name == "Dummy"
	})

	if any1 {
		t.Errorf("Exists should not evaluate true for Name == Dummy")
	}

	any2 := things.Exists(func(x Thing) bool {
		return x.Number > 5
	})

	if !any2 {
		t.Errorf("Exists should evaluate true for Number > 5")
	}

	any3 := thingCollection().Exists(func(x Thing) bool {
		return true
	})

	if any3 {
		t.Errorf("Exists should evaluate false for empty slices")
	}
}
