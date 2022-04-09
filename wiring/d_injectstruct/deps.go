package main

type Foo struct{}

func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct{}

func NewBar() *Bar {
	return &Bar{}
}

type Foobar struct {
	f *Foo
	b *Bar
}

func NewFoobar(f *Foo, b *Bar) *Foobar {
	return &Foobar{f: f, b: b}
}

type Service struct {
	f *Foo
	b *Bar
	m *Foobar
}
