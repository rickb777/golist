// Generated by: setup
// TypeWriter: slice
// Directive: +test on Thing

package main

import (
	"errors"
	"math/rand"
)

// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.

// ThingSlice is a slice of type Thing. Use it where you would use []Thing.
type ThingSlice []Thing

// Len returns the number of items in the slice.
// There is no Size() method; use Len() instead.
func (rcv ThingSlice) Len() int {
	return len(rcv)
}

// IsEmpty tests whether ThingSlice is empty.
func (slice ThingSlice) IsEmpty() bool {
	return len(slice) == 0
}

// NonEmpty tests whether ThingSlice is empty.
func (slice ThingSlice) NonEmpty() bool {
	return len(slice) > 0
}

// Exists verifies that one or more elements of ThingSlice return true for the passed func.
func (slice ThingSlice) Exists(fn func(Thing) bool) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of ThingSlice return true for the passed func.
func (slice ThingSlice) Forall(fn func(Thing) bool) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over ThingSlice and executes the passed func against each element.
func (slice ThingSlice) Foreach(fn func(Thing)) {
	for _, v := range slice {
		fn(v)
	}
}

// Filter returns a new ThingSlice whose elements return true for func.
func (rcv ThingSlice) Filter(fn func(Thing) bool) (result ThingSlice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Partition returns two new ThingSlices whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original slice.
func (slice ThingSlice) Partition(p func(Thing) bool) (matching ThingSlice, others ThingSlice) {
	for _, v := range slice {
		if p(v) {
			matching = append(matching, v)
		} else {
			others = append(others, v)
		}
	}
	return
}

// Reverse returns a copy of ThingSlice with all elements in the reverse order.
func (rcv ThingSlice) Reverse() ThingSlice {
	numItems := len(rcv)
	result := make(ThingSlice, numItems)
	last := numItems - 1
	for i, v := range rcv {
		result[last-i] = v
	}
	return result
}

// Shuffle returns a shuffled copy of ThingSlice, using a version of the Fisher-Yates shuffle. See: http://clipperhouse.github.io/gen/#Shuffle
func (rcv ThingSlice) Shuffle() ThingSlice {
	numItems := len(rcv)
	result := make(ThingSlice, numItems)
	copy(result, rcv)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result[r], result[i] = result[i], result[r]
	}
	return result
}

// CountBy gives the number elements of ThingSlice that return true for the passed predicate.
func (rcv ThingSlice) CountBy(predicate func(Thing) bool) (result int) {
	for _, v := range rcv {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of ThingSlice containing the minimum value, when compared to other elements using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MinBy
func (rcv ThingSlice) MinBy(less func(Thing, Thing) bool) (result Thing, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the Min of an empty slice")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(rcv[i], rcv[m]) {
			m = i
		}
	}
	result = rcv[m]
	return
}

// MaxBy returns an element of ThingSlice containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MaxBy
func (rcv ThingSlice) MaxBy(less func(Thing, Thing) bool) (result Thing, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the MaxBy of an empty slice")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if rcv[i] != rcv[m] && !less(rcv[i], rcv[m]) {
			m = i
		}
	}
	result = rcv[m]
	return
}

// DistinctBy returns a new ThingSlice whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
func (rcv ThingSlice) DistinctBy(equal func(Thing, Thing) bool) (result ThingSlice) {
Outer:
	for _, v := range rcv {
		for _, r := range result {
			if equal(v, r) {
				continue Outer
			}
		}
		result = append(result, v)
	}
	return result
}

// Contains verifies that a given value is contained in ThingSlice.
func (slice ThingSlice) Contains(value Thing) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Count gives the number elements of ThingSlice that match a certain value.
func (rcv ThingSlice) Count(value Thing) (result int) {
	for _, v := range rcv {
		if v == value {
			result++
		}
	}
	return
}

// Distinct returns a new ThingSlice whose elements are unique. See: http://clipperhouse.github.io/gen/#Distinct
func (rcv ThingSlice) Distinct() (result ThingSlice) {
	appended := make(map[Thing]bool)
	for _, v := range rcv {
		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}
	}
	return result
}

