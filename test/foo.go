package main

// +test List:" Option "
type Foo1 Underlying

// +test Option:" List "
type Foo2 Underlying

// +test * List:" Option "
type Foo3 Underlying

// +test * Option:" List "
type Foo4 Underlying

type Underlying int
