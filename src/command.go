package main

import "fmt"

type moveFileCommand struct { 
    src string
    dest string
}

func (p *moveFileCommand) do() { 
    fmt.Printf("move %s to %s\n" , p.src, p.dest)
}

func (p *moveFileCommand) undo() { 
    fmt.Printf("move %s to %s\n", p.dest, p.src)
}

func main() {
    var list []moveFileCommand
    list = append(list, moveFileCommand{"foo.txt", "bar.txt"})
    list = append(list, moveFileCommand{"bar.txt", "baz.txt"})

    for _, ren := range list {
        ren.do()
    }
    
    var ren1, ren2 moveFileCommand
    ren2, list = list[len(list)-1], list[:len(list)-1]
    ren2.undo()
    ren1, list = list[len(list)-1], list[:len(list)-1]
    ren1.undo()
}
