package main

import "fmt"

type Clonable interface {
	clone() Clonable
	get_data() int
	set_data(data int)
}

type ConcreteClonable struct {
	name string
	data int
}

func (p *ConcreteClonable) set_data(data int) {
	p.data = data
}

func (p *ConcreteClonable) get_data() int {
	return p.data
}

func (p *ConcreteClonable) clone() Clonable {
	new_obj := (*p)
	return &new_obj
}

func main() {
	var data1 = ConcreteClonable{"foo", 1}

	for i := 0; i < 10; i++ {
		new_data := data1.clone()
		new_data.set_data(i * new_data.get_data())
		fmt.Println("Check data: ", new_data)
	}
}
