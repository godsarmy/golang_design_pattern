package main

import "fmt"

type MoveFileCommand struct {
	src  string
	dest string
}

func (p *MoveFileCommand) do() {
	fmt.Printf("move %s to %s\n", p.src, p.dest)
}

func (p *MoveFileCommand) undo() {
	fmt.Printf("move %s to %s\n", p.dest, p.src)
}

func main() {
	var list []MoveFileCommand
	list = append(list, MoveFileCommand{"foo.txt", "bar.txt"})
	list = append(list, MoveFileCommand{"bar.txt", "baz.txt"})

	for _, ren := range list {
		ren.do()
	}

	var ren1, ren2 MoveFileCommand
	ren2, list = list[len(list)-1], list[:len(list)-1]
	ren2.undo()
	ren1, list = list[len(list)-1], list[:len(list)-1]
	ren1.undo()
}
