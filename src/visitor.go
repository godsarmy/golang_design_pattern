package main

import (
    "fmt"
    "reflect"
)

//Elements
type Element interface {
    accept(visitor *Visitor)
}

type A struct{}

func (p *A) accept(visitor *Visitor) {
    fmt.Println("In accept A")
    visitor.visit(p)
}

type B struct{}

func (p *B) accept(visitor *Visitor) {
    fmt.Println("In accept B")
    visitor.visit(p)
}

type C struct {
    a   A
    b   B
}

func (p *C) accept(visitor *Visitor) {
    fmt.Println("In accept C")
    visitor.visit(&p.a)
    visitor.visit(&p.b)
    visitor.visit(p)
}

//vistor 
type Visitor struct{}

func (p *Visitor) visit(e Element) {
    if typ := reflect.TypeOf(e); typ == reflect.TypeOf(&B{}) {
        p.visit_B(e)
        return
    }

    p.generic_visit(e)
}

func (p *Visitor) visit_B(e Element) {
    typ := reflect.TypeOf(e)
    fmt.Println("In Visit B:", typ)
}

func (p *Visitor) generic_visit(e Element) {
    typ := reflect.TypeOf(e)
    fmt.Println("In Generic Visit:", typ)
}

func main() {
    v := &Visitor{}
    a := A{}
    b := B{}
    c := C{}
    a.accept(v)
    b.accept(v)
    c.accept(v)
}
