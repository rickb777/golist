// Generated by: setup
// TypeWriter: List
// Directive: +test on Foo

package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
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

	// ToSlice returns a plain slice containing all the elements in the collection. This is useful for bespoke iteration etc.
	// For sequences, the order of the elements is simple and well defined.
	// For non-sequences (i.e. sets) the order of the elements is stable but not well defined. This means it will give
	// the same order each subsequent time it is used as it did the first time.
	ToSlice() []Foo

	// ToStrings gets all the elements in a slice of the underlying type, []string.
	ToStrings() []string

	// ToList gets all the elements in a List.
	ToList() FooList

	// ToSet gets all the elements in a Set.
	ToSet() FooSet

	// Send sends all elements along a channel of type Foo.
	// For sequences, the order of the elements is simple and well defined.
	// For non-sequences (i.e. sets) the order of the elements is stable but not well defined. This means it will give
	// the same order each subsequent time it is used as it did the first time.
	Send() <-chan Foo

	// Exists returns true if there exists at least one element in the collection that matches
	// the predicate supplied.
	Exists(predicate func(Foo) bool) bool

	// Forall returns true if every element in the collection matches the predicate supplied.
	Forall(predicate func(Foo) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(Foo))

	// Filter returns a new FooCollection whose elements return true for a predicate function.
	// The relative order of the elements in the result is the same as in the
	// original collection.
	Filter(predicate func(Foo) bool) (result FooCollection)

	// Partition returns two new FooCollections whose elements return true or false for the predicate, p.
	// The first consists of all elements that satisfy the predicate and the second consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original collection.
	Partition(p func(Foo) bool) (matching FooCollection, others FooCollection)

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

// FooList is a slice of type Foo. Use it where you would use []Foo.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type FooList []Foo

//-------------------------------------------------------------------------------------------------

// NewFooList constructs a new list containing the supplied values, if any.
func NewFooList(values ...Foo) FooList {
	list := make(FooList, len(values))
	for i, v := range values {
		list[i] = v
	}
	return list
}

// NewFooListFromStrings constructs a new FooList from a []string.
func NewFooListFromStrings(values []string) FooList {
	list := make(FooList, len(values))
	for i, v := range values {
		list[i] = Foo(v)
	}
	return list
}

// BuildFooListFromChan constructs a new FooList from a channel that supplies a sequence
// of values until it is closed. The function doesn't return until then.
func BuildFooListFromChan(source <-chan Foo) FooList {
	result := make(FooList, 0)
	for v := range source {
		result = append(result, v)
	}
	return result
}

//-------------------------------------------------------------------------------------------------

//-------------------------------------------------------------------------------------------------

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty
func (list FooList) Head() Foo {
	return list[0]
}

