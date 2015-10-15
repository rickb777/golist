// Generated by: setup
// TypeWriter: List
// Directive: +test on Foo1

package main

import (
	"errors"
	"math/rand"
	"sort"
)

// Foo1Seq is an interface for sequences of type Foo1, including lists and options (where present).
type Foo1Seq interface {
	// Len gets the size/length of the sequence.
	Len() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool

	// Exists returns true if there exists at least one element in the sequence that matches
	// the predicate supplied.
	Exists(predicate func(Foo1) bool) bool

	// Forall returns true if every element in the sequence matches the predicate supplied.
	Forall(predicate func(Foo1) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Foo1))

	// Filter returns a new Foo1Seq whose elements return true for a predicate function.
	Filter(predicate func(Foo1) bool) (result Foo1Seq)

	// Partition returns two new Foo1Lists whose elements return true or false for the predicate, p.
	// The first result consists of all elements that satisfy the predicate and the second result consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original list.
	Partition(p func(Foo1) bool) (matching Foo1Seq, others Foo1Seq)

	// Find searches for the first value that matches a given predicate. It may or may not find one.
	Find(predicate func(Foo1) bool) OptionalFoo1

	// Converts the sequence to a list. For lists, this is merely a type conversion.
	ToList() Foo1List

	// Contains tests whether a given value is present in the sequence.
	// Omitted if Foo1 is not comparable.
	Contains(value Foo1) bool

	// Count counts the number of times a given value occurs in the sequence.
	// Omitted if Foo1 is not comparable.
	Count(value Foo1) int

	// Distinct returns a new Foo1Seq whose elements are all unique.
	// Omitted if Foo1 is not comparable.
	Distinct() Foo1Seq

	// Sum sums Foo1 elements.
	// Omitted if Foo1 is not numeric.
	Sum() Foo1
}

//-------------------------------------------------------------------------------------------------
// Foo1List is a slice of type Foo1. Use it where you would use []Foo1.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type Foo1List []Foo1

//-------------------------------------------------------------------------------------------------

// Len returns the number of items in the list.
// There is no Size() method; use Len() instead.
// This is one of the three methods in the standard sort.Interface.
func (list Foo1List) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is neceessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list Foo1List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// These methods require that Foo1 be ordered.

// Less determines whether one specified element is less than another specified element.
// This is one of the three methods in the standard sort.Interface.
func (list Foo1List) Less(i, j int) bool {
	return list[i] < list[j]
}

// Sort returns a new ordered Foo1List.
func (list Foo1List) Sort() Foo1List {
	result := make(Foo1List, len(list))
	copy(result, list)
	sort.Sort(result)
	return result
}

// IsSorted reports whether Foo1List is sorted.
func (list Foo1List) IsSorted() bool {
	return sort.IsSorted(list)
}

// SortDesc returns a new reverse-ordered Foo1List.
func (list Foo1List) SortDesc() Foo1List {
	result := make(Foo1List, len(list))
	copy(result, list)
	sort.Sort(sort.Reverse(result))
	return result
}

// IsSortedDesc reports whether Foo1List is reverse-sorted.
func (list Foo1List) IsSortedDesc() bool {
	return sort.IsSorted(sort.Reverse(list))
}

// IsEmpty tests whether Foo1List is empty.
func (list Foo1List) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether Foo1List is empty.
func (list Foo1List) NonEmpty() bool {
	return len(list) > 0
}

// ToList simply returns the list in this case, but is part of the Seq interface.
func (list Foo1List) ToList() Foo1List {
	return list
}

// Exists verifies that one or more elements of Foo1List return true for the passed func.
func (list Foo1List) Exists(fn func(Foo1) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of Foo1List return true for the passed func.
func (list Foo1List) Forall(fn func(Foo1) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over Foo1List and executes the passed func against each element.
func (list Foo1List) Foreach(fn func(Foo1)) {
	for _, v := range list {
		fn(v)
	}
}

// Reverse returns a copy of Foo1List with all elements in the reverse order.
func (list Foo1List) Reverse() Foo1List {
	numItems := len(list)
	result := make(Foo1List, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of Foo1List, using a version of the Fisher-Yates shuffle.
func (list Foo1List) Shuffle() Foo1List {
	numItems := len(list)
	result := make(Foo1List, numItems)
	copy(result, list)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.Swap(i, r)
	}
	return result
}

// Take returns a new Foo1List containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Foo1List) Take(n int) Foo1List {
	if n > len(list) {
		return list
	} else {
		return list[0:n]
	}
}

// Drop returns a new Foo1List without the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Foo1List) Drop(n int) Foo1List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[n:]
	}
}

