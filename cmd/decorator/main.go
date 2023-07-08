package main

import "fmt"

type Foo struct{}

func (p *Foo) f1() {
	fmt.Println("origianl f1")
}

func (p *Foo) f2() {
	fmt.Println("origianl f2")
}

type FooDecorator struct{ decoratee *Foo }

func (p *FooDecorator) f1() {
	fmt.Println("decorated f1")
	p.decoratee.f1()
}

func (p *FooDecorator) f2() {
	p.decoratee.f2()
}

func main() {
	var u = Foo{}
	var v = FooDecorator{&u}
	v.f1()
	v.f2()
}