// Last gets the last element in the list. Init plus Last include the whole list. Last is the opposite of Head.
// Panics if list is empty
func (list FooList) Last() Foo {
	return list[len(list)-1]
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty
func (list FooList) Tail() FooCollection {
	return FooList(list[1:])
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty
func (list FooList) Init() FooCollection {
	return FooList(list[:len(list)-1])
}

// IsEmpty tests whether FooList is empty.
func (list FooList) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether FooList is empty.
func (list FooList) NonEmpty() bool {
	return len(list) > 0
}

// IsSequence returns true for lists.
func (list FooList) IsSequence() bool {
	return true
}

// IsSet returns false for lists.
func (list FooList) IsSet() bool {
	return false
}

//-------------------------------------------------------------------------------------------------

// ToSlice gets all the list's elements in a plain slice. This is simply a type conversion and is hardly needed
// for lists, because the underlying type can be used directly also.
// It is part of the FooCollection interface.
func (list FooList) ToSlice() []Foo {
	return []Foo(list)
}

// ToStrings gets all the elements in a []string.
func (list FooList) ToStrings() []string {
	slice := make([]string, len(list))
	for i, v := range list {
		slice[i] = string(v)
	}
	return slice
}

// ToList simply returns the list in this case, but is part of the Collection interface.
func (list FooList) ToList() FooList {
	return list
}

// ToSet gets all the list's elements in a FooSet.
func (list FooList) ToSet() FooSet {
	set := make(map[Foo]struct{})
	for _, v := range list {
		set[v] = struct{}{}
	}
	return FooSet(set)
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list FooList) Size() int {
	return len(list)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list FooList) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is neceessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list FooList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// These methods require that Foo be ordered.

// Less determines whether one specified element is less than another specified element.
// This is one of the three methods in the standard sort.Interface.
func (list FooList) Less(i, j int) bool {
	return list[i] < list[j]
}

// Sort returns a new ordered FooList.
func (list FooList) Sort() FooList {
	result := make(FooList, len(list))
	copy(result, list)
	sort.Sort(result)
	return result
}

// IsSorted reports whether FooList is sorted.
func (list FooList) IsSorted() bool {
	return sort.IsSorted(list)
}

// SortDesc returns a new reverse-ordered FooList.
func (list FooList) SortDesc() FooList {
	result := make(FooList, len(list))
	copy(result, list)
	sort.Sort(sort.Reverse(result))
	return result
}

// IsSortedDesc reports whether FooList is reverse-sorted.
func (list FooList) IsSortedDesc() bool {
	return sort.IsSorted(sort.Reverse(list))
}

//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of FooList return true for the passed func.
func (list FooList) Exists(fn func(Foo) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of FooList return true for the passed func.
func (list FooList) Forall(fn func(Foo) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over FooList and executes the passed func against each element.
func (list FooList) Foreach(fn func(Foo)) {
	for _, v := range list {
		fn(v)
	}
}

// Send gets a channel that will send all the elements in order.
func (list FooList) Send() <-chan Foo {
	ch := make(chan Foo)
	go func() {
		for _, v := range list {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Reverse returns a copy of FooList with all elements in the reverse order.
func (list FooList) Reverse() FooList {
	numItems := len(list)
	result := make(FooList, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of FooList, using a version of the Fisher-Yates shuffle.
func (list FooList) Shuffle() FooList {
	numItems := len(list)
	result := make(FooList, numItems)
	copy(result, list)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.Swap(i, r)
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// Take returns a new FooList containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list FooList) Take(n int) FooList {
	if n > len(list) {
		return list
	} else {
		return list[0:n]
	}
}

// Drop returns a new FooList without the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list FooList) Drop(n int) FooList {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[n:]
	}
}

// TakeLast returns a new FooList containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list FooList) TakeLast(n int) FooList {
	l := len(list)
	if n > l {
		return list
	} else {
		return list[l-n:]
	}
}

// DropLast returns a new FooList without the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list FooList) DropLast(n int) FooList {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new FooList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list FooList) TakeWhile(p func(Foo) bool) (result FooList) {
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new FooList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list FooList) DropWhile(p func(Foo) bool) (result FooList) {
	adding := false
	for _, v := range list {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}
	return
}

//-------------------------------------------------------------------------------------------------

// Filter returns a new FooList whose elements return true for func.
func (list FooList) Filter(fn func(Foo) bool) FooCollection {
	result := make(FooList, 0, len(list)/2)
	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Partition returns two new FooLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list FooList) Partition(p func(Foo) bool) (FooCollection, FooCollection) {
	matching := make(FooList, 0, len(list)/2)
	others := make(FooList, 0, len(list)/2)
	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of FooList that return true for the passed predicate.
func (list FooList) CountBy(predicate func(Foo) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of FooList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list FooList) MinBy(less func(Foo, Foo) bool) (result Foo) {
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

// MaxBy returns an element of FooList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list FooList) MaxBy(less func(Foo, Foo) bool) (result Foo) {
	l := len(list)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(list[m], list[i]) {
			m = i
		}
	}
	result = list[m]
	return
}

// DistinctBy returns a new FooList whose elements are unique, where equality is defined by a passed func.
func (list FooList) DistinctBy(equal func(Foo, Foo) bool) (result FooList) {
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
func (list FooList) IndexWhere(p func(Foo) bool) int {
	for i, v := range list {
		if p(v) {
			return i
		}
	}
	return -1
}

// IndexWhere2 finds the index of the first element satisfying some predicate at or after some start index.
// If none exists, -1 is returned.
func (list FooList) IndexWhere2(p func(Foo) bool, from int) int {
	for i, v := range list {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying some predicate.
// If none exists, -1 is returned.
func (list FooList) LastIndexWhere(p func(Foo) bool) int {
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
func (list FooList) LastIndexWhere2(p func(Foo) bool, before int) int {
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}

// Equals verifies that another FooCollection has the same size and elements as this one. Also,
// because this collection is a sequence, the order must be the same.
// Omitted if Foo is not comparable.
func (list FooList) Equals(other FooCollection) bool {
	if len(list) != other.Size() {
		return false
	}
	eq := true
	i := 0
	other.Foreach(func(a Foo) {
		if eq {
			v := list[i]
			if v != a {
				eq = false
			}
			i++
		}
	})
	return eq
}

//-------------------------------------------------------------------------------------------------
// These methods are provided because Foo is comparable.

// IndexOf finds the index of the first element specified. If none exists, -1 is returned.
func (list FooList) IndexOf(value Foo) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return -1
}

// IndexOf2 finds the index of the first element specified at or after some start index.
// If none exists, -1 is returned.
func (list FooList) IndexOf2(value Foo, from int) int {
	for i, v := range list {
		if i >= from && v == value {
			return i
		}
	}
	return -1
}

// Contains verifies that a given value is contained in FooList.
func (list FooList) Contains(value Foo) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// Count gives the number elements of FooList that match a certain value.
func (list FooList) Count(value Foo) (result int) {
	for _, v := range list {
		if v == value {
			result++
		}
	}
	return
}

// Distinct returns a new FooList whose elements are unique, retaining the original order.
func (list FooList) Distinct() FooCollection {
	result := make(FooList, 0)
	appended := make(map[Foo]bool)
	for _, v := range list {
		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}
	}
	return result
}

//-------------------------------------------------------------------------------------------------
// These methods are provided because Foo is ordered.

// Min returns the element with the minimum value. In the case of multiple items being equally minimal,
// the first such element is returned. Panics if the collection is empty.
func (list FooList) Min() (result Foo) {
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
func (list FooList) Max() (result Foo) {
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

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the list as a comma-separated string enclosed in square brackets.
func (list FooList) String() string {
	return list.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string using a supplied separator. No enclosing marks are added.
func (list FooList) MkString(sep string) string {
	return list.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string, using the prefix, separator and suffix supplied.
func (list FooList) MkString3(pfx, mid, sfx string) string {
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

//-------------------------------------------------------------------------------------------------
// Methods to interface lists with options.

// First returns the first element that returns true for the passed func. Returns none if no elements return true.
func (list FooList) Find(fn func(Foo) bool) OptionalFoo {
	for _, v := range list {
		if fn(v) {
			//return SomeFoo(v)
		}
	}
	return NoFoo()
}

// HeadOption gets the first item in the list, provided there is one.
func (list FooList) HeadOption() OptionalFoo {
	l := len(list)
	if l > 0 {
		return SomeFoo(list[0])
	} else {
		return NoFoo()
	}
}

// TailOption gets the last item in the list, provided there is one.
func (list FooList) LastOption() OptionalFoo {
	l := len(list)
	if l > 0 {
		return SomeFoo(list[l-1])
	} else {
		return NoFoo()
	}
}

//-------------------------------------------------------------------------------------------------
// List:MapTo[Foo]

// MapToFoo transforms FooList to FooList.
func (list FooList) MapToFoo(fn func(Foo) Foo) FooCollection {
	result := make(FooList, 0, len(list))
	for _, v := range list {
		u := fn(v)
		result = append(result, u)
	}
	return result
}

// FlatMapToFoo transforms FooList to FooList, by repeatedly
// calling the supplied function and concatenating the results as a single flat list.
func (list FooList) FlatMapToFoo(fn func(Foo) FooCollection) FooCollection {
	result := make(FooList, 0, len(list))
	for _, v := range list {
		u := fn(v)
		if u.NonEmpty() {
			result = append(result, (u.ToList())...)
		}
	}
	return result
}

//-------------------------------------------------------------------------------------------------

// OptionalFoo is an optional of type Foo. Use it where you want to be explicit about
// the presence or absence of data.
//
// Optional values follow a similar pattern to Scala Options. In particular, an option is a collection
// with a maximum cardinality of one. As such, options can be converted to/from lists and sets.
// See e.g. http://www.scala-lang.org/api/2.11.7/index.html#scala.Option
type OptionalFoo struct {
	x *Foo
}

// shared none value
var noneFoo = OptionalFoo{nil}

// NoFoo gets an empty instance.
func NoFoo() OptionalFoo {
	return noneFoo
}

// SomeFoo gets a non-empty instance wrapping some value *x*.
func SomeFoo(x Foo) OptionalFoo {

	return OptionalFoo{&x}

}

//-------------------------------------------------------------------------------------------------

// panics if option is empty
func (o OptionalFoo) Head() Foo {
	return o.Get()
}

func (o OptionalFoo) Get() Foo {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return *o.x
}

func (o OptionalFoo) GetOrElse(d func() Foo) Foo {
	if o.IsEmpty() {
		return d()
	}
	return *o.x
}

func (o OptionalFoo) OrElse(alternative func() OptionalFoo) OptionalFoo {
	if o.IsEmpty() {
		return alternative()
	}
	return o
}

//-------------------------------------------------------------------------------------------------

func (o OptionalFoo) Size() int {
	if o.IsEmpty() {
		return 0
	}
	return 1
}

func (o OptionalFoo) Len() int {
	return o.Size()
}

func (o OptionalFoo) IsEmpty() bool {
	return o.x == nil
}

func (o OptionalFoo) NonEmpty() bool {
	return o.x != nil
}

// IsSequence returns false for options.
func (o OptionalFoo) IsSequence() bool {
	return false
}

// IsSet returns false for options.
func (o OptionalFoo) IsSet() bool {
	return false
}

// IsDefined returns true if the option is defined, i.e. non-empty. This is an alias for NonEmpty().
func (o OptionalFoo) IsDefined() bool {
	return o.NonEmpty()
}

//-------------------------------------------------------------------------------------------------

func (o OptionalFoo) Find(predicate func(Foo) bool) OptionalFoo {
	if o.IsEmpty() {
		return o
	}
	if predicate(*o.x) {
		return o
	}
	return noneFoo
}

func (o OptionalFoo) Exists(predicate func(Foo) bool) bool {
	if o.IsEmpty() {
		return false
	}
	return predicate(*o.x)
}

func (o OptionalFoo) Forall(predicate func(Foo) bool) bool {
	if o.IsEmpty() {
		return true
	}
	return predicate(*o.x)
}

func (o OptionalFoo) Foreach(fn func(Foo)) {
	if o.NonEmpty() {
		fn(*o.x)
	}
}

// Send gets a channel that will send all the elements in order.
func (o OptionalFoo) Send() <-chan Foo {
	ch := make(chan Foo)
	go func() {
		if o.NonEmpty() {
			ch <- *o.x
		}
		close(ch)
	}()
	return ch
}

func (o OptionalFoo) Filter(predicate func(Foo) bool) FooCollection {
	return o.Find(predicate)
}

func (o OptionalFoo) Partition(predicate func(Foo) bool) (FooCollection, FooCollection) {
	if o.IsEmpty() {
		return o, o
	}
	if predicate(*o.x) {
		return o, noneFoo
	}
	return noneFoo, o
}

func (o OptionalFoo) ToSlice() []Foo {
	slice := make([]Foo, o.Size())
	if o.NonEmpty() {
		slice[0] = *o.x
	}
	return slice
}

// ToStrings gets all the elements in a []string.
func (o OptionalFoo) ToStrings() []string {
	slice := make([]string, o.Size())
	if o.NonEmpty() {
		slice[0] = string(*o.x)
	}
	return slice
}

// ToList gets the option's element in a FooList.
func (o OptionalFoo) ToList() FooList {
	return FooList(o.ToSlice())
}

// ToSet gets the option's element in a FooSet.
func (o OptionalFoo) ToSet() FooSet {
	return NewFooSet(o.ToSlice()...)
}

//-------------------------------------------------------------------------------------------------
// These methods require Foo be comparable.

// Equals verifies that one or more elements of FooList return true for the passed func.
func (o OptionalFoo) Equals(other FooCollection) bool {
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

func (o OptionalFoo) Contains(value Foo) bool {
	if o.IsEmpty() {
		return false
	}
	return *(o.x) == value
}

func (o OptionalFoo) Count(value Foo) int {
	if o.Contains(value) {
		return 1
	}
	return 0
}

// Distinct returns a new FooCollection whose elements are all unique. For options, this simply returns the
// receiver.
// Omitted if Foo is not comparable.
func (o OptionalFoo) Distinct() FooCollection {
	return o
}

// Min returns the minimum value of FooList. In the case of multiple items being equally minimal,
// the first such element is returned. Panics if the collection is empty.
func (o OptionalFoo) Min() Foo {
	return o.Get()
}

// Max returns the maximum value of FooList. In the case of multiple items being equally maximal,
// the first such element is returned. Panics if the collection is empty.
func (o OptionalFoo) Max() Foo {
	return o.Get()
}

//-------------------------------------------------------------------------------------------------

// String implements the Stringer interface to render the option as an array of one element.
func (o OptionalFoo) String() string {
	return o.MkString3("[", ",", "]")
}

// MkString concatenates the values as a string.
func (o OptionalFoo) MkString(sep string) string {
	return o.MkString3("", sep, "")
}

// MkString3 concatenates the values as a string.
func (o OptionalFoo) MkString3(pfx, mid, sfx string) string {
	if o.IsEmpty() {
		return fmt.Sprintf("%s%s", pfx, sfx)
	}
	return fmt.Sprintf("%s%v%s", pfx, *(o.x), sfx)
}

//-------------------------------------------------------------------------------------------------

// FooSet is a typesafe set of Foo items.
// The implementation is based on an underyling Go map, which can be manipulated directly,
// but otherwise instances are essentially immutable.
// The set-agebra functions *Union*, *Intersection* and *Difference* allow new variants to be constructed
// easily; these methods do not modify the input sets.
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

// BuildFooSetFrom constructs a new FooSet from a channel that supplies a stream of values
// until it is closed. The function returns all these values in a set (i.e. without any duplicates).
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

// ToList gets all the set's elements in a in SetList.
func (set FooSet) ToList() FooList {
	return FooList(set.ToSlice())
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

// List flags: {Collection:false List:true Option:true Set:true Plumbing:false Tag:map[MapTo:true]}
