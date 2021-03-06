# GoList - Lists

## The List typewriter

The **List** typewriter generates a type and corresponding methods that follow a pattern inspired by Scala's List and
its parent trait [IndexedSeq](http://www.scala-lang.org/api/2.11.7/#scala.collection.IndexedSeq).

Although Go slices are mutable, the generated code *does not provide* any mutation methods, by design. Some functions
make a copy of the list, mutate the copy and return it without changing the original. There is one exception:
A mutating `Swap` method is needed for efficient list sorting; however, the main sort methods operate on *copies*
of the list.

You can manipulate the underlying slice using normal Go syntax when it needs to be changed, of course.

The **List** typewriter also offers easier ad-hoc sorts.

At it simplest, the annotation looks like:

````go
// +gen List
type Example struct { ... }
````

#### Core methods

The generated code is a list to hold `Example` values. It provides methods including:

 * **IsEmpty**, **NonEmpty** - get simple properties
 * **Size**, **Len** - get simple properties (these are aliases)
 * **Exists**, **Forall** - test whether any or all elements match some specified condition
 * **Foreach** - apply a function to every element in turn, typically causing side-effects
 * **Filter**, **Partition** - get a subset, or two disjoint subsets, of the collection
 * **CountBy**, **MaxBy**, **MinBy**, **DistinctBy** - statistics based on supplied operator functions
 * **MkString**, **MkString3**, **String** - constructs a string representation of the collection
 * **Send** - get a channel that supplies values in sequence

#### List Methods

 * **Head**, **Tail** - get the first element and the rest
 * **Init**, **Last** - get the last element and the rest
 * **Reverse**, **Shuffle** - get a new list that is reversed or shuffled
 * **Take**, **TakeWhile**, **DropLast** - get a new list without some trailing elements
 * **Drop**, **DropWhile**, **TakeLast** - get a new list without some leading elements
 * **IndexWhere**, **IndexWhere2** - find the index of the first match using a predicate function
 * **LastIndexWhere**, **LastIndexWhere2** - find the index of the last match using a predicate function

#### Comparable Methods

If the element type is *comparable*, it adds:

 * **Equals** - compare with another collection
 * **Contains** - compare with a specified value
 * **Count** - enumerates a specified value
 * **Distinct** - remove duplicates
 * **IndexOf**, **IndexOf2** - find the index of the first match
 * **LastIndexOf**, **LastIndexOf2** - find the index of the last match

#### Numeric Methods

If the element type is *numeric*, it adds:

 * **Sum** - sum all elements
 * **Mean** - compute the arithmetic mean

#### Min and Max

It always adds:

 * **Min**, **Max** - find the minimum/maximum value

but the implementation depends on whether the element type is *ordered* or not. For ordered elements, Min and Max use
simple inequality operators '<' and '>'. Otherwise a comparison function must be supplied.

#### Sorting

Also in the case of *ordered* elements, the list implements sorting using the [standard Go api](https://golang.org/pkg/sort/)
within these methods:

* **Sort**, **IsSorted** - sort in ascending order
* **SortDesc**, **IsSortedDesc** - sort in descending order

#### Optional Extremae

Finally, if a companion option is present, it adds:

 * **HeadOption** - get the first element if present (similar to Head)
 * **LastOption** - get the last element if present (similar to Last)

### List For Pointer Elements

A variant is to include a star:

````go
// +gen * List
type Example struct { ... }
````

This creates a basic list of `*Example` pointers, i.e. `[]*Example`.

### Tags

Extra tags can be included to add more features. You can include a comma-separated list of as many tags as you need.

#### `Option` and `Set` Tags

`Option` adds option implementation on the same type (e.g. `Example`). Some additional methods are provided to convert
between options and lists.

`Set` adds set implementation on the same type (e.g. `Example`).  Some additional methods are provided to convert
between options and sets.

#### `MapTo` Tag

`MapTo[T]` adds code to transform the original list to a new list by transforming each element using a function you provide.
`MapTo[T]` can be used more than once: 

````go
// +gen List:"MapTo[Fred], MapTo[Jim]"
type Example struct { ... }
````

This adds methods according to the parameter type `T` (e.g. `Fred` and `Jim` above):

 * **MapTo**T - converts to type `T`
 * **FlatMapTo**T - converts to collections of `T`

For `MapToT`, a mapping function transforms an individual `Example` item to its `Fred` equivalent.

For `FlatMapToT`, a mapping function transforms an individual `Example` item to a
sequence of its `T` equivalent. This sequence is added to the result set.

Normally the result is a `TCollection`. 
But if the target type is a basic type, such as `string` or `int`, the result is a slice of the basic type instead.

#### `With[T]` Tag

`With[T]` adds methods that need functions depending on another type, `T`.

````go
// +gen List:"With[Colour], With[Texture]"
type Example struct { ... }
````

This adds methods according to the type `T`:

 * **FoldLeft**T, **FoldRight**T (always) - provide general summation functions
 * **GroupBy**T (only if `T` is *comparable*) - clustering and histogramming
 * **Sum**T, **Mean**T (only if `T` is *numeric*)
 * **MinBy**T, **MaxBy**T (only if `T` is *ordered*) - find the min/max list entry according to a projection function

#### `SortWith` Tag

This adds extra sort methods (as alternatives) that simply depend on providing a comparator function. The benefit is
simpler usage at the calling sites, but much more code is generated to implement the sorting, so there is a trade-off
you can decide upon.

````go
// +gen List:"SortWith"
type Example struct { ... }
````

 * **SortWith**, **IsSortedWith**
 * **SortWithDesc**, **IsSortedWithDesc**

### Next: [Sets](Set.md)
#### Contents:

 * [Intro](README.md)
 * **Lists**
 * [Sets](Set.md)
 * [Options](Option.md)
 * [Joint Lists with Options and/or Sets](Unified.md)
 * [Plumbing functions](Plumbing.md)

### Generated Examples

 * [Num1: a simple numeric type](internal/list/test/other.go) generates [Num1List](internal/list/test/num1_list.go)
 * [Num2: a pointer numeric type](internal/list/test/other.go) generates [Num2List](internal/list/test/num2_list.go)
 * [Thing: a struct type](internal/list/test/thing.go) generates [ThingList](internal/list/test/thing_list.go) 

