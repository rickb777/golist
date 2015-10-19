// Generated by: setup
// TypeWriter: Option
// Directive: +test on Foo2

package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

// Foo2Seq is an interface for sequences of type Foo2, including lists and options (where present).
type Foo2Seq interface {
	// Len gets the size/length of the sequence.
	Len() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool

	//-------------------------------------------------------------------------
	// Gets the first element from the sequence. This panics if the sequence is empty.
	Head() Foo2

	// Gets the last element from the sequence. This panics if the sequence is empty.
	Last() Foo2

	// Gets the remainder after the first element from the sequence. This panics if the sequence is empty.
	Tail() Foo2Seq

	// Gets everything except the last element from the sequence. This panics if the sequence is empty.
	Init() Foo2Seq

	//-------------------------------------------------------------------------
	// Exists returns true if there exists at least one element in the sequence that matches
	// the predicate supplied.
	Exists(predicate func(Foo2) bool) bool

	// Forall returns true if every element in the sequence matches the predicate supplied.
	Forall(predicate func(Foo2) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Foo2))

	//-------------------------------------------------------------------------
	// Filter returns a new Foo2Seq whose elements return true for a predicate function.
	Filter(predicate func(Foo2) bool) (result Foo2Seq)

	// Partition returns two new Foo2Lists whose elements return true or false for the predicate, p.
	// The first result consists of all elements that satisfy the predicate and the second result consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original list.
	Partition(p func(Foo2) bool) (matching Foo2Seq, others Foo2Seq)

	//-------------------------------------------------------------------------
	// Find searches for the first value that matches a given predicate. It may or may not find one.
	Find(predicate func(Foo2) bool) OptionalFoo2

	// Converts the sequence to a list. For lists, this is merely a type assertion.
	ToList() Foo2List

	//-------------------------------------------------------------------------
	// Tests whether this sequence has the same length and the same elements as another sequence.
	// Omitted if Foo2 is not comparable.
	Equals(other Foo2Seq) bool

	// Contains tests whether a given value is present in the sequence.
	// Omitted if Foo2 is not comparable.
	Contains(value Foo2) bool

	// Count counts the number of times a given value occurs in the sequence.
	// Omitted if Foo2 is not comparable.
	Count(value Foo2) int

	// Distinct returns a new Foo2Seq whose elements are all unique.
	// Omitted if Foo2 is not comparable.
	Distinct() Foo2Seq

	//-------------------------------------------------------------------------
	// Sum sums Foo2 elements.
	// Omitted if Foo2 is not numeric.
	Sum() Foo2

	// Mean computes the arithmetic mean of all elements.
	// Panics if the list is empty.
	Mean() Foo2
}

//-------------------------------------------------------------------------------------------------
// OptionalFoo2 is an optional of type Foo2. Use it where you want to be explicit about
// the presence or absence of data.
//
// Optional values follow a similar pattern to Scala Options.
// See e.g. http://www.scala-lang.org/api/2.11.7/index.html#scala.Option

type OptionalFoo2 struct {
	x *Foo2
}

// shared none value
var noneFoo2 = OptionalFoo2{nil}

func NoFoo2() OptionalFoo2 {
	return noneFoo2
}

func SomeFoo2(x Foo2) OptionalFoo2 {

	return OptionalFoo2{&x}

}

//-------------------------------------------------------------------------------------------------

// panics if option is empty
func (o OptionalFoo2) Head() Foo2 {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return *(o.x)
}

// panics if option is empty
func (o OptionalFoo2) Last() Foo2 {
	return o.Head()
}

// panics if option is empty
func (o OptionalFoo2) Tail() Foo2Seq {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return noneFoo2
}

// panics if option is empty
func (o OptionalFoo2) Init() Foo2Seq {
	return o.Tail()
}

//-------------------------------------------------------------------------------------------------

func (o OptionalFoo2) Get() Foo2 {
	return o.Head()
}

func (o OptionalFoo2) GetOrElse(d func() Foo2) Foo2 {
	if o.IsEmpty() {
		return d()
	}
	return *o.x
}

func (o OptionalFoo2) OrElse(alternative func() OptionalFoo2) OptionalFoo2 {
	if o.IsEmpty() {
		return alternative()
	}
	return o
}

