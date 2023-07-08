package singleton

import "fmt"

var _instance *object

type object struct {
	name string
}

func Instance() *object {
	if _instance == nil {
		_instance = new(object)
	}
	return _instance
}

func (p *object) Setname(name string) {
	p.name = name
}

func (p *object) Say() {
	fmt.Println(p.name)
}
