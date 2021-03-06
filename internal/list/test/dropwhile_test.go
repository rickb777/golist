package main

import (
	"reflect"
	"testing"
)

func TestDropWhileNum(t *testing.T) {
	things := num1Collection(1, 3, 17, 5, 6, 17, 8, 9)

	where1 := things.DropWhile(func(x Num1) bool {
		return x % 2 == 1
	})

	expected1 := num1Collection(6, 17, 8, 9)

	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("DropWhile should result in %#v, got %#v", expected1, where1)
	}

	where2 := things.DropWhile(func(x Num1) bool {
		return x == 0
	})

	if !reflect.DeepEqual(where2, things) {
		t.Errorf("DropWhile should result in %#v, got %#v", things, where2)
	}

	where3 := things.DropWhile(func(x Num1) bool {
		return true
	})

	if len(where3) != 0 {
		t.Errorf("DropWhile should result in empty list, got %#v", where3)
	}

	where4 := num1Collection().DropWhile(func(x Num1) bool {
		return true
	})

	if len(where4) != 0 {
		t.Errorf("DropWhile should result in empty list, got %#v", where4)
	}
}

func TestDropWhileThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	where1 := things.DropWhile(func(x Thing) bool {
		return x.Name[0] == 'F'
	})

	expected1 := thingCollection(
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)


	if !reflect.DeepEqual(where1, expected1) {
		t.Errorf("DropWhile should result in %#v, got %#v", expected1, where1)
	}

	where2 := things.DropWhile(func(x Thing) bool {
		return true
	})

	if len(where2) != 0 {
		t.Errorf("DropWhile should result in empty list, got %#v", where2)
	}

	where3 := things.DropWhile(func(x Thing) bool {
		return x.Name == "Dummy"
	})

	if !reflect.DeepEqual(where3, things) {
		t.Errorf("DropWhile should result in %#v, got %#v", things, where3)
	}

	where4 := thingCollection().DropWhile(func(x Thing) bool {
		return true
	})

	if len(where4) != 0 {
		t.Errorf("DropWhile should result in empty list, got %#v", where4)
	}
}