//-------------------------------------------------------------------------------------------------

func (o OptionalFoo2) Len() int {
	if o.IsEmpty() {
		return 0
	}
	return 1
}

func (o OptionalFoo2) IsEmpty() bool {
	return o.x == nil
}

func (o OptionalFoo2) NonEmpty() bool {
	return o.x != nil
}

// IsDefined returns true if the option is defined, i.e. non-empty. This is an alias for NonEmpty().
func (o OptionalFoo2) IsDefined() bool {
	return o.NonEmpty()
}

//-------------------------------------------------------------------------------------------------

func (o OptionalFoo2) Find(predicate func(Foo2) bool) OptionalFoo2 {
	if o.IsEmpty() {
		return o
	}
	if predicate(*o.x) {
		return o
	}
	return noneFoo2
}

func (o OptionalFoo2) Exists(predicate func(Foo2) bool) bool {
	if o.IsEmpty() {
		return false
	}
	return predicate(*o.x)
}

func (o OptionalFoo2) Forall(predicate func(Foo2) bool) bool {
	if o.IsEmpty() {
		return true
	}
	return predicate(*o.x)
}

func (o OptionalFoo2) Foreach(fn func(Foo2)) {
	if o.NonEmpty() {
		fn(*o.x)
	}
}

func (o OptionalFoo2) Filter(predicate func(Foo2) bool) Foo2Seq {
	return o.Find(predicate)
}

func (o OptionalFoo2) Partition(predicate func(Foo2) bool) (Foo2Seq, Foo2Seq) {
	if o.IsEmpty() {
		return o, o
	}
	if predicate(*o.x) {
		return o, noneFoo2
	}
	return noneFoo2, o
}

//-------------------------------------------------------------------------------------------------
// These methods require Foo2 be comparable.

// Equals verifies that one or more elements of Foo2List return true for the passed func.
func (o OptionalFoo2) Equals(other Foo2Seq) bool {
	if o.IsEmpty() {
		return other.IsEmpty()
	}
	if other.IsEmpty() || other.Len() > 1 {
		return false
	}
	a := o.Head()
	b := other.Head()
	return a == b
}

func (o OptionalFoo2) Contains(value Foo2) bool {
	if o.IsEmpty() {
		return false
	}
	return *(o.x) == value
}

func (o OptionalFoo2) Count(value Foo2) int {
	if o.Contains(value) {
		return 1
	}
	return 0
}

// Distinct returns a new Foo2Seq whose elements are all unique. For options, this simply returns the receiver.
// Omitted if Foo2 is not comparable.
func (o OptionalFoo2) Distinct() Foo2Seq {
	return o
}

//-------------------------------------------------------------------------------------------------
// Sum sums Foo2 elements.
// Omitted if Foo2 is not numeric.
func (o OptionalFoo2) Sum() Foo2 {

	if o.IsEmpty() {
		return 0
	}
	return *(o.x)

}

// Mean computes the arithmetic mean of all elements.
// Panics if the list is empty.
func (o OptionalFoo2) Mean() Foo2 {
	if o.IsEmpty() {
		panic("Cannot compute the arithmetic mean of zero-length OptionalFoo2")
	}
	return o.Sum()
}

func (o OptionalFoo2) ToList() Foo2List {
	if o.IsEmpty() {
		return Foo2List{}
	}
	return Foo2List{*o.x}
}

//-------------------------------------------------------------------------------------------------
// Foo2List is a slice of type Foo2. Use it where you would use []Foo2.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type Foo2List []Foo2

