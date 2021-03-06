package option

import "github.com/rickb777/golist/internal/collection"

const Optional = collection.Collection + `
//-------------------------------------------------------------------------------------------------

// Optional{{.TName}} is an optional of type {{.PName}}. Use it where you want to be explicit about
// the presence or absence of data.
//
// Optional values follow a similar pattern to Scala Options. In particular, an option is a collection
// with a maximum cardinality of one. As such, options can be converted to/from lists and sets.
// See e.g. http://www.scala-lang.org/api/2.11.7/index.html#scala.Option
type Optional{{.TName}} struct {
	x *{{.TName}}
}

// shared none value
var none{{.TName}} = Optional{{.TName}}{ nil }

// No{{.TName}} gets an empty instance.
func No{{.TName}}() Optional{{.TName}} {
	return none{{.TName}}
}

// Some{{.TName}} gets a non-empty instance wrapping some value *x*.
func Some{{.TName}}(x {{.PName}}) Optional{{.TName}} {
	{{if .Type.Pointer}}
	if x == nil {
		return No{{.TName}}()
	}
	return Optional{{.TName}}{ x }
	{{else}}
	return Optional{{.TName}}{ &x }
	{{end}}
}

//-------------------------------------------------------------------------------------------------

// panics if option is empty
func (o Optional{{.TName}}) Head() {{.PName}} {
	return o.Get()
}

func (o Optional{{.TName}}) Get() {{.PName}} {
	if o.IsEmpty() {
		panic("Attempt to access non-existent value")
	}
	return {{.Deref}}o.x
}

func (o Optional{{.TName}}) GetOrElse(d func() {{.PName}}) {{.PName}} {
	if o.IsEmpty() {
		return d()
	}
	return {{.Deref}}o.x
}

func (o Optional{{.TName}}) OrElse(alternative func() Optional{{.TName}}) Optional{{.TName}} {
	if o.IsEmpty() {
		return alternative()
	}
	return o
}

//-------------------------------------------------------------------------------------------------

func (o Optional{{.TName}}) Size() int {
	if o.IsEmpty() {
		return 0
	}
	return 1
}

func (o Optional{{.TName}}) Len() int {
	return o.Size()
}

func (o Optional{{.TName}}) IsEmpty() bool {
	return o.x == nil
}

func (o Optional{{.TName}}) NonEmpty() bool {
	return o.x != nil
}

// IsSequence returns false for options.
func (o Optional{{.TName}}) IsSequence() bool {
	return false
}

// IsSet returns false for options.
func (o Optional{{.TName}}) IsSet() bool {
	return false
}

// IsDefined returns true if the option is defined, i.e. non-empty. This is an alias for NonEmpty().
func (o Optional{{.TName}}) IsDefined() bool {
	return o.NonEmpty()
}

//-------------------------------------------------------------------------------------------------

func (o Optional{{.TName}}) Find(predicate func({{.PName}}) bool) Optional{{.TName}} {
	if o.IsEmpty() {
		return o
	}
	if predicate({{.Deref}}o.x) {
		return o
	}
	return none{{.TName}}
}

func (o Optional{{.TName}}) Exists(predicate func({{.PName}}) bool) bool {
	if o.IsEmpty() {
		return false
	}
	return predicate({{.Deref}}o.x)
}

func (o Optional{{.TName}}) Forall(predicate func({{.PName}}) bool) bool {
	if o.IsEmpty() {
		return true
	}
	return predicate({{.Deref}}o.x)
}

func (o Optional{{.TName}}) Foreach(fn func({{.PName}})) {
	if o.NonEmpty() {
		fn({{.Deref}}o.x)
	}
}

// Send returns a channel that will send all the elements in order.
func (o Optional{{.TName}}) Send() <-chan {{.PName}} {
	ch := make(chan {{.PName}})
	go func() {
		if o.NonEmpty() {
			ch <- {{.Deref}}o.x
		}
		close(ch)
	}()
	return ch
}

func (o Optional{{.TName}}) Filter(predicate func({{.PName}}) bool) {{.TName}}Collection {
	return o.Find(predicate)
}

func (o Optional{{.TName}}) Partition(predicate func({{.PName}}) bool) ({{.TName}}Collection, {{.TName}}Collection) {
	if o.IsEmpty() {
		return o, o
	}
	if predicate({{.Deref}}o.x) {
		return o, none{{.TName}}
	}
	return none{{.TName}}, o
}

func (o Optional{{.TName}}) ToSlice() []{{.PName}} {
	slice := make([]{{.PName}}, o.Size())
	if o.NonEmpty() {
		slice[0] = {{.Deref}}o.x
	}
	return slice
}

{{if .Type.Underlying.IsBasic}}
// To{{.Type.Underlying.LongName}}s gets all the elements in a []{{.Type.Underlying}}.
func (o Optional{{.TName}}) To{{.Type.Underlying.LongName}}s() []{{.Type.Underlying}} {
	slice := make([]{{.Type.Underlying}}, o.Size())
	if o.NonEmpty() {
		slice[0] = {{.Type.Underlying}}({{.Deref}}o.x)
	}
	return slice
}

{{end}}
{{if .Has.List}}
// ToList gets the option's element in a {{.TName}}List.
func (o Optional{{.TName}}) ToList() {{.TName}}List {
	return {{.TName}}List(o.ToSlice())
}

{{end}}
{{if .Has.Set}}
// ToSet gets the option's element in a {{.TName}}Set.
func (o Optional{{.TName}}) ToSet() {{.TName}}Set {
	return New{{.TName}}Set(o.ToSlice()...)
}

{{end}}

` + comparable + ordered + numeric + mkstring
