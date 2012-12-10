//implement of http://en.wikipedia.org/wiki/Memento_pattern
package main

import "fmt"

//Abstract class
type Abstract interface {
    modify_list()
    get_result() string
}

//Concrete class
type ByteList []byte
type IntList []int

type Template struct {
    list Abstract
}

func (p *Template) templateMethod() {
    p.list.iterate_list()
    fmt.Println(p.list.get_result())
}

func (p IntList) modify_list() {
    for i, _ := range []int(p) {
        p[i] *= 2
    }
}

func (p ByteList) modify_list() {
    for i, _ := range []byte(p) {
        p[i]++
    }
}

func (p IntList) get_result() string {
    var s string
    for i, o := range []int(p) {
        s += fmt.Sprintf("%d:%d ", i, o)
    }
    return s
}

func (p ByteList) get_result() string {
    var s string
    for i, o := range []byte(p) {
        s += fmt.Sprintf("%d-%c ", i, o)
    }
    return s
}

func main() {
    t1 := Template{ByteList{'w', 'a', 'l', 't'}}
    t2 := Template{IntList{1, 9, 8, 1, 1, 0, 1, 0}}
    t1.templateMethod()
    t2.templateMethod()
}
