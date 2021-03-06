package set

const predicatedFunctions = `
// Filter returns a new {{.TName}}Set whose elements return true for func.
func (set {{.TName}}Set) Filter(fn func({{.PName}}) bool) {{.TName}}Collection {
	result := make(map[{{.PName}}]struct{})
	for v := range set {
		if fn(v) {
			result[v] = struct{}{}
		}
	}
	return {{.TName}}Set(result)
}

// Partition returns two new {{.TName}}Lists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original set.
func (set {{.TName}}Set) Partition(p func({{.PName}}) bool) ({{.TName}}Collection, {{.TName}}Collection) {
	matching := make(map[{{.PName}}]struct{})
	others := make(map[{{.PName}}]struct{})
	for v := range set {
		if p(v) {
			matching[v] = struct{}{}
		} else {
			others[v] = struct{}{}
		}
	}
	return {{.TName}}Set(matching), {{.TName}}Set(others)
}

// CountBy gives the number elements of {{.TName}}Set that return true for the passed predicate.
func (set {{.TName}}Set) CountBy(predicate func({{.PName}}) bool) (result int) {
	for v := range set {
		if predicate(v) {
			result++
		}
	}
	return
}

// MinBy returns an element of {{.TName}}Set containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (set {{.TName}}Set) MinBy(less func({{.PName}}, {{.PName}}) bool) (result {{.PName}}) {
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

// MaxBy returns an element of {{.TName}}Set containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such
// element is returned. Panics if there are no elements.
func (set {{.TName}}Set) MaxBy(less func({{.PName}}, {{.PName}}) bool) (result {{.PName}}) {
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

`
