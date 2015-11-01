// Generated by: setup
// TypeWriter: List
// Directive: +test on Foo1

package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
)

//-------------------------------------------------------------------------------------------------
// Foo1Collection is an interface for collections of type Foo1, including sets, lists and options (where present).
type Foo1Collection interface {
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
	Head() Foo1

	//-------------------------------------------------------------------------
	// ToSlice returns a plain slice containing all the elements in the collection.
	// This is useful for bespoke iteration etc.
	// For sequences, the order is well defined.
	// For non-sequences (i.e. sets) the first time it is used, order of the elements is not well defined. But
	// the order is stable, which means it will give the same order each subsequent time it is used.
	ToSlice() []Foo1

	// ToList gets all the elements in a in List.
	ToList() Foo1List

	// ToSet gets all the elements in a in Set.
	ToSet() Foo1Set

	// Send sends all elements along a channel of type Foo1.
	// For sequences, the order is well defined.
	// For non-sequences (i.e. sets) the first time it is used, order of the elements is not well defined. But
	// the order is stable, which means it will give the same order each subsequent time it is used.
	Send() <-chan Foo1

	//-------------------------------------------------------------------------
	// Exists returns true if there exists at least one element in the collection that matches
	// the predicate supplied.
	Exists(predicate func(Foo1) bool) bool

	// Forall returns true if every element in the collection matches the predicate supplied.
	Forall(predicate func(Foo1) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Foo1))

	//-------------------------------------------------------------------------
	// Filter returns a new Foo1Collection whose elements return true for a predicate function.
	// The relative order of the elements in the result is the same as in the
	// original collection.
	Filter(predicate func(Foo1) bool) (result Foo1Collection)

	// Partition returns two new Foo1Collections whose elements return true or false for the predicate, p.
	// The first consists of all elements that satisfy the predicate and the second consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original collection.
	Partition(p func(Foo1) bool) (matching Foo1Collection, others Foo1Collection)

	//-------------------------------------------------------------------------

	// Equals verifies that another Foo1Collection has the same size and elements as this one. Also,
	// if the collection is a sequence, the order must be the same.
	// Omitted if Foo1 is not comparable.
	Equals(other Foo1Collection) bool

	// Contains tests whether a given value is present in the collection.
	// Omitted if Foo1 is not comparable.
	Contains(value Foo1) bool

	//-------------------------------------------------------------------------
	// Sum sums Foo1 elements.
	// Omitted if Foo1 is not numeric.
	Sum() Foo1

	// Mean computes the arithmetic mean of all elements. Panics if the collection is empty.
	// Omitted if Foo1 is not numeric.
	Mean() float64

	// Min returns the minimum value of Foo1List. In the case of multiple items being equally minimal,
	// the first such element is returned. Panics if the collection is empty.
	Min() Foo1

	// Max returns the maximum value of Foo1List. In the case of multiple items being equally maximal,
	// the first such element is returned. Panics if the collection is empty.
	Max() Foo1

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
// Foo1List is a slice of type Foo1. Use it where you would use []Foo1.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type Foo1List []Foo1