// TakeLast returns a new Foo1List containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Foo1List) TakeLast(n int) Foo1List {
	l := len(list)
	if n > l {
		return list
	} else {
		return list[l-n:]
	}
}

// DropLast returns a new Foo1List without the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Foo1List) DropLast(n int) Foo1List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new Foo1List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list Foo1List) TakeWhile(p func(Foo1) bool) (result Foo1List) {
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new Foo1List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list Foo1List) DropWhile(p func(Foo1) bool) (result Foo1List) {
	adding := false
	for _, v := range list {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}
	return
}

// Filter returns a new Foo1List whose elements return true for func.
func (list Foo1List) Filter(fn func(Foo1) bool) Foo1Seq {
	result := make(Foo1List, 0, len(list)/2)
	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Partition returns two new Foo1Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list Foo1List) Partition(p func(Foo1) bool) (Foo1Seq, Foo1Seq) {
	matching := make(Foo1List, 0, len(list)/2)
	others := make(Foo1List, 0, len(list)/2)
	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of Foo1List that return true for the passed predicate.
func (list Foo1List) CountBy(predicate func(Foo1) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of Foo1List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Returns error if no elements.
func (list Foo1List) MinBy(less func(Foo1, Foo1) bool) (result Foo1, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("Cannot determine the MinBy of an empty list.")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(list[i], list[m]) {
			m = i
		}
	}
	result = list[m]
	return
}

// MaxBy returns an element of Foo1List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements.
func (list Foo1List) MaxBy(less func(Foo1, Foo1) bool) (result Foo1, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("Cannot determine the MaxBy of an empty list.")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if list[i] != list[m] && !less(list[i], list[m]) {
			m = i
		}
	}
	result = list[m]
	return
}

// DistinctBy returns a new Foo1List whose elements are unique, where equality is defined by a passed func.
func (list Foo1List) DistinctBy(equal func(Foo1, Foo1) bool) (result Foo1List) {
Outer:
	for _, v := range list {
		for _, r := range result {
			if equal(v, r) {
				continue Outer
			}
		}
		result = append(result, v)
	}
	return result
}

// These methods require Foo1 be comparable.

// Contains verifies that a given value is contained in Foo1List.
func (list Foo1List) Contains(value Foo1) bool {
	for _, v := range list {

		if v == value {
			return true
		}

	}
	return false
}

// Count gives the number elements of Foo1List that match a certain value.
func (list Foo1List) Count(value Foo1) (result int) {
	for _, v := range list {

		if v == value {
			result++
		}

	}
	return
}

// Distinct returns a new Foo1List whose elements are unique.
func (list Foo1List) Distinct() Foo1Seq {
	result := make(Foo1List, 0)
	appended := make(map[Foo1]bool)
	for _, v := range list {

		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}

	}
	return result
}

// These methods require Foo1 be numeric.

// Sum sums Foo1 elements in Foo1List.
func (list Foo1List) Sum() (result Foo1) {
	for _, v := range list {
		result += v
	}
	return
}

// Mean sums Foo1List over all elements and divides by len(Foo1List).
func (list Foo1List) Mean() (Foo1, error) {
	var result Foo1

	l := len(list)
	if l == 0 {
		return result, errors.New("cannot determine Mean of zero-length Foo1List")
	}
	for _, v := range list {
		result += v
	}
	result = result / Foo1(l)
	return result, nil
}

// These methods require Foo1 be ordered.

// Min returns the minimum value of Foo1List. In the case of multiple items being equally minimal,
// the first such element is returned. Returns error if no elements.
func (list Foo1List) Min() (result Foo1, err error) {
	if len(list) == 0 {
		err = errors.New("Cannot determine the Min of an empty list.")
		return
	}
	result = list[0]
	for _, v := range list {
		if v < result {
			result = v
		}
	}
	return
}

// Max returns the maximum value of Foo1List. In the case of multiple items being equally maximal,
// the first such element is returned. Returns error if no elements.
func (list Foo1List) Max() (result Foo1, err error) {
	if len(list) == 0 {
		err = errors.New("Cannot determine the Max of an empty list.")
		return
	}
	result = list[0]
	for _, v := range list {
		if v > result {
			result = v
		}
	}
	return
}