//-------------------------------------------------------------------------------------------------

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// panics if list is empty
func (list Foo2List) Head() Foo2 {
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// panics if list is empty
func (list Foo2List) Last() Foo2 {
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// panics if list is empty
func (list Foo2List) Tail() Foo2Seq {
	return Foo2List(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// panics if list is empty
func (list Foo2List) Init() Foo2Seq {
	return Foo2List(list[:len(list)-1])
}

// IsEmpty tests whether Foo2List is empty.
func (list Foo2List) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether Foo2List is empty.
func (list Foo2List) NonEmpty() bool {
	return len(list) > 0
}

// ToList simply returns the list in this case, but is part of the Seq interface.
func (list Foo2List) ToList() Foo2List {
	return list
}

// Len returns the number of items in the list.
// There is no Size() method; use Len() instead.
// This is one of the three methods in the standard sort.Interface.
func (list Foo2List) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is neceessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list Foo2List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// These methods require that Foo2 be ordered.

// Less determines whether one specified element is less than another specified element.
// This is one of the three methods in the standard sort.Interface.
func (list Foo2List) Less(i, j int) bool {
	return list[i] < list[j]
}

// Sort returns a new ordered Foo2List.
func (list Foo2List) Sort() Foo2List {
	result := make(Foo2List, len(list))
	copy(result, list)
	sort.Sort(result)
	return result
}

// IsSorted reports whether Foo2List is sorted.
func (list Foo2List) IsSorted() bool {
	return sort.IsSorted(list)
}

// SortDesc returns a new reverse-ordered Foo2List.
func (list Foo2List) SortDesc() Foo2List {
	result := make(Foo2List, len(list))
	copy(result, list)
	sort.Sort(sort.Reverse(result))
	return result
}

// IsSortedDesc reports whether Foo2List is reverse-sorted.
func (list Foo2List) IsSortedDesc() bool {
	return sort.IsSorted(sort.Reverse(list))
}

// Exists verifies that one or more elements of Foo2List return true for the passed func.
func (list Foo2List) Exists(fn func(Foo2) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of Foo2List return true for the passed func.
func (list Foo2List) Forall(fn func(Foo2) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over Foo2List and executes the passed func against each element.
func (list Foo2List) Foreach(fn func(Foo2)) {
	for _, v := range list {
		fn(v)
	}
}

// Reverse returns a copy of Foo2List with all elements in the reverse order.
func (list Foo2List) Reverse() Foo2List {
	numItems := len(list)
	result := make(Foo2List, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of Foo2List, using a version of the Fisher-Yates shuffle.
func (list Foo2List) Shuffle() Foo2List {
	numItems := len(list)
	result := make(Foo2List, numItems)
	copy(result, list)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.Swap(i, r)
	}
	return result
}

// Take returns a new Foo2List containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Foo2List) Take(n int) Foo2List {
	if n > len(list) {
		return list
	} else {
		return list[0:n]
	}
}

// Drop returns a new Foo2List without the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Foo2List) Drop(n int) Foo2List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[n:]
	}
}

// TakeLast returns a new Foo2List containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Foo2List) TakeLast(n int) Foo2List {
	l := len(list)
	if n > l {
		return list
	} else {
		return list[l-n:]
	}
}

// DropLast returns a new Foo2List without the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Foo2List) DropLast(n int) Foo2List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new Foo2List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list Foo2List) TakeWhile(p func(Foo2) bool) (result Foo2List) {
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new Foo2List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list Foo2List) DropWhile(p func(Foo2) bool) (result Foo2List) {
	adding := false
	for _, v := range list {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}
	return
}

// Filter returns a new Foo2List whose elements return true for func.
func (list Foo2List) Filter(fn func(Foo2) bool) Foo2Seq {
	result := make(Foo2List, 0, len(list)/2)
	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Partition returns two new Foo2Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list Foo2List) Partition(p func(Foo2) bool) (Foo2Seq, Foo2Seq) {
	matching := make(Foo2List, 0, len(list)/2)
	others := make(Foo2List, 0, len(list)/2)
	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of Foo2List that return true for the passed predicate.
func (list Foo2List) CountBy(predicate func(Foo2) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of Foo2List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Returns error if no elements.
func (list Foo2List) MinBy(less func(Foo2, Foo2) bool) (result Foo2, err error) {
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

// MaxBy returns an element of Foo2List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements.
func (list Foo2List) MaxBy(less func(Foo2, Foo2) bool) (result Foo2, err error) {
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

// DistinctBy returns a new Foo2List whose elements are unique, where equality is defined by a passed func.
func (list Foo2List) DistinctBy(equal func(Foo2, Foo2) bool) (result Foo2List) {
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

// IndexWhere finds the index of the first element satisfying some predicate. If none exists, -1 is returned.
func (list Foo2List) IndexWhere(p func(Foo2) bool) int {
	for i, v := range list {
		if p(v) {
			return i
		}
	}
	return -1
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list Foo2List) IndexWhere2(p func(Foo2) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list Foo2List) LastIndexWhere(p func(Foo2) bool) int {
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		if p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere2 finds the index of the last element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list Foo2List) LastIndexWhere2(p func(Foo2) bool, before int) int {
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}

// These methods require Foo2 be comparable.

// Equals verifies that one or more elements of Foo2List return true for the passed func.
func (list Foo2List) Equals(other Foo2Seq) bool {
	if len(list) != other.Len() {
		return false
	}
	eq := true
	i := 0
	other.Foreach(func(a Foo2) {
		if eq {
			v := list[i]
			if v != a {
				eq = false
			}
			i += 1
		}
	})
	return eq
}

// These methods require Foo2 be comparable.

// IndexOf finds the index of the first element specified. If none exists, -1 is returned.
func (list Foo2List) IndexOf(value Foo2) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return -1
}

// IndexOf2 finds the index of the first element specified at or after some start index.
// If none exists, -1 is returned.
func (list Foo2List) IndexOf2(value Foo2, from int) int {
	for i, v := range list {
		if i >= from && v == value {
			return i
		}
	}
	return -1
}

// Contains verifies that a given value is contained in Foo2List.
func (list Foo2List) Contains(value Foo2) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// Count gives the number elements of Foo2List that match a certain value.
func (list Foo2List) Count(value Foo2) (result int) {
	for _, v := range list {
		if v == value {
			result++
		}
	}
	return
}

// Distinct returns a new Foo2List whose elements are unique.
func (list Foo2List) Distinct() Foo2Seq {
	result := make(Foo2List, 0)
	appended := make(map[Foo2]bool)
	for _, v := range list {
		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}
	}
	return result
}

// These methods require Foo2 be numeric.

// Sum sums Foo2 elements in Foo2List.
func (list Foo2List) Sum() (result Foo2) {
	for _, v := range list {
		result += v
	}
	return
}

// Mean computes the arithmetic mean of all elements.
// Panics if the list is empty.
func (list Foo2List) Mean() Foo2 {
	l := len(list)
	if l == 0 {
		panic("Cannot compute the arithmetic mean of zero-length Foo2List")
	}
	return list.Sum() / Foo2(l)
}

// These methods require Foo2 be ordered.

// Min returns the minimum value of Foo2List. In the case of multiple items being equally minimal,
// the first such element is returned. Returns error if no elements.
func (list Foo2List) Min() (result Foo2, err error) {
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

// Max returns the maximum value of Foo2List. In the case of multiple items being equally maximal,
// the first such element is returned. Returns error if no elements.
func (list Foo2List) Max() (result Foo2, err error) {
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

// MkString concatenates the values as a string.
func (list Foo2List) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string.
func (list Foo2List) MkString3(pfx, mid, sfx string) string {
	b := bytes.Buffer{}
	b.WriteString(pfx)
	l := len(list)
	if l > 0 {
		v := list[0]
		b.WriteString(fmt.Sprintf("%v", v))
		for i := 1; i < l; i++ {
			v := list[i]
			b.WriteString(mid)
			b.WriteString(fmt.Sprintf("%v", v))
		}
	}
	b.WriteString(sfx)
	return b.String()
}

// optionForList

// First returns the first element that returns true for the passed func. Returns error if no elements return true.
func (list Foo2List) Find(fn func(Foo2) bool) OptionalFoo2 {
	for _, v := range list {
		if fn(v) {
			//return SomeFoo2(v)
		}
	}
	return NoFoo2()
}

// HeadOption gets the first item in the list, provided there is one.
func (list Foo2List) HeadOption() OptionalFoo2 {
	l := len(list)
	if l > 0 {
		return SomeFoo2(list[0])
	} else {
		return NoFoo2()
	}
}

// TailOption gets the last item in the list, provided there is one.
func (list Foo2List) LastOption() OptionalFoo2 {
	l := len(list)
	if l > 0 {
		return SomeFoo2(list[l-1])
	} else {
		return NoFoo2()
	}
}
