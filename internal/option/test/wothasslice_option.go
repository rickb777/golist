// Generated by: setup
// TypeWriter: Option
// Directive: +test on WotHasSlice

package main

import "fmt"

//-------------------------------------------------------------------------------------------------

// WotHasSliceCollection is an interface for collections of type WotHasSlice, including sets, lists and options (where present).
type WotHasSliceCollection interface {
	// Size gets the size/length of the collection.
	Size() int

	// IsEmpty returns true if the collection is empty.
	IsEmpty() bool

	// NonEmpty returns true if the collection is non-empty.
	NonEmpty() bool

	// IsSequence returns true for lists, but false otherwise.
	IsSequence() bool

	// IsSet returns true for sets, but false otherwise.
	IsSet() bool

	// Head returns the first element of a list or an arbitrary element of a set or the contents of an option.
	// Panics if the collection is empty.
	Head() WotHasSlice

	// ToSlice returns a plain slice containing all the elements in the collection. This is useful for bespoke iteration etc.
	// For sequences, the order of the elements is simple and well defined.
	// For non-sequences (i.e. sets) the order of the elements is stable but not well defined. This means it will give
	// the same order each subsequent time it is used as it did the first time.
	ToSlice() []WotHasSlice

	// Send sends all elements along a channel of type WotHasSlice.
	// For sequences, the order of the elements is simple and well defined.
	// For non-sequences (i.e. sets) the order of the elements is stable but not well defined. This means it will give
	// the same order each subsequent time it is used as it did the first time.
	Send() <-chan WotHasSlice

	// Exists returns true if there exists at least one element in the collection that matches
	// the predicate supplied.
	Exists(predicate func(WotHasSlice) bool) bool

	// Forall returns true if every element in the collection matches the predicate supplied.
	Forall(predicate func(WotHasSlice) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(WotHasSlice))

	// Filter returns a new WotHasSliceCollection whose elements return true for a predicate function.
	// The relative order of the elements in the result is the same as in the
	// original collection.
	Filter(predicate func(WotHasSlice) bool) (result WotHasSliceCollection)

	// Partition returns two new WotHasSliceCollections whose elements return true or false for the predicate, p.
	// The first consists of all elements that satisfy the predicate and the second consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original collection.
	Partition(p func(WotHasSlice) bool) (matching WotHasSliceCollection, others WotHasSliceCollection)

	// Min returns an element of WotHasSliceList containing the minimum value, when compared to other elements
	// using a specified comparator function defining ‘less’. For ordered sequences, Min returns the first such element.
	// Panics if the collection is empty.
	Min(less func(WotHasSlice, WotHasSlice) bool) WotHasSlice

	// Max returns an element of WotHasSliceList containing the maximum value, when compared to other elements
	// using a specified comparator function defining ‘less’. For ordered sequences, Max returns the first such element.
	// Panics if the collection is empty.
	Max(less func(WotHasSlice, WotHasSlice) bool) WotHasSlice

	// String gets a string representation of the collection. "[" and "]" surround
	// a comma-separated list of the elements.
	String() string

	// MkString gets a string representation of the collection. "[" and "]" surround a list
	// of the elements joined by the separator you provide.
	MkString(sep string) string

	// MkString3 gets a string representation of the collection. 'pfx' and 'sfx' surround a list
	// of the elements joined by the 'mid' separator you provide.
	MkString3(pfx, mid, sfx string) string
}

//-------------------------------------------------------------------------------------------------

// OptionalWotHasSlice is an optional of type WotHasSlice. Use it where you want to be explicit about
// the presence or absence of data.
//
// Optional values follow a similar pattern to Scala Options. In particular, an option is a collection
// with a maximum cardinality of one. As such, options can be converted to/from lists and sets.
// See e.g. http://www.scala-lang.org/api/2.11.7/index.html#scala.Option
type OptionalWotHasSlice struct {
	x *WotHasSlice
}

// shared none value
var noneWotHasSlice = OptionalWotHasSlice{nil}

// NoWotHasSlice gets an empty instance.
func NoWotHasSlice() OptionalWotHasSlice {
	return noneWotHasSlice
}

// SomeWotHasSlice gets a non-empty instance wrapping some value *x*.
func SomeWotHasSlice(x WotHasSlice) OptionalWotHasSlice {

	return OptionalWotHasSlice{&x}

}

//-------------------------------------------------------------------------------------------------

// panics if option is empty
func (o OptionalWotHasSlice) Head() WotHasSlice {
	return o.Get()
}

