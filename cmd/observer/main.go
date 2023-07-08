package main

import "fmt"

// data with list of observers
type Data struct {
	name      string
	data      int
	observers []Observer
}

func (p *Data) attach(obj Observer) {
	p.observers = append(p.observers, obj)
}

func (p *Data) detach(obj Observer) {
	new_list := []Observer{}
	for _, viewer := range p.observers {
		if viewer != obj {
			new_list = append(new_list, viewer)
		}
	}
	p.observers = new_list
}

func (p *Data) notify() {
	for _, viewer := range p.observers {
		viewer.update(p)
	}
}

func (p *Data) set_data(data int) {
	p.data = data
	p.notify()
}

type Observer interface {
	update(data *Data)
}

// implementor of Observer
type HexViewer struct{}
type DecViewer struct{}

func (p *HexViewer) update(data *Data) {
	fmt.Printf("HexViewer: Subject %s has data 0x%x\n", data.name, data.data)
}

func (p *DecViewer) update(data *Data) {
	fmt.Printf("DecViewer: Subject %s has data %d\n", data.name, data.data)
}

func main() {
	data1 := &Data{name: "Data 1"}
	data2 := &Data{name: "Data 2"}
	view1 := &HexViewer{}
	view2 := &DecViewer{}

	data1.attach(view1)
	data1.attach(view2)
	data2.attach(view2)
	data2.attach(view1)

	fmt.Println("Setting Data 1 = 10")
	data1.set_data(10)
	fmt.Println("Setting Data 2 = 15")
	data2.set_data(15)
	fmt.Println("Setting Data 1 = 3")
	data1.set_data(3)
	fmt.Println("Setting Data 2 = 5")
	data2.set_data(5)

	data1.detach(view1)
	data2.detach(view1)
	fmt.Println("Setting Data 1 = 10")
	data1.set_data(10)
	fmt.Println("Setting Data 2 = 15")
	data2.set_data(15)
}