// optionForList

// First returns the first element that returns true for the passed func. Returns error if no elements return true.
func (list Foo1List) Find(fn func(Foo1) bool) OptionalFoo1 {
	for _, v := range list {
		if fn(v) {
			//return SomeFoo1(v)
		}
	}
	return NoFoo1()
}

// HeadOption gets the first item in the list, provided there is one.
func (list Foo1List) HeadOption() OptionalFoo1 {
	l := len(list)
	if l > 0 {
		return SomeFoo1(list[0])
	} else {
		return NoFoo1()
	}
}

// TailOption gets the last item in the list, provided there is one.
func (list Foo1List) TailOption() OptionalFoo1 {
	l := len(list)
	if l > 0 {
		return SomeFoo1(list[l-1])
	} else {
		return NoFoo1()
	}
}

//-------------------------------------------------------------------------------------------------
// OptionalFoo1 is an optional of type Foo1. Use it where you want to be explicit about
// the presence or absence of data.
//
// Optional values follow a similar pattern to Scala Options.
// See e.g. http://www.scala-lang.org/api/2.11.7/index.html#scala.Option

type OptionalFoo1 struct {
	x *Foo1
}

// shared none value
var noneFoo1 = OptionalFoo1{nil}

func NoFoo1() OptionalFoo1 {
	return noneFoo1
}

func SomeFoo1(x Foo1) OptionalFoo1 {

	return OptionalFoo1{&x}

}

//-------------------------------------------------------------------------------------------------

func (o OptionalFoo1) Get() Foo1 {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return *o.x
}

func (o OptionalFoo1) GetOrElse(d func() Foo1) Foo1 {
	if o.IsEmpty() {
		return d()
	}
	return *o.x
}

func (o OptionalFoo1) OrElse(alternative func() OptionalFoo1) OptionalFoo1 {
	if o.IsEmpty() {
		return alternative()
	}
	return o
}

//----- Foo1Seq Methods -----

func (o OptionalFoo1) Len() int {
	if o.IsEmpty() {
		return 0
	}
	return 1
}

func (o OptionalFoo1) IsEmpty() bool {
	return o.x == nil
}

func (o OptionalFoo1) NonEmpty() bool {
	return o.x != nil
}

func (o OptionalFoo1) Find(predicate func(Foo1) bool) OptionalFoo1 {
	if o.IsEmpty() {
		return o
	}
	if predicate(*o.x) {
		return o
	}
	return noneFoo1
}

func (o OptionalFoo1) Exists(predicate func(Foo1) bool) bool {
	if o.IsEmpty() {
		return false
	}
	return predicate(*o.x)
}

func (o OptionalFoo1) Forall(predicate func(Foo1) bool) bool {
	if o.IsEmpty() {
		return true
	}
	return predicate(*o.x)
}

func (o OptionalFoo1) Foreach(fn func(Foo1)) {
	if o.NonEmpty() {
		fn(*o.x)
	}
}

func (o OptionalFoo1) Filter(predicate func(Foo1) bool) Foo1Seq {
	return o.Find(predicate)
}

func (o OptionalFoo1) Partition(predicate func(Foo1) bool) (Foo1Seq, Foo1Seq) {
	if o.IsEmpty() {
		return o, o
	}
	if predicate(*o.x) {
		return o, noneFoo1
	}
	return noneFoo1, o
}

func (o OptionalFoo1) Contains(value Foo1) bool {
	if o.IsEmpty() {
		return false
	}
	return *(o.x) == value
}

func (o OptionalFoo1) Count(value Foo1) int {
	if o.Contains(value) {
		return 1
	}
	return 0
}

// Distinct returns a new Foo1Seq whose elements are all unique. For options, this simply returns the receiver.
// Omitted if Foo1 is not comparable.
func (o OptionalFoo1) Distinct() Foo1Seq {
	return o
}

// Sum sums Foo1 elements.
// Omitted if Foo1 is not numeric.
func (o OptionalFoo1) Sum() Foo1 {

	if o.IsEmpty() {
		return 0
	}
	return *(o.x)

}

func (o OptionalFoo1) ToList() Foo1List {
	if o.IsEmpty() {
		return Foo1List{}
	}
	return Foo1List{*o.x}
}