func (o OptionalWotHasSlice) Get() WotHasSlice {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return *o.x
}

func (o OptionalWotHasSlice) GetOrElse(d func() WotHasSlice) WotHasSlice {
	if o.IsEmpty() {
		return d()
	}
	return *o.x
}

func (o OptionalWotHasSlice) OrElse(alternative func() OptionalWotHasSlice) OptionalWotHasSlice {
	if o.IsEmpty() {
		return alternative()
	}
	return o
}

//-------------------------------------------------------------------------------------------------

func (o OptionalWotHasSlice) Size() int {
	if o.IsEmpty() {
		return 0
	}
	return 1
}

func (o OptionalWotHasSlice) Len() int {
	return o.Size()
}

func (o OptionalWotHasSlice) IsEmpty() bool {
	return o.x == nil
}

func (o OptionalWotHasSlice) NonEmpty() bool {
	return o.x != nil
}

// IsSequence returns false for options.
func (o OptionalWotHasSlice) IsSequence() bool {
	return false
}

// IsSet returns false for options.
func (o OptionalWotHasSlice) IsSet() bool {
	return false
}

// IsDefined returns true if the option is defined, i.e. non-empty. This is an alias for NonEmpty().
func (o OptionalWotHasSlice) IsDefined() bool {
	return o.NonEmpty()
}

//-------------------------------------------------------------------------------------------------

func (o OptionalWotHasSlice) Find(predicate func(WotHasSlice) bool) OptionalWotHasSlice {
	if o.IsEmpty() {
		return o
	}
	if predicate(*o.x) {
		return o
	}
	return noneWotHasSlice
}

func (o OptionalWotHasSlice) Exists(predicate func(WotHasSlice) bool) bool {
	if o.IsEmpty() {
		return false
	}
	return predicate(*o.x)
}

func (o OptionalWotHasSlice) Forall(predicate func(WotHasSlice) bool) bool {
	if o.IsEmpty() {
		return true
	}
	return predicate(*o.x)
}

func (o OptionalWotHasSlice) Foreach(fn func(WotHasSlice)) {
	if o.NonEmpty() {
		fn(*o.x)
	}
}

// Send gets a channel that will send all the elements in order.
func (o OptionalWotHasSlice) Send() <-chan WotHasSlice {
	ch := make(chan WotHasSlice)
	go func() {
		if o.NonEmpty() {
			ch <- *o.x
		}
		close(ch)
	}()
	return ch
}

func (o OptionalWotHasSlice) Filter(predicate func(WotHasSlice) bool) WotHasSliceCollection {
	return o.Find(predicate)
}

func (o OptionalWotHasSlice) Partition(predicate func(WotHasSlice) bool) (WotHasSliceCollection, WotHasSliceCollection) {
	if o.IsEmpty() {
		return o, o
	}
	if predicate(*o.x) {
		return o, noneWotHasSlice
	}
	return noneWotHasSlice, o
}

func (o OptionalWotHasSlice) ToSlice() []WotHasSlice {
	slice := make([]WotHasSlice, o.Size())
	if o.NonEmpty() {
		slice[0] = *o.x
	}
	return slice
}

// Min returns an element of WotHasSliceList containing the minimum value, when compared to other elements
// using a specified comparator function defining ‘less’. For ordered sequences, Min returns the first such element.
// Panics if the collection is empty.
func (o OptionalWotHasSlice) Min(less func(WotHasSlice, WotHasSlice) bool) WotHasSlice {
	return o.Get()
}

// Max returns an element of WotHasSliceList containing the maximum value, when compared to other elements
// using a specified comparator function defining ‘less’. For ordered sequences, Max returns the first such element.
// Panics if the collection is empty.
func (o OptionalWotHasSlice) Max(less func(WotHasSlice, WotHasSlice) bool) WotHasSlice {
	return o.Get()
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the option as an array of one element.
func (o OptionalWotHasSlice) String() string {
	return o.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string.
func (o OptionalWotHasSlice) MkString(sep string) string {
	return o.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string.
func (o OptionalWotHasSlice) MkString3(pfx, mid, sfx string) string {
	if o.IsEmpty() {
		return fmt.Sprintf("%s%s", pfx, sfx)
	}
	return fmt.Sprintf("%s%v%s", pfx, *(o.x), sfx)
}

// Option flags: {Collection:false List:false Option:true Set:false Plumbing:false Tag:map[]}
