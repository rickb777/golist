// Generated by: setup
// TypeWriter: Set
// Directive: +test on Foo

package main

import (
	"bytes"
	"fmt"
)

//-------------------------------------------------------------------------------------------------
// FooCollection is an interface for collections of type Foo, including sets, lists and options (where present).
type FooCollection interface {
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
	Head() Foo

	//-------------------------------------------------------------------------
	// ToSlice returns a plain slice containing all the elements in the collection.
	// This is useful for bespoke iteration etc.
	// For sequences, the order is well defined.
	// For non-sequences (i.e. sets) the first time it is used, order of the elements is not well defined. But
	// the order is stable, which means it will give the same order each subsequent time it is used.
	ToSlice() []Foo

	// ToStrings gets all the elements in a slice of the underlying type, []string.
	ToStrings() []string

	// ToSet gets all the elements in a Set.
	ToSet() FooSet

	// Send sends all elements along a channel of type Foo.
	// For sequences, the order is well defined.
	// For non-sequences (i.e. sets) the first time it is used, order of the elements is not well defined. But
	// the order is stable, which means it will give the same order each subsequent time it is used.
	Send() <-chan Foo

	//-------------------------------------------------------------------------
	// Exists returns true if there exists at least one element in the collection that matches
	// the predicate supplied.
	Exists(predicate func(Foo) bool) bool

	// Forall returns true if every element in the collection matches the predicate supplied.
	Forall(predicate func(Foo) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Foo))

	//-------------------------------------------------------------------------
	// Filter returns a new FooCollection whose elements return true for a predicate function.
	// The relative order of the elements in the result is the same as in the
	// original collection.
	Filter(predicate func(Foo) bool) (result FooCollection)

	// Partition returns two new FooCollections whose elements return true or false for the predicate, p.
	// The first consists of all elements that satisfy the predicate and the second consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original collection.
	Partition(p func(Foo) bool) (matching FooCollection, others FooCollection)

	//-------------------------------------------------------------------------
	// Equals verifies that another FooCollection has the same size and elements as this one. Also,
	// if the collection is a sequence, the order must be the same.
	// Omitted if Foo is not comparable.
	Equals(other FooCollection) bool

	// Contains tests whether a given value is present in the collection.
	// Omitted if Foo is not comparable.
	Contains(value Foo) bool

	// Min returns the minimum value of FooList. In the case of multiple items being equally minimal,
	// the first such element is returned. Panics if the collection is empty.
	Min() Foo

	// Max returns the maximum value of FooList. In the case of multiple items being equally maximal,
	// the first such element is returned. Panics if the collection is empty.
	Max() Foo

	//-------------------------------------------------------------------------
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
// FooSet is a typesafe set of Foo items. Instances are essentially immutable.
// The set-agebra functions Union, Intersection and Difference allow new variants to be constructed
// easily.
//
// The implementation is based on Go maps.

type FooSet map[Foo]struct{}

//-------------------------------------------------------------------------------------------------
// NewFooSet constructs a new set containing the supplied values, if any.
func NewFooSet(values ...Foo) FooSet {
	set := make(map[Foo]struct{})
	for _, v := range values {
		set[v] = struct{}{}
	}
	return FooSet(set)
}

