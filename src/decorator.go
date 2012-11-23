package main

import "fmt"

type foo struct {  }

func (p *foo) f1() { 
    fmt.Println("origianl f1")
}

func (p *foo) f2() { 
    fmt.Println("origianl f2")
}

type foo_decorator struct { decoratee *foo }

func (p *foo_decorator) f1() { 
    fmt.Println("decorated f1")
    p.decoratee.f1()
}

func (p *foo_decorator) f2() { 
    p.decoratee.f2()
}


func main() {
    var u = foo{}
    var v = foo_decorator{&u}
    v.f1()
    v.f2()
}
