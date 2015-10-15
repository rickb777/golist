// Generated by: setup
// TypeWriter: List
// Directive: +test on *Num2

package main

import (
	"errors"
	"math/rand"
)

// Num2Seq is an interface for sequences of type *Num2, including lists and options (where present).
type Num2Seq interface {
	// Len gets the size/length of the sequence.
	Len() int

	// IsEmpty returns true if the sequence is empty.
	IsEmpty() bool

	// NonEmpty returns true if the sequence is non-empty.
	NonEmpty() bool

	// Exists returns true if there exists at least one element in the sequence that matches
	// the predicate supplied.
	Exists(predicate func(*Num2) bool) bool

	// Forall returns true if every element in the sequence matches the predicate supplied.
	Forall(predicate func(*Num2) bool) bool

	// Foreach iterates over every element, executing a supplied function against each.
	Foreach(fn func(*Num2))

	// Filter returns a new Num2Seq whose elements return true for a predicate function.
	Filter(predicate func(*Num2) bool) (result Num2Seq)

	// Partition returns two new Num2Lists whose elements return true or false for the predicate, p.
	// The first result consists of all elements that satisfy the predicate and the second result consists of
	// all elements that don't. The relative order of the elements in the results is the same as in the
	// original list.
	Partition(p func(*Num2) bool) (matching Num2Seq, others Num2Seq)

	// Converts the sequence to a list. For lists, this is merely a type conversion.
	ToList() Num2List

	// Contains tests whether a given value is present in the sequence.
	// Omitted if Num2 is not comparable.
	Contains(value *Num2) bool

	// Count counts the number of times a given value occurs in the sequence.
	// Omitted if Num2 is not comparable.
	Count(value *Num2) int

	// Distinct returns a new Num2Seq whose elements are all unique.
	// Omitted if Num2 is not comparable.
	Distinct() Num2Seq
}

//-------------------------------------------------------------------------------------------------
// Num2List is a slice of type *Num2. Use it where you would use []*Num2.
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// Importantly, *none of its methods ever mutate a list*; they merely return new lists where required.
// When a list needs mutating, use normal Go slice operations, e.g. *append()*.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type Num2List []*Num2

//-------------------------------------------------------------------------------------------------

// Len returns the number of items in the list.
// There is no Size() method; use Len() instead.
// This is one of the three methods in the standard sort.Interface.
func (list Num2List) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is neceessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list Num2List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// IsEmpty tests whether Num2List is empty.
func (list Num2List) IsEmpty() bool {
	return len(list) == 0
}

// NonEmpty tests whether Num2List is empty.
func (list Num2List) NonEmpty() bool {
	return len(list) > 0
}

// ToList simply returns the list in this case, but is part of the Seq interface.
func (list Num2List) ToList() Num2List {
	return list
}

// Exists verifies that one or more elements of Num2List return true for the passed func.
func (list Num2List) Exists(fn func(*Num2) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of Num2List return true for the passed func.
func (list Num2List) Forall(fn func(*Num2) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over Num2List and executes the passed func against each element.
func (list Num2List) Foreach(fn func(*Num2)) {
	for _, v := range list {
		fn(v)
	}
}

// Reverse returns a copy of Num2List with all elements in the reverse order.
func (list Num2List) Reverse() Num2List {
	numItems := len(list)
	result := make(Num2List, numItems)
	last := numItems - 1
	for i, v := range list {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of Num2List, using a version of the Fisher-Yates shuffle.
func (list Num2List) Shuffle() Num2List {
	numItems := len(list)
	result := make(Num2List, numItems)
	copy(result, list)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result.Swap(i, r)
	}
	return result
}

// Take returns a new Num2List containing the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Num2List) Take(n int) Num2List {
	if n > len(list) {
		return list
	} else {
		return list[0:n]
	}
}

// Drop returns a new Num2List without the leading n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Num2List) Drop(n int) Num2List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[n:]
	}
}

// TakeLast returns a new Num2List containing the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Num2List) TakeLast(n int) Num2List {
	l := len(list)
	if n > l {
		return list
	} else {
		return list[l-n:]
	}
}

// DropLast returns a new Num2List without the trailing n elements of the source list.
// If n is greater than the size of the list, the whole list is returned.
func (list Num2List) DropLast(n int) Num2List {
	l := len(list)
	if n > l {
		return list[l:]
	} else {
		return list[0 : l-n]
	}
}

// TakeWhile returns a new Num2List containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elemense are excluded.
func (list Num2List) TakeWhile(p func(*Num2) bool) (result Num2List) {
	for _, v := range list {
		if p(v) {
			result = append(result, v)
		} else {
			return
		}
	}
	return
}

// DropWhile returns a new Num2List containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elemense are added.
func (list Num2List) DropWhile(p func(*Num2) bool) (result Num2List) {
	adding := false
	for _, v := range list {
		if !p(v) || adding {
			adding = true
			result = append(result, v)
		}
	}
	return
}

// Filter returns a new Num2List whose elements return true for func.
func (list Num2List) Filter(fn func(*Num2) bool) Num2Seq {
	result := make(Num2List, 0, len(list)/2)
	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Partition returns two new Num2Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
func (list Num2List) Partition(p func(*Num2) bool) (Num2Seq, Num2Seq) {
	matching := make(Num2List, 0, len(list)/2)
	others := make(Num2List, 0, len(list)/2)
	for _, v := range list {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return matching, others
}

// CountBy gives the number elements of Num2List that return true for the passed predicate.
func (list Num2List) CountBy(predicate func(*Num2) bool) (result int) {
	for _, v := range list {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of Num2List containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Returns error if no elements.
func (list Num2List) MinBy(less func(*Num2, *Num2) bool) (result *Num2, err error) {
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

// MaxBy returns an element of Num2List containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements.
func (list Num2List) MaxBy(less func(*Num2, *Num2) bool) (result *Num2, err error) {
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

// DistinctBy returns a new Num2List whose elements are unique, where equality is defined by a passed func.
func (list Num2List) DistinctBy(equal func(*Num2, *Num2) bool) (result Num2List) {
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

// These methods require *Num2 be comparable.

// Contains verifies that a given value is contained in Num2List.
func (list Num2List) Contains(value *Num2) bool {
	for _, v := range list {

		if *v == *value {
			return true
		}

	}
	return false
}

// Count gives the number elements of Num2List that match a certain value.
func (list Num2List) Count(value *Num2) (result int) {
	for _, v := range list {

		if *v == *value {
			result++
		}

	}
	return
}

// Distinct returns a new Num2List whose elements are unique.
func (list Num2List) Distinct() Num2Seq {
	result := make(Num2List, 0)
	appended := make(map[Num2]bool)
	for _, v := range list {

		if !appended[*v] {
			result = append(result, v)
			appended[*v] = true
		}

	}
	return result
}

// Min returns the first element of Num2List containing the minimum value, when compared to other elements
// using a specified comparator function defining ‘less’.
// Returns an error if the Num2List is empty.
func (list Num2List) Min(less func(*Num2, *Num2) bool) (result *Num2, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("Cannot determine the minimum of an empty list.")
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

// Max returns the first element of Num2List containing the maximum value, when compared to other elements
// using a specified comparator function defining ‘less’.
// Returns an error if the Num2List is empty.
func (list Num2List) Max(less func(*Num2, *Num2) bool) (result *Num2, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("Cannot determine the maximum of an empty list.")
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

// optionForList