// BuildFooSetFrom constructs a new FooSet from a channel that supplies values
// until it is closed.
func BuildFooSetFrom(source <-chan Foo) FooSet {
	set := make(map[Foo]struct{})
	for v := range source {
		set[v] = struct{}{}
	}
	return FooSet(set)
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns false for sets.
func (set FooSet) IsSequence() bool {
	return false
}

// IsSet returns true for sets.
func (set FooSet) IsSet() bool {
	return true
}

func (set FooSet) Size() int {
	return len(set)
}

func (set FooSet) IsEmpty() bool {
	return len(set) == 0
}

func (set FooSet) NonEmpty() bool {
	return len(set) > 0
}

// Head gets an arbitrary element.
func (set FooSet) Head() Foo {
	for v := range set {
		return v
	}
	panic("Set is empty")
}

// ToSlice gets all the set's elements in a plain slice.
func (set FooSet) ToSlice() []Foo {
	slice := make([]Foo, set.Size())
	i := 0
	for v := range set {
		slice[i] = v
		i++
	}
	return slice
}

// ToStrings gets all the elements in a []string.
func (set FooSet) ToStrings() []string {
	slice := make([]string, len(set))
	i := 0
	for v := range set {
		slice[i] = string(v)
		i++
	}
	return slice
}

// ToSet gets the current set, which requires no further conversion.
func (set FooSet) ToSet() FooSet {
	return set
}

// Contains tests whether an item is already in the FooSet.
func (set FooSet) Contains(i Foo) bool {
	_, found := set[i]
	return found
}

// ContainsAll tests whether many items are all in the FooSet.
func (set FooSet) ContainsAll(i ...Foo) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

func (set FooSet) actualSubset(other FooSet) bool {
	for item := range set {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// Equals determines if two sets are equal to each other.
// They are considered equal if both are the same size and both have the same items.
func (set FooSet) Equals(other FooCollection) bool {
	otherSet, isSet := other.(FooSet)
	return isSet && set.Size() == other.Size() && set.actualSubset(otherSet)
}

// IsSubset determines if every item in the other set is in this set.
func (set FooSet) IsSubset(other FooSet) bool {
	return set.Size() <= other.Size() && set.actualSubset(other)
}

// IsProperSubset determines if every item in the other set is in this set and this set is
// smaller than the other.
func (set FooSet) IsProperSubset(other FooSet) bool {
	return set.Size() < other.Size() && set.actualSubset(other)
}

// IsSuperset determines if every item of this set is in the other set.
func (set FooSet) IsSuperset(other FooSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set FooSet) Union(other FooSet) FooSet {
	union := NewFooSet()
	for item := range set {
		union[item] = struct{}{}
	}
	for item := range other {
		union[item] = struct{}{}
	}
	return union
}

// Intersection returns a new set with items that exist only in both sets.
func (set FooSet) Intersection(other FooSet) FooSet {
	intersection := NewFooSet()
	// loop over the smaller set
	if set.Size() < other.Size() {
		for item := range set {
			if other.Contains(item) {
				intersection[item] = struct{}{}
			}
		}
	} else {
		for item := range other {
			if set.Contains(item) {
				intersection[item] = struct{}{}
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set FooSet) Difference(other FooSet) FooSet {
	diffs := NewFooSet()
	for item := range set {
		if !other.Contains(item) {
			diffs[item] = struct{}{}
		}
	}
	return diffs
}

// Add creates a new set with elements added. This is similar to Union, but takes a slice of extra values.
// The receiver is not modified.
func (set FooSet) Add(others ...Foo) FooSet {
	added := NewFooSet()
	for item := range set {
		added[item] = struct{}{}
	}
	for _, item := range others {
		added[item] = struct{}{}
	}
	return added
}

// Remove creates a new set with elements removed. This is similar to Difference, but takes a slice of unwanted values.
// The receiver is not modified.
func (set FooSet) Remove(unwanted ...Foo) FooSet {
	removed := NewFooSet()
	for item := range set {
		removed[item] = struct{}{}
	}
	for _, item := range unwanted {
		delete(removed, item)
	}
	return removed
}

// Exists verifies that one or more elements of FooSet return true for the passed func.
func (set FooSet) Exists(fn func(Foo) bool) bool {
	for v := range set {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of FooSet return true for the passed func.
func (set FooSet) Forall(fn func(Foo) bool) bool {
	for v := range set {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over FooSet and executes the passed func against each element.
// The order of the elements is not well defined but is probably repeatably stable until the set is changed.
func (set FooSet) Foreach(fn func(Foo)) {
	for v := range set {
		fn(v)
	}
}

// Send sends all elements along a channel of type Foo.
// The order of the elements is not well defined but is probably repeatably stable until the set is changed.
func (set FooSet) Send() <-chan Foo {
	ch := make(chan Foo)
	go func() {
		for v := range set {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Filter returns a new FooSet whose elements return true for func.
func (set FooSet) Filter(fn func(Foo) bool) FooCollection {
	result := make(map[Foo]struct{})
	for v := range set {
		if fn(v) {
			result[v] = struct{}{}
		}
	}
	return FooSet(result)
}

// Partition returns two new FooLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original set.
func (set FooSet) Partition(p func(Foo) bool) (FooCollection, FooCollection) {
	matching := make(map[Foo]struct{})
	others := make(map[Foo]struct{})
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return FooSet(matching), FooSet(others)
}

// CountBy gives the number elements of FooSet that return true for the passed predicate.
func (set FooSet) CountBy(predicate func(Foo) bool) (result int) {
	for v := range set {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of FooSet containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set FooSet) MinBy(less func(Foo, Foo) bool) (result Foo) {
	l := len(set)
	if l == 0 {
		panic("Cannot determine the minimum of an empty set.")
	}
	first := true
	for v := range set {
		if first {
			first = false
			result = v
		} else if less(v, result) {
			result = v
		}
	}
	return
}

// MaxBy returns an element of FooSet containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Panics if there are no elements.
func (set FooSet) MaxBy(less func(Foo, Foo) bool) (result Foo) {
	l := len(set)
	if l == 0 {
		panic("Cannot determine the maximum of an empty set.")
	}
	first := true
	for v := range set {
		if first {
			first = false
			result = v
		} else if less(result, v) {
			result = v
		}
	}
	return
}

//-------------------------------------------------------------------------------------------------
// These methods require Foo be ordered.

// Min returns the element with the minimum value. In the case of multiple items being equally minimal,
// any such element is returned. Panics if the collection is empty.
func (set FooSet) Min() (result Foo) {
	if len(set) == 0 {
		panic("Cannot determine the minimum of an empty set.")
	}
	first := true
	for v := range set {
		if first {
			first = false
			result = v
		} else if v < result {
			result = v
		}
	}
	return
}

// Max returns the element with the maximum value. In the case of multiple items being equally maximal,
// any such element is returned. Panics if the collection is empty.
func (set FooSet) Max() (result Foo) {
	if len(set) == 0 {
		panic("Cannot determine the maximum of an empty set.")
	}
	first := true
	for v := range set {
		if first {
			first = false
			result = v
		} else if v > result {
			result = v
		}
	}
	return
}

// String implements the Stringer interface to render the set as a comma-separated array.
func (set FooSet) String() string {
	return set.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string.
func (set FooSet) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string.
func (set FooSet) MkString3(pfx, mid, sfx string) string {
	b := bytes.Buffer{}
	b.WriteString(pfx)
	l := len(set)
	if l > 0 {
		sep := ""
		for v := range set {
			b.WriteString(sep)
			b.WriteString(fmt.Sprintf("%v", v))
			sep = mid
		}
	}
	b.WriteString(sfx)
	return b.String()
}

// MapToNum1 transforms FooSet to Num1Set.
func (set FooSet) MapToNum1(fn func(Foo) Num1) Num1Collection {
	result := make(map[Num1]struct{})
	for v := range set {
		u := fn(v)
		result[u] = struct{}{}
	}
	return Num1Set(result)
}

// FlatMapToNum1 transforms FooSet to Num1Set, by
// calling the supplied function on each of the enclosed set elements, and returning a new set.
func (set FooSet) FlatMapToNum1(fn func(Foo) Num1Collection) Num1Collection {
	result := make(map[Num1]struct{})
	for a := range set {
		b := fn(a)
		b.Foreach(func(c Num1) {
			result[c] = struct{}{}
		})
	}
	return Num1Set(result)
}

// GroupByNum1 groups elements into a map keyed by Num1.
// This method requires Num1 be comparable.
func (set FooSet) GroupByNum1(fn func(Foo) Num1) map[Num1]FooSet {
	result := make(map[Num1]FooSet)
	for v := range set {
		key := fn(v)
		group, exists := result[key]
		if !exists {
			group = NewFooSet()
		}
		group[v] = struct{}{}
		result[key] = group
	}
	return result
}

// Set flags: {Collection:false Sequence:false List:false Option:false Set:true Tag:map[MapTo:true]}
