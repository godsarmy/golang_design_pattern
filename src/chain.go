package main

import "fmt"

type Handler interface {
    handle(request int)
}

//implementor 1
type ConcreteHandler1 struct {
    successor Handler
}

func (p *ConcreteHandler1) handle(request int) {
    if request > 0 && request <= 10 {
        fmt.Println(request, "in handler1")
    } else {
        p.successor.handle(request)
    }
}

//implementor 2
type ConcreteHandler2 struct {
    successor Handler
}

func (p *ConcreteHandler2) handle(request int) {
    if request > 10 && request <= 20 {
        fmt.Println(request, "in handler2")
    } else {
        p.successor.handle(request)
    }
}

//implementor 3 
type ConcreteHandler3 struct {
    successor Handler
}

func (p *ConcreteHandler3) handle(request int) {
    if request > 20 && request <= 30 {
        fmt.Println(request, "in handler3")
    } else {
        fmt.Println("end of chain, no handler for", request)
    }
}

func main() {
    h3 := &ConcreteHandler3{}
    h2 := &ConcreteHandler2{h3}
    h1 := &ConcreteHandler1{h2}

    list := []int{2, 5, 14, 22, 18, 3, 35, 27, 20}
    for _, i := range list {
        h1.handle(i)
    }
}
