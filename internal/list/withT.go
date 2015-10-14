package list

const WithParamFunctions = `

// FoldLeft{{.TypeParameter.LongName}} applies a binary operator to a start value and all elements of this list, going left to right.
func (list {{.TName}}List) FoldLeft{{.TypeParameter.LongName}}(zero {{.TypeParameter}}, fn func({{.TypeParameter}}, {{.PName}}) {{.TypeParameter}}) {{.TypeParameter}} {
	sum := zero
	for _, v := range list {
		sum = fn(sum, v)
	}
	return sum
}

// FoldRight{{.TypeParameter.LongName}} applies a binary operator to a start value and all elements of this list, going right to left.
func (list {{.TName}}List) FoldRight{{.TypeParameter.LongName}}(zero {{.TypeParameter}}, fn func({{.TypeParameter}}, {{.PName}}) {{.TypeParameter}}) {{.TypeParameter}} {
	sum := zero
	for i := len(list) - 1; i >= 0; i-- {
		sum = fn(sum, list[i])
	}
	return sum
}

{{if .TypeParameter.Comparable}}
// This methods require {{.PName}} be comparable.

// GroupBy{{.TypeParameter.LongName}} groups elements into a map keyed by {{.TypeParameter}}.
func (list {{.TName}}List) GroupBy{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) map[{{.TypeParameter}}]{{.TName}}List {
	result := make(map[{{.TypeParameter}}]{{.TName}}List)
	for _, v := range list {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}
{{end}}

{{if .TypeParameter.Numeric}}
// These methods require {{.PName}} be numeric.

// Sum{{.TypeParameter.LongName}} sums {{.PName}} over elements in {{.TName}}List. See: http://clipperhouse.github.io/gen/#Sum
func (list {{.TName}}List) Sum{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) (result {{.TypeParameter}}) {
	for _, v := range list {
		result += fn(v)
	}
	return
}

// Mean{{.TypeParameter.LongName}} sums {{.TypeParameter}} over all elements and divides by len({{.TName}}List). See: http://clipperhouse.github.io/gen/#Mean
func (list {{.TName}}List) Mean{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) (result {{.TypeParameter}}, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("cannot determine Mean[{{.TypeParameter}}] of zero-length {{.TName}}List")
		return
	}
	for _, v := range list {
		result += fn(v)
	}
	result = result / {{.TypeParameter}}(l)
	return
}
{{end}}

{{if .TypeParameter.Ordered}}
// These methods require {{.PName}} be ordered.

// Min{{.TypeParameter.LongName}} selects the least value of {{.TypeParameter}} in {{.TName}}List.
// Returns error on {{.TName}}List with no elements.
func (list {{.TName}}List) Min{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) (result {{.TypeParameter}}, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("cannot determine Min of zero-length {{.TName}}List")
		return
	}
	result = fn(list[0])
	if l > 1 {
		for _, v := range list[1:] {
			f := fn(v)
			if f < result {
				result = f
			}
		}
	}
	return
}

// Max{{.TypeParameter.LongName}} selects the largest value of {{.TypeParameter}} in {{.TName}}List.
// Returns error on {{.TName}}List with no elements.
func (list {{.TName}}List) Max{{.TypeParameter.LongName}}(fn func({{.PName}}) {{.TypeParameter}}) (result {{.TypeParameter}}, err error) {
	l := len(list)
	if l == 0 {
		err = errors.New("cannot determine Max of zero-length {{.TName}}List")
		return
	}
	result = fn(list[0])
	if l > 1 {
		for _, v := range list[1:] {
			f := fn(v)
			if f > result {
				result = f
			}
		}
	}
	return
}
{{end}}
`
