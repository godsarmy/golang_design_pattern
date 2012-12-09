//implement of http://en.wikipedia.org/wiki/Memento_pattern
package main

import "fmt"

type Memento struct {
    state string
}

func (p *Memento) getSavedState() string {
    return p.state
}

type Originator struct {
    state string
}

func (p *Originator) set(state string) {
    fmt.Println("Originator: Setting state to", state)
    p.state = state
}

func (p *Originator) saveToMemento() Memento {
    fmt.Println("Originator: Saving to Memento.")
    return Memento{p.state}
}

func (p *Originator) restoreFromMemento(memento Memento) {
    p.state = memento.getSavedState()
    fmt.Println("Originator: State after restoring from Memento:", p.state)
}

func main() {
    //do some caretaker staff
    var savedStates []Memento

    originator := Originator{}
    originator.set("State1")
    originator.set("State2")
    savedStates = append(savedStates, originator.saveToMemento())
    originator.set("State3")

    // We can request multiple mementos, and choose which one to roll back to.
    savedStates = append(savedStates, originator.saveToMemento())
    originator.set("State4")

    originator.restoreFromMemento(savedStates[0])
    originator.restoreFromMemento(savedStates[1])
}
