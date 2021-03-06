package golist

import (
	"github.com/rickb777/typewriter"
	"github.com/rickb777/golist/internal/set"
)

var coreSetTemplate = &typewriter.Template{
	Name: setName,
	Text: set.Set,
}

var otherSetTemplates = typewriter.TemplateSlice{
	setMapToT,
	setWithT,
	coreListTemplate,
	coreOptionTemplate,
	corePlumbingTemplate,
}

var setMapToT = &typewriter.Template{
	Name: "MapTo",
	Text: set.SetMapToParamFunctions,
	// exactly one type parameter is required, but no constraints on that type
	TypeParameterConstraints: []typewriter.Constraint{{}},
}

var setWithT = &typewriter.Template{
	Name: "With",
	Text: set.WithParamFunctions,
	// exactly one type parameter is required, but no constraints on that type
	TypeParameterConstraints: []typewriter.Constraint{{}},
}

