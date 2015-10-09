package slice

import "github.com/clipperhouse/typewriter"

var groupByT = &typewriter.Template{
	Name: "GroupBy",
	Text: `
// GroupBy{{.TypeParameter.LongName}} groups elements into a map keyed by {{.TypeParameter}}. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv {{.ListName}}) GroupBy{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) map[{{.TypeParameter}}]{{.ListName}} {
	result := make(map[{{.TypeParameter}}]{{.ListName}})
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}
`,
	TypeParameterConstraints: []typewriter.Constraint{
		// exactly one type parameter is required, and it must be comparable
		{Comparable: true},
	},
}