package main

import "fmt"

type Cat struct{ name string }

func (p *Cat) meow() string {
	return "meow"
}

type Human struct{ name string }

func (p *Human) speak() string {
	return "hello"
}

type Car struct{ name string }

func (p *Car) make_noise(level int) string {
	return fmt.Sprintf("vroom %d", level)
}

type Soundable interface {
	make_noise() string
	get_name() string
}

type CarAdaptor struct{ obj *Car }

func (p *CarAdaptor) make_noise() string {
	return p.obj.make_noise(3)
}

func (p *CarAdaptor) get_name() string {
	return p.obj.name
}

type HumanAdaptor struct{ obj *Human }

func (p *HumanAdaptor) make_noise() string {
	return p.obj.speak()
}

func (p *HumanAdaptor) get_name() string {
	return p.obj.name
}

type CatAdaptor struct{ obj *Cat }

func (p *CatAdaptor) make_noise() string {
	return p.obj.meow()
}

func (p *CatAdaptor) get_name() string {
	return p.obj.name
}

func main() {
	list := []Soundable{&CarAdaptor{&Car{"Jeep"}}, &HumanAdaptor{&Human{"Tom"}}, &CatAdaptor{&Cat{"Persian Cat"}}}

	for _, adaptor := range list {
		fmt.Printf("%s goes %s\n", adaptor.get_name(), adaptor.make_noise())
	}
}
