package list

const sortable = `

// Len returns the number of items in the list.
// There is no Size() method; use Len() instead.
// This is one of the three methods in the standard sort.Interface.
func (list {{.TName}}List) Len() int {
	return len(list)
}

// Swap exchanges two elements, which is neceessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list {{.TName}}List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

{{if .Type.Ordered}}
// These methods require that {{.TName}} be ordered.

// Less determines whether one specified element is less than another specified element.
// This is one of the three methods in the standard sort.Interface.
func (list {{.TName}}List) Less(i, j int) bool {
	return list[i] < list[j]
}

// Sort returns a new ordered {{.TName}}List.
func (list {{.TName}}List) Sort() {{.TName}}List {
	result := make({{.TName}}List, len(list))
	copy(result, list)
	sort.Sort(result)
	return result
}

// IsSorted reports whether {{.TName}}List is sorted.
func (list {{.TName}}List) IsSorted() bool {
	return sort.IsSorted(list)
}

// SortDesc returns a new reverse-ordered {{.TName}}List.
func (list {{.TName}}List) SortDesc() {{.TName}}List {
	result := make({{.TName}}List, len(list))
	copy(result, list)
	sort.Sort(sort.Reverse(result))
	return result
}

// IsSortedDesc reports whether {{.TName}}List is reverse-sorted.
func (list {{.TName}}List) IsSortedDesc() bool {
	return sort.IsSorted(sort.Reverse(list))
}

{{end}}
`