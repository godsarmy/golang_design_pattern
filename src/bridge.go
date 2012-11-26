package main

import "fmt"

//implementor
type DrawAPI interface { 
    drawCircle (x int, y int, radius int)
}

//implementor 1
type DrawAPI1 struct {}

func (p DrawAPI1) drawCircle(x int, y int, radius int) { 
    fmt.Printf("API1.circle at %d:%d radius %d\n", x, y, radius)
}

//implementor 2
type DrawAPI2 struct {}

func (p DrawAPI2) drawCircle(x int, y int, radius int) { 
    fmt.Printf("API2.circle at %d:%d radius %d\n", x, y, radius)
}

//refined abstraction
type CircleShape struct { 
    x, y, radius int
    drawingAPI DrawAPI
}

func (p *CircleShape) draw () {
    p.drawingAPI.drawCircle(p.x, p.y, p.radius)
}

func (p *CircleShape) resize (factor int) {
    p.radius *= factor
}

func main() {
    list := []CircleShape{CircleShape{2,3,4,new(DrawAPI1)}, CircleShape{5,7,11,new(DrawAPI2)}}

    for _, shape := range list {
        shape.resize(3)
        shape.draw()
    }
}
