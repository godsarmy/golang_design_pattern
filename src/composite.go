package main

import "fmt"

type Graphic interface {
    print ()
}

type CompositeGraphic struct {
    name string
    childGraphics []Graphic
}

func (p *CompositeGraphic) print () {
    fmt.Println("Composite", p.name)
    for _, graphic := range p.childGraphics {
        graphic.print()
    }
}

func (p *CompositeGraphic) add (graphic Graphic) {
    p.childGraphics = append(p.childGraphics, graphic)
}

func (p *CompositeGraphic) remove (graphic Graphic) {
    new_list := []Graphic{}
    for _, obj := range p.childGraphics {
        if graphic != obj {
            new_list = append(new_list, obj)
        }
    }
    p.childGraphics = new_list
}

type Ellipse struct {
    name string
}

func (p *Ellipse) print () {
    fmt.Println("Ellipse", p.name)
}

func main() {
    ellipse1 := &Ellipse{"ellipse1"}
    ellipse2 := &Ellipse{"ellipse2"}
    ellipse3 := &Ellipse{"ellipse3"}
    ellipse4 := &Ellipse{"ellipse4"}

    //Initialize three composite graphics
    graphic := &CompositeGraphic{name:"composite"}
    graphic1 := &CompositeGraphic{name:"composite1"}
    graphic2 := &CompositeGraphic{name:"composite2"}

    //Composes the graphics
    graphic1.add(ellipse1)
    graphic1.add(ellipse2)
    graphic1.add(ellipse3)

    graphic2.add(ellipse4)

    graphic.add(graphic1)
    graphic.add(graphic2)

    //Prints the complete graphic (four times the string "Ellipse").
    graphic.print()
}
