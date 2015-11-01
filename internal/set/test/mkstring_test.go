package main

import (
	"testing"
)

func TestMkStringNum1(t *testing.T) {
	things := num1Collection(1, 2, 3, 5, 8, 13)

	s1 := things.MkString3("<", ", ", ">")
	e1 := "<1, 2, 3, 5, 8, 13>"
	if len(s1) != len(e1) || s1[0] != '<' {
		t.Errorf("MkString got %q not %q", s1, e1)
	}

	s2 := things.MkString(", ")
	e2 := "1, 2, 3, 5, 8, 13"
	if len(s2) != len(e2) || s2[0] == '[' {
		t.Errorf("MkString got %q not %q", s2, e2)
	}

	s3 := things.String()
	e3 := "[1,2,3,5,8,13]"
	if len(s3) != len(e3) || s3[0] != '[' {
		t.Errorf("MkString got %q not %q", s3, e3)
	}
}

func TestMkStringNum1Empty(t *testing.T) {
	empty := num1Collection()

	s1 := empty.MkString3("<", ", ", ">")
	e1 := "<>"
	if s1 != e1 {
		t.Errorf("MkString got %q", s1)
	}

	s2 := empty.MkString(", ")
	e2 := ""
	if s2 != e2 {
		t.Errorf("MkString got %q", s2)
	}

	s3 := empty.String()
	e3 := "[]"
	if s3 != e3 {
		t.Errorf("MkString got %q", s3)
	}
}

func TestMkStringThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	s1 := things.MkString3("<", ", ", ">")
	e1 := "<{Fee 1}, {Fie 2}, {Foe 3}, {Boo 5}, {Boo 8}, {Bam 13}, {Bam 21}>"
	if len(s1) != len(e1) || s1[0] != '<' {
		t.Errorf("MkString got %q not %q", s1, e1)
	}

	s2 := things.MkString(", ")
	e2 := "{Fee 1}, {Fie 2}, {Foe 3}, {Boo 5}, {Boo 8}, {Bam 13}, {Bam 21}"
	if len(s2) != len(e2) || s2[0] != '{' {
		t.Errorf("MkString got %q not %q", s2, e2)
	}

	s3 := things.String()
	e3 := "[{Fee 1},{Fie 2},{Foe 3},{Boo 5},{Boo 8},{Bam 13},{Bam 21}]"
	if len(s3) != len(e3) || s3[0] != '[' {
		t.Errorf("MkString got %q not %q", s3, e3)
	}
}

func TestMkStringThingEmpty(t *testing.T) {
	empty := thingCollection()

	s1 := empty.MkString3("<", ", ", ">")
	e1 := "<>"
	if s1 != e1 {
		t.Errorf("MkString got %q not %q", s1, e1)
	}

	s2 := empty.MkString(", ")
	e2 := ""
	if s2 != e2 {
		t.Errorf("MkString got %q not %q", s2, e2)
	}

	s3 := empty.String()
	e3 := "[]"
	if s3 != e3 {
		t.Errorf("MkString got %q not %q", s3, e3)
	}
}