// GroupByOther groups elements into a map keyed by Other. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv ThingSlice) GroupByOther(fn func(Thing) Other) map[Other]ThingSlice {
	result := make(map[Other]ThingSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// MinOther selects the least value of Other in ThingSlice.
// Returns error on ThingSlice with no elements. See: http://clipperhouse.github.io/gen/#MinCustom
func (rcv ThingSlice) MinOther(fn func(Thing) Other) (result Other, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine Min of zero-length ThingSlice")
		return
	}
	result = fn(rcv[0])
	if l > 1 {
		for _, v := range rcv[1:] {
			f := fn(v)
			if f < result {
				result = f
			}
		}
	}
	return
}

// MaxOther selects the largest value of Other in ThingSlice.
// Returns error on ThingSlice with no elements. See: http://clipperhouse.github.io/gen/#MaxCustom
func (rcv ThingSlice) MaxOther(fn func(Thing) Other) (result Other, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine Max of zero-length ThingSlice")
		return
	}
	result = fn(rcv[0])
	if l > 1 {
		for _, v := range rcv[1:] {
			f := fn(v)
			if f > result {
				result = f
			}
		}
	}
	return
}

// SumOther sums Thing over elements in ThingSlice. See: http://clipperhouse.github.io/gen/#Sum
func (rcv ThingSlice) SumOther(fn func(Thing) Other) (result Other) {
	for _, v := range rcv {
		result += fn(v)
	}
	return
}

// MeanOther sums Other over all elements and divides by len(ThingSlice). See: http://clipperhouse.github.io/gen/#Mean
func (rcv ThingSlice) MeanOther(fn func(Thing) Other) (result Other, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine Mean[Other] of zero-length ThingSlice")
		return
	}
	for _, v := range rcv {
		result += fn(v)
	}
	result = result / Other(l)
	return
}

// AggregateOther iterates over ThingSlice, operating on each element while maintaining ‘state’. See: http://clipperhouse.github.io/gen/#Aggregate
func (rcv ThingSlice) AggregateOther(fn func(Other, Thing) Other) (result Other) {
	for _, v := range rcv {
		result = fn(result, v)
	}
	return
}

// MapToOther transforms a slice of Other from ThingSlice.
func (rcv ThingSlice) MapToOther(fn func(Thing) Other) (result OtherSlice) {
	for _, v := range rcv {
		result = append(result, fn(v))
	}
	return
}

// First returns the first element that returns true for the passed func. Returns error if no elements return true. See: http://clipperhouse.github.io/gen/#First
func (rcv ThingSlice) First(fn func(Thing) bool) (result Thing, err error) {
	for _, v := range rcv {
		if fn(v) {
			result = v
			return
		}
	}
	err = errors.New("no ThingSlice elements return true for passed func")
	return
}

// SortWith returns a new ordered ThingSlice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv ThingSlice) SortWith(less func(Thing, Thing) bool) ThingSlice {
	result := make(ThingSlice, len(rcv))
	copy(result, rcv)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSortThingSlice(result, less, 0, n, maxDepth)
	return result
}

// IsSortedWith reports whether an instance of ThingSlice is sorted, using the pass func to define ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv ThingSlice) IsSortedWith(less func(Thing, Thing) bool) bool {
	n := len(rcv)
	for i := n - 1; i > 0; i-- {
		if less(rcv[i], rcv[i-1]) {
			return false
		}
	}
	return true
}

// SortByDesc returns a new, descending-ordered ThingSlice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv ThingSlice) SortByDesc(less func(Thing, Thing) bool) ThingSlice {
	greater := func(a, b Thing) bool {
		return less(b, a)
	}
	return rcv.SortWith(greater)
}

// IsSortedDesc reports whether an instance of ThingSlice is sorted in descending order, using the pass func to define ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv ThingSlice) IsSortedByDesc(less func(Thing, Thing) bool) bool {
	greater := func(a, b Thing) bool {
		return less(b, a)
	}
	return rcv.IsSortedWith(greater)
}

// Sort implementation based on http://golang.org/pkg/sort/#Sort, see top of this file

func swapThingSlice(rcv ThingSlice, a, b int) {
	rcv[a], rcv[b] = rcv[b], rcv[a]
}

