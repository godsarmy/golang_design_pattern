// golang version of java code in http://en.wikipedia.org/wiki/Strategy_pattern
package main

import "fmt"

// Mediator
type Mediator interface {
	send(message string, colleague Colleague)
}

type ConcreteMediator struct {
	colleague1 *ConcreteColleague1
	colleague2 *ConcreteColleague2
}

func (p *ConcreteMediator) send(message string, colleague Colleague) {
	if colleague == p.colleague1 {
		p.colleague2.notify(message)
	} else {
		p.colleague1.notify(message)
	}
}

// Colleague
type Colleague interface{}

// Concrete Colleagues
type ConcreteColleague1 struct {
	mediator Mediator
}

func (p *ConcreteColleague1) send(message string) {
	p.mediator.send(message, p)
}

func (p *ConcreteColleague1) notify(message string) {
	fmt.Println("Colleague1 gets message:", message)
}

type ConcreteColleague2 struct {
	mediator Mediator
}

func (p *ConcreteColleague2) send(message string) {
	p.mediator.send(message, p)
}

func (p *ConcreteColleague2) notify(message string) {
	fmt.Println("Colleague2 gets message:", message)
}

func main() {
	m := &ConcreteMediator{}
	c1 := &ConcreteColleague1{m}
	c2 := &ConcreteColleague2{m}
	m.colleague1 = c1
	m.colleague2 = c2
	c1.send("How are you?")
	c2.send("Fine, thanks")
}
