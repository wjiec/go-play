package main

type FooConfig string

type Foo struct {
	cfg FooConfig
}

func NewFoo(cfg FooConfig) *Foo {
	return &Foo{cfg: cfg}
}

type BarConfig string

type Bar struct {
	cfg BarConfig
}

func NewBar(cfg BarConfig) (*Bar, error) {
	return &Bar{cfg: cfg}, nil
}

type FooBar struct {
	f *Foo
	b *Bar
}

func NewFooBar(f *Foo, b *Bar) *FooBar {
	return &FooBar{f: f, b: b}
}