//-------------------------------------------------------------------------------------------------
// BuildFoo1ListFrom constructs a new Foo1List from a channel that supplies values
// until it is closed.
func BuildFoo1ListFrom(source <-chan Foo1) Foo1List {
	result := make(Foo1List, 0)
	for v := range source {
		result = append(result, v)
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// panics if list is empty
func (list Foo1List) Head() Foo1 {
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// panics if list is empty
func (list Foo1List) Last() Foo1 {
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// panics if list is empty
func (list Foo1List) Tail() Foo1Collection {
	return Foo1List(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// panics if list is empty
func (list Foo1List) Init() Foo1Collection {
	return Foo1List(list[:len(list)-1])
}

// IsEmpty tests whether Foo1List is empty.
func (list Foo1List) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether Foo1List is empty.
func (list Foo1List) NonEmpty() bool {
	return len(list) > 0
}

// IsSequence returns true for lists.
func (list Foo1List) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list Foo1List) IsSet() bool {
	return false
}

// ToSlice gets all the list's elements in a plain slice. This is simply a type conversion.
func (list Foo1List) ToSlice() []Foo1 {
	return []Foo1(list)
}

// ToList simply returns the list in this case, but is part of the Collection interface.
func (list Foo1List) ToList() Foo1List {
	return list
}

// ToSet gets all the list's elements in a Foo1Set.
func (list Foo1List) ToSet() Foo1Set {
	set := make(map[Foo1]struct{})
	for _, v := range list {
		set[v] = struct{}{}
	}
	return Foo1Set(set)
}

// Size returns the number of items in the list - an alias of Len().
func (list Foo1List) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
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

// Send gets a channel that will send all the elements in order.
func (list Foo1List) Send() <-chan Foo1 {
	ch := make(chan Foo1)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
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
func (list Foo1List) Filter(fn func(Foo1) bool) Foo1Collection {
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
func (list Foo1List) Partition(p func(Foo1) bool) (Foo1Collection, Foo1Collection) {
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
// element is returned. Panics if there are no elements.
func (list Foo1List) MinBy(less func(Foo1, Foo1) bool) (result Foo1) {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
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
// element is returned. Panics if there are no elements.
func (list Foo1List) MaxBy(less func(Foo1, Foo1) bool) (result Foo1) {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
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

// IndexWhere finds the index of the first element satisfying some predicate. If none exists, -1 is returned.
func (list Foo1List) IndexWhere(p func(Foo1) bool) int {
	for i, v := range list {
		if p(v) {
			return i
		}
	}
	return -1
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list Foo1List) IndexWhere2(p func(Foo1) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list Foo1List) LastIndexWhere(p func(Foo1) bool) int {
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
func (list Foo1List) LastIndexWhere2(p func(Foo1) bool, before int) int {
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}

// These methods require Foo1 be comparable.

// Equals verifies that one or more elements of Foo1List return true for the passed func.
func (list Foo1List) Equals(other Foo1Collection) bool {
	if len(list) != other.Size() {
		return false
	}
	eq := true
	i := 0
	other.Foreach(func(a Foo1) {
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

// These methods require Foo1 be comparable.

// IndexOf finds the index of the first element specified. If none exists, -1 is returned.
func (list Foo1List) IndexOf(value Foo1) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return -1
}

// IndexOf2 finds the index of the first element specified at or after some start index.
// If none exists, -1 is returned.
func (list Foo1List) IndexOf2(value Foo1, from int) int {
	for i, v := range list {
		if i >= from && v == value {
			return i
		}
	}
	return -1
}

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

// Distinct returns a new Foo1List whose elements are unique, retaining the original order.
func (list Foo1List) Distinct() Foo1Collection {
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

//-------------------------------------------------------------------------------------------------
// These methods require Foo1 be numeric.

// Sum sums all elements in the list.
func (list Foo1List) Sum() (result Foo1) {
	for _, v := range list {
		result += v
	}
	return
}

// Mean computes the arithmetic mean of all elements.
// Panics if the list is empty.
func (list Foo1List) Mean() float64 {
	l := len(list)
	if l == 0 {
		panic("Cannot compute the arithmetic mean of zero-length Foo1List")
	}
	sum := list.Sum()
	return float64(sum) / float64(l)
}

//-------------------------------------------------------------------------------------------------
// These methods require Foo1 be ordered.

// Min returns the element with the minimum value. In the case of multiple items being equally minimal,
// the first such element is returned. Panics if the collection is empty.
func (list Foo1List) Min() (result Foo1) {
	if len(list) == 0 {
		panic("Cannot determine the Min of an empty list.")
	}
	result = list[0]
	for _, v := range list {
		if v < result {
			result = v
		}
	}
	return
}

// Max returns the element with the maximum value. In the case of multiple items being equally maximal,
// the first such element is returned. Panics if the collection is empty.
func (list Foo1List) Max() (result Foo1) {
	if len(list) == 0 {
		panic("Cannot determine the Max of an empty list.")
	}
	result = list[0]
	for _, v := range list {
		if v > result {
			result = v
		}
	}
	return
}

// String implements the Stringer interface to render the list as a comma-separated array.
func (list Foo1List) String() string {
	return list.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string.
func (list Foo1List) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string.
func (list Foo1List) MkString3(pfx, mid, sfx string) string {
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

// First returns the first element that returns true for the passed func. Returns none if no elements return true.
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
func (list Foo1List) LastOption() OptionalFoo1 {
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

// panics if option is empty
func (o OptionalFoo1) Head() Foo1 {
	return o.Get()
}

func (o OptionalFoo1) Get() Foo1 {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return *(o.x)
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

//-------------------------------------------------------------------------------------------------

func (o OptionalFoo1) Size() int {
	if o.IsEmpty() {
		return 0
	}
	return 1
}

func (o OptionalFoo1) Len() int {
	return o.Size()
}

func (o OptionalFoo1) IsEmpty() bool {
	return o.x == nil
}

func (o OptionalFoo1) NonEmpty() bool {
	return o.x != nil
}

// IsSequence returns false for options.
func (o OptionalFoo1) IsSequence() bool {
	return false
}

// IsSet returns false for options.
func (o OptionalFoo1) IsSet() bool {
	return false
}

// IsDefined returns true if the option is defined, i.e. non-empty. This is an alias for NonEmpty().
func (o OptionalFoo1) IsDefined() bool {
	return o.NonEmpty()
}

//-------------------------------------------------------------------------------------------------

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

// Send gets a channel that will send all the elements in order.
func (o OptionalFoo1) Send() <-chan Foo1 {
	ch := make(chan Foo1)
	go func() {
		if o.NonEmpty() {
			ch <- *o.x
		}
		close(ch)
	}()
	return ch
}

func (o OptionalFoo1) Filter(predicate func(Foo1) bool) Foo1Collection {
	return o.Find(predicate)
}

func (o OptionalFoo1) Partition(predicate func(Foo1) bool) (Foo1Collection, Foo1Collection) {
	if o.IsEmpty() {
		return o, o
	}
	if predicate(*o.x) {
		return o, noneFoo1
	}
	return noneFoo1, o
}

func (o OptionalFoo1) ToSlice() []Foo1 {
	slice := make([]Foo1, o.Size())
	if o.NonEmpty() {
		slice[0] = *o.x
	}
	return slice
}

// ToList gets the option's element in a Foo1List.
func (o OptionalFoo1) ToList() Foo1List {
	return Foo1List(o.ToSlice())
}

// ToSet gets the option's element in a Foo1Set.
func (o OptionalFoo1) ToSet() Foo1Set {
	return NewFoo1Set(o.ToSlice()...)
}

//-------------------------------------------------------------------------------------------------
// These methods require Foo1 be comparable.

// Equals verifies that one or more elements of Foo1List return true for the passed func.
func (o OptionalFoo1) Equals(other Foo1Collection) bool {
	if o.IsEmpty() {
		return other.IsEmpty()
	}
	if other.IsEmpty() || other.Size() > 1 {
		return false
	}
	a := o.x
	s := other.ToSlice()
	b := s[0]
	return *a == b
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

// Distinct returns a new Foo1Collection whose elements are all unique. For options, this simply returns the
// receiver.
// Omitted if Foo1 is not comparable.
func (o OptionalFoo1) Distinct() Foo1Collection {
	return o
}

// Min returns the minimum value of Foo1List. In the case of multiple items being equally minimal,
// the first such element is returned. Panics if the collection is empty.
func (o OptionalFoo1) Min() Foo1 {
	return o.Get()
}

// Max returns the maximum value of Foo1List. In the case of multiple items being equally maximal,
// the first such element is returned. Panics if the collection is empty.
func (o OptionalFoo1) Max() Foo1 {
	return o.Get()
}

//-------------------------------------------------------------------------------------------------
// Sum sums Foo1 elements.
// Omitted if Foo1 is not numeric.
func (o OptionalFoo1) Sum() Foo1 {

	if o.IsEmpty() {
		return 0
	}
	return *(o.x)

}

// Mean computes the arithmetic mean of all elements.
// Panics if the list is empty.
func (o OptionalFoo1) Mean() float64 {
	if o.IsEmpty() {
		panic("Cannot compute the arithmetic mean of zero-length OptionalFoo1")
	}
	return float64(*(o.x))
}

//-------------------------------------------------------------------------------------------------
// String implements the Stringer interface to render the option as an array of one element.
func (o OptionalFoo1) String() string {
	return o.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string.
func (o OptionalFoo1) MkString(sep string) string {
	return o.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string.
func (o OptionalFoo1) MkString3(pfx, mid, sfx string) string {
	if o.IsEmpty() {
		return fmt.Sprintf("%s%s", pfx, sfx)
	}
	return fmt.Sprintf("%s%v%s", pfx, *(o.x), sfx)
}

//-------------------------------------------------------------------------------------------------
// Foo1Set is a typesafe set of Foo1 items. Instances are essentially immutable.
// The set-agebra functions Union, Intersection and Difference allow new variants to be constructed
// easily.
//
// The implementation is based on Go maps.

type Foo1Set map[Foo1]struct{}

//-------------------------------------------------------------------------------------------------
// NewFoo1Set constructs a new set containing the supplied values, if any.
func NewFoo1Set(e ...Foo1) Foo1Set {
	set := make(map[Foo1]struct{})
	for _, v := range e {
		set[v] = struct{}{}
	}
	return Foo1Set(set)
}

// BuildFoo1SetFrom constructs a new Foo1Set from a channel that supplies values
// until it is closed.
func BuildFoo1SetFrom(source <-chan Foo1) Foo1Set {
	set := make(map[Foo1]struct{})
	for v := range source {
		set[v] = struct{}{}
	}
	return Foo1Set(set)
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns false for sets.
func (set Foo1Set) IsSequence() bool {
	return false
}

// IsSet returns true for sets.
func (set Foo1Set) IsSet() bool {
	return true
}

func (set Foo1Set) Size() int {
	return len(set)
}

func (set Foo1Set) IsEmpty() bool {
	return len(set) == 0
}

func (set Foo1Set) NonEmpty() bool {
	return len(set) > 0
}

// Head gets an arbitrary element.
func (set Foo1Set) Head() Foo1 {
	for v := range set {
		return v
	}
	panic("Set is empty")
}

// ToSlice gets all the set's elements in a plain slice.
func (set Foo1Set) ToSlice() []Foo1 {
	slice := make([]Foo1, set.Size())
	i := 0
	for v := range set {
		slice[i] = v
		i++
	}
	return slice
}

// ToList gets all the set's elements in a in SetList.
func (set Foo1Set) ToList() Foo1List {
	return Foo1List(set.ToSlice())
}

// ToSet gets the current set, which requires no further conversion.
func (set Foo1Set) ToSet() Foo1Set {
	return set
}

// Contains tests whether an item is already in the Foo1Set.
func (set Foo1Set) Contains(i Foo1) bool {
	_, found := set[i]
	return found
}

// ContainsAll tests whether many items are all in the Foo1Set.
func (set Foo1Set) ContainsAll(i ...Foo1) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

func (set Foo1Set) actualSubset(other Foo1Set) bool {
	for item := range set {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// Equals determines if two sets are equal to each other.
// They are considered equal if both are the same size and both have the same items.
func (set Foo1Set) Equals(other Foo1Collection) bool {
	otherSet, isSet := other.(Foo1Set)
	return isSet && set.Size() == other.Size() && set.actualSubset(otherSet)
}

// IsSubset determines if every item in the other set is in this set.
func (set Foo1Set) IsSubset(other Foo1Set) bool {
	return set.Size() <= other.Size() && set.actualSubset(other)
}

// IsProperSubset determines if every item in the other set is in this set and this set is
// smaller than the other.
func (set Foo1Set) IsProperSubset(other Foo1Set) bool {
	return set.Size() < other.Size() && set.actualSubset(other)
}

// IsSuperset determines if every item of this set is in the other set.
func (set Foo1Set) IsSuperset(other Foo1Set) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set Foo1Set) Union(other Foo1Set) Foo1Set {
	union := NewFoo1Set()
	for item := range set {
		union[item] = struct{}{}
	}
	for item := range other {
		union[item] = struct{}{}
	}
	return union
}

// Intersection returns a new set with items that exist only in both sets.
func (set Foo1Set) Intersection(other Foo1Set) Foo1Set {
	intersection := NewFoo1Set()
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
func (set Foo1Set) Difference(other Foo1Set) Foo1Set {
	diffs := NewFoo1Set()
	for item := range set {
		if !other.Contains(item) {
			diffs[item] = struct{}{}
		}
	}
	return diffs
}

// Add creates a new set with elements added. This is similar to Union, but takes a slice of extra values.
// The receiver is not modified.
func (set Foo1Set) Add(others ...Foo1) Foo1Set {
	added := NewFoo1Set()
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
func (set Foo1Set) Remove(unwanted ...Foo1) Foo1Set {
	removed := NewFoo1Set()
	for item := range set {
		removed[item] = struct{}{}
	}
	for _, item := range unwanted {
		delete(removed, item)
	}
	return removed
}

// Exists verifies that one or more elements of Foo1Set return true for the passed func.
func (set Foo1Set) Exists(fn func(Foo1) bool) bool {
	for v := range set {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of Foo1Set return true for the passed func.
func (set Foo1Set) Forall(fn func(Foo1) bool) bool {
	for v := range set {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over Foo1Set and executes the passed func against each element.
// The order of the elements is not well defined but is probably repeatably stable until the set is changed.
func (set Foo1Set) Foreach(fn func(Foo1)) {
	for v := range set {
		fn(v)
	}
}

// Send sends all elements along a channel of type Foo1.
// The order of the elements is not well defined but is probably repeatably stable until the set is changed.
func (set Foo1Set) Send() <-chan Foo1 {
	ch := make(chan Foo1)
	go func() {
		for v := range set {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Filter returns a new Foo1Set whose elements return true for func.
func (set Foo1Set) Filter(fn func(Foo1) bool) Foo1Collection {
	result := make(map[Foo1]struct{})
	for v := range set {
		if fn(v) {
			result[v] = struct{}{}
		}
	}
	return Foo1Set(result)
}

// Partition returns two new Foo1Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original set.
func (set Foo1Set) Partition(p func(Foo1) bool) (Foo1Collection, Foo1Collection) {
	matching := make(map[Foo1]struct{})
	others := make(map[Foo1]struct{})
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return Foo1Set(matching), Foo1Set(others)
}

// CountBy gives the number elements of Foo1Set that return true for the passed predicate.
func (set Foo1Set) CountBy(predicate func(Foo1) bool) (result int) {
	for v := range set {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of Foo1Set containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set Foo1Set) MinBy(less func(Foo1, Foo1) bool) (result Foo1) {
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

// MaxBy returns an element of Foo1Set containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Panics if there are no elements.
func (set Foo1Set) MaxBy(less func(Foo1, Foo1) bool) (result Foo1) {
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

// These methods require Foo1 be numeric.

// Sum sums all elements in the set.
func (set Foo1Set) Sum() (result Foo1) {
	for v := range set {
		result += v
	}
	return
}

// Mean computes the arithmetic mean of all elements.
// Panics if the set is empty.
func (set Foo1Set) Mean() float64 {
	l := len(set)
	if l == 0 {
		panic("Cannot compute the arithmetic mean of zero-length Foo1Set")
	}
	sum := set.Sum()
	return float64(sum) / float64(l)
}

//-------------------------------------------------------------------------------------------------
// These methods require Foo1 be ordered.

// Min returns the element with the minimum value. In the case of multiple items being equally minimal,
// any such element is returned. Panics if the collection is empty.
func (set Foo1Set) Min() (result Foo1) {
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
func (set Foo1Set) Max() (result Foo1) {
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
func (set Foo1Set) String() string {
	return set.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string.
func (set Foo1Set) MkString(sep string) string {
	return set.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string.
func (set Foo1Set) MkString3(pfx, mid, sfx string) string {
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

// MapToFoo2 transforms Foo1List to Foo2List.
func (list Foo1List) MapToFoo2(fn func(Foo1) Foo2) Foo2Collection {
	result := make(Foo2List, 0, len(list))
	for _, v := range list {
		u := fn(v)
		result = append(result, u)
	}
	return result
}

// FlatMapToFoo2 transforms Foo1List to Foo2List, by repeatedly
// calling the supplied function and concatenating the results as a single flat list.
func (list Foo1List) FlatMapToFoo2(fn func(Foo1) Foo2Collection) Foo2Collection {
	result := make(Foo2List, 0, len(list))
	for _, v := range list {
		u := fn(v)
		if u.NonEmpty() {
			result = append(result, (u.ToList())...)
		}
	}
	return result
}

// List flags: {Collection:false Sequence:false List:true Option:true Set:true Tag:map[MapTo:true]}
