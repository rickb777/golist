// Generated by: setup
// TypeWriter: Option
// Directive: +test on *Bar

package main

// BarSeq is an interface for sequences of type *Bar, including lists and options (where present).
type BarSeq interface {
	// Gets the first element from the sequence. This panics if the sequence is empty.
	Head() *Bar

	// Gets the last element from the sequence. This panics if the sequence is empty.
	Last() *Bar

	// Gets the remainder after the first element from the sequence. This panics if the sequence is empty.
	Tail() BarSeq

	// Gets everything except the last element from the sequence. This panics if the sequence is empty.
	Init() BarSeq

	// Len gets the size/length of the sequence.
	Len() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool

	// Exists returns true if there exists at least one element in the sequence that matches
	// the predicate supplied.
	Exists(predicate func(*Bar) bool) bool

	// Forall returns true if every element in the sequence matches the predicate supplied.
	Forall(predicate func(*Bar) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(*Bar))

	// Filter returns a new BarSeq whose elements return true for a predicate function.
	Filter(predicate func(*Bar) bool) (result BarSeq)

	// Partition returns two new BarLists whose elements return true or false for the predicate, p.
	// The first result consists of all elements that satisfy the predicate and the second result consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original list.
	Partition(p func(*Bar) bool) (matching BarSeq, others BarSeq)

	// Find searches for the first value that matches a given predicate. It may or may not find one.
	Find(predicate func(*Bar) bool) OptionalBar

	// Tests whether this sequence has the same length and the same elements as another sequence.
	// Omitted if Bar is not comparable.
	Equals(other BarSeq) bool

	// Contains tests whether a given value is present in the sequence.
	// Omitted if Bar is not comparable.
	Contains(value *Bar) bool

	// Count counts the number of times a given value occurs in the sequence.
	// Omitted if Bar is not comparable.
	Count(value *Bar) int
}

//-------------------------------------------------------------------------------------------------
// OptionalBar is an optional of type *Bar. Use it where you want to be explicit about
// the presence or absence of data.
//
// Optional values follow a similar pattern to Scala Options.
// See e.g. http://www.scala-lang.org/api/2.11.7/index.html#scala.Option

type OptionalBar struct {
	x *Bar
}

// shared none value
var noneBar = OptionalBar{nil}

func NoBar() OptionalBar {
	return noneBar
}

func SomeBar(x *Bar) OptionalBar {

	if x == nil {
		return NoBar()
	}
	return OptionalBar{x}

}

//-------------------------------------------------------------------------------------------------

// panics if option is empty
func (o OptionalBar) Head() *Bar {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return (o.x)
}

// panics if option is empty
func (o OptionalBar) Last() *Bar {
	return o.Head()
}

// panics if option is empty
func (o OptionalBar) Tail() BarSeq {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return noneBar
}

// panics if option is empty
func (o OptionalBar) Init() BarSeq {
	return o.Tail()
}

//-------------------------------------------------------------------------------------------------

func (o OptionalBar) Get() *Bar {
	return o.Head()
}

func (o OptionalBar) GetOrElse(d func() *Bar) *Bar {
	if o.IsEmpty() {
		return d()
	}
	return o.x
}

func (o OptionalBar) OrElse(alternative func() OptionalBar) OptionalBar {
	if o.IsEmpty() {
		return alternative()
	}
	return o
}

//-------------------------------------------------------------------------------------------------

func (o OptionalBar) Len() int {
	if o.IsEmpty() {
		return 0
	}
	return 1
}

func (o OptionalBar) IsEmpty() bool {
	return o.x == nil
}

func (o OptionalBar) NonEmpty() bool {
	return o.x != nil
}

// IsDefined returns true if the option is defined, i.e. non-empty. This is an alias for NonEmpty().
func (o OptionalBar) IsDefined() bool {
	return o.NonEmpty()
}

//-------------------------------------------------------------------------------------------------

func (o OptionalBar) Find(predicate func(*Bar) bool) OptionalBar {
	if o.IsEmpty() {
		return o
	}
	if predicate(o.x) {
		return o
	}
	return noneBar
}

func (o OptionalBar) Exists(predicate func(*Bar) bool) bool {
	if o.IsEmpty() {
		return false
	}
	return predicate(o.x)
}

func (o OptionalBar) Forall(predicate func(*Bar) bool) bool {
	if o.IsEmpty() {
		return true
	}
	return predicate(o.x)
}

func (o OptionalBar) Foreach(fn func(*Bar)) {
	if o.NonEmpty() {
		fn(o.x)
	}
}

func (o OptionalBar) Filter(predicate func(*Bar) bool) BarSeq {
	return o.Find(predicate)
}

func (o OptionalBar) Partition(predicate func(*Bar) bool) (BarSeq, BarSeq) {
	if o.IsEmpty() {
		return o, o
	}
	if predicate(o.x) {
		return o, noneBar
	}
	return noneBar, o
}

//-------------------------------------------------------------------------------------------------
// These methods require *Bar be comparable.

// Equals verifies that one or more elements of BarList return true for the passed func.
func (o OptionalBar) Equals(other BarSeq) bool {
	if o.IsEmpty() {
		return other.IsEmpty()
	}
	if other.IsEmpty() || other.Len() > 1 {
		return false
	}
	a := o.Head()
	b := other.Head()
	return *a == *b
}

func (o OptionalBar) Contains(value *Bar) bool {
	if o.IsEmpty() {
		return false
	}
	return *(o.x) == *value
}

func (o OptionalBar) Count(value *Bar) int {
	if o.Contains(value) {
		return 1
	}
	return 0
}

// Distinct returns a new BarSeq whose elements are all unique. For options, this simply returns the receiver.
// Omitted if Bar is not comparable.
func (o OptionalBar) Distinct() BarSeq {
	return o
}
