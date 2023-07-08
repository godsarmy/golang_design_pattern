// golang version of java code in http://en.wikipedia.org/wiki/Strategy_pattern
package main

import "fmt"

// strategy interface
type Strategy interface {
	execute(a, b int) int
}

// Concrete strategies
type StrategyAdd struct{}
type StrategySubstract struct{}
type StrategyMultiply struct{}

func (p *StrategyAdd) execute(a, b int) int {
	fmt.Println("Called StrategyAdd's execute()")
	return a + b
}

func (p *StrategySubstract) execute(a, b int) int {
	fmt.Println("Called StrategySubstract's execute()")
	return a - b
}

func (p *StrategyMultiply) execute(a, b int) int {
	fmt.Println("Called StrategyMultiply's execute()")
	return a * b
}

type Context struct {
	strategy Strategy
}

func (p *Context) execute_strategy(a, b int) int {
	return p.strategy.execute(a, b)
}

func main() {
	var context Context

	// Three contexts following different strategies
	context = Context{new(StrategyAdd)}
	resultA := context.execute_strategy(3, 4)

	context = Context{new(StrategySubstract)}
	resultB := context.execute_strategy(3, 4)

	context = Context{new(StrategyMultiply)}
	resultC := context.execute_strategy(3, 4)

	fmt.Println("Result A : ", resultA)
	fmt.Println("Result B : ", resultB)
	fmt.Println("Result C : ", resultC)
}
