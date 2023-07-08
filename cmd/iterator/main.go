package main

import "fmt"

// original iterator
type Iterator struct {
	current  int
	data_set *DataSet
}

func (p *Iterator) has_next() bool {
	length := len(*p.data_set)
	rc := true
	if p.current >= length {
		rc = false
	}
	return rc
}

func (p *Iterator) next() interface{} {
	data := (*p.data_set)[p.current]
	p.current++
	return data
}

type DataSet []string

func (p *DataSet) get_iterator() *Iterator {
	iter := &Iterator{0, p}
	return iter
}

func main() {
	data := &DataSet{"foo", "bar", "hello", "world"}

	for iter := data.get_iterator(); iter.has_next(); {
		value := iter.next()
		fmt.Println("Value: ", value)
	}
}