// Insertion sort
func insertionSortThingSlice(rcv ThingSlice, less func(Thing, Thing) bool, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(rcv[j], rcv[j-1]); j-- {
			swapThingSlice(rcv, j, j-1)
		}
	}
}

// siftDown implements the heap property on rcv[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDownThingSlice(rcv ThingSlice, less func(Thing, Thing) bool, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(rcv[first+child], rcv[first+child+1]) {
			child++
		}
		if !less(rcv[first+root], rcv[first+child]) {
			return
		}
		swapThingSlice(rcv, first+root, first+child)
		root = child
	}
}

func heapSortThingSlice(rcv ThingSlice, less func(Thing, Thing) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDownThingSlice(rcv, less, i, hi, first)
	}

	// Pop elements, largest first, into end of rcv.
	for i := hi - 1; i >= 0; i-- {
		swapThingSlice(rcv, first, first+i)
		siftDownThingSlice(rcv, less, lo, i, first)
	}
}

// Quicksort, following Bentley and McIlroy,
// Engineering a Sort Function, SP&E November 1993.

// medianOfThree moves the median of the three values rcv[a], rcv[b], rcv[c] into rcv[a].
func medianOfThreeThingSlice(rcv ThingSlice, less func(Thing, Thing) bool, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if less(rcv[m1], rcv[m0]) {
		swapThingSlice(rcv, m1, m0)
	}
	if less(rcv[m2], rcv[m1]) {
		swapThingSlice(rcv, m2, m1)
	}
	if less(rcv[m1], rcv[m0]) {
		swapThingSlice(rcv, m1, m0)
	}
	// now rcv[m0] <= rcv[m1] <= rcv[m2]
}

func swapRangeThingSlice(rcv ThingSlice, a, b, n int) {
	for i := 0; i < n; i++ {
		swapThingSlice(rcv, a+i, b+i)
	}
}

func doPivotThingSlice(rcv ThingSlice, less func(Thing, Thing) bool, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's Ninther, median of three medians of three.
		s := (hi - lo) / 8
		medianOfThreeThingSlice(rcv, less, lo, lo+s, lo+2*s)
		medianOfThreeThingSlice(rcv, less, m, m-s, m+s)
		medianOfThreeThingSlice(rcv, less, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThreeThingSlice(rcv, less, lo, m, hi-1)

	// Invariants are:
	//	rcv[lo] = pivot (set up by ChoosePivot)
	//	rcv[lo <= i < a] = pivot
	//	rcv[a <= i < b] < pivot
	//	rcv[b <= i < c] is unexamined
	//	rcv[c <= i < d] > pivot
	//	rcv[d <= i < hi] = pivot
	//
	// Once b meets c, can swap the "= pivot" sections
	// into the middle of the slice.
	pivot := lo
	a, b, c, d := lo+1, lo+1, hi, hi
	for {
		for b < c {
			if less(rcv[b], rcv[pivot]) { // rcv[b] < pivot
				b++
			} else if !less(rcv[pivot], rcv[b]) { // rcv[b] = pivot
				swapThingSlice(rcv, a, b)
				a++
				b++
			} else {
				break
			}
		}
		for b < c {
			if less(rcv[pivot], rcv[c-1]) { // rcv[c-1] > pivot
				c--
			} else if !less(rcv[c-1], rcv[pivot]) { // rcv[c-1] = pivot
				swapThingSlice(rcv, c-1, d-1)
				c--
				d--
			} else {
				break
			}
		}
		if b >= c {
			break
		}
		// rcv[b] > pivot; rcv[c-1] < pivot
		swapThingSlice(rcv, b, c-1)
		b++
		c--
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := min(b-a, a-lo)
	swapRangeThingSlice(rcv, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRangeThingSlice(rcv, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSortThingSlice(rcv ThingSlice, less func(Thing, Thing) bool, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSortThingSlice(rcv, less, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivotThingSlice(rcv, less, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSortThingSlice(rcv, less, a, mlo, maxDepth)
			a = mhi // i.e., quickSortThingSlice(rcv, mhi, b)
		} else {
			quickSortThingSlice(rcv, less, mhi, b, maxDepth)
			b = mlo // i.e., quickSortThingSlice(rcv, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSortThingSlice(rcv, less, a, b)
	}
}
