package main

import "fmt"

type Flyweight struct {
	a string
	b string
}

type FlyweightFactory struct {
	pool map[string]*Flyweight
}

func (p *FlyweightFactory) getFlyweight(a string, b string) *Flyweight {
	key := a + b
	if _, present := p.pool[key]; !present {
		p.pool[key] = &Flyweight{a, b}
	}
	return p.pool[key]
}

func compare(a, b *Flyweight) {
	if a == b {
		fmt.Println(a, "and", b, "are the same object")
	} else {
		fmt.Println(a, "and", b, "are different objects")
	}
}

func main() {
	factory := &FlyweightFactory{pool: make(map[string]*Flyweight)}

	obj1 := factory.getFlyweight("foo", "bar")
	obj2 := factory.getFlyweight("foo", "bar")
	obj3 := factory.getFlyweight("foo", "baz")

	compare(obj1, obj2)
	compare(obj1, obj3)
}
