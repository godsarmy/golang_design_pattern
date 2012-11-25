package main

import (
    "fmt"
    "time"
)

const (
    SLEEP = 1*time.Second
)

type Runnable interface {
    run()
}

type T1 struct {}
type T2 struct {}
type T3 struct {}

func (p *T1) run() { 
    fmt.Println("###### In Test 1 ######")
    time.Sleep(SLEEP)
    fmt.Println("Setting up")
    time.Sleep(SLEEP)
    fmt.Println("Running test")
    time.Sleep(SLEEP)
    fmt.Println("Tearing down")
    time.Sleep(SLEEP)
    fmt.Println("Test Finished")
}

func (p *T2) run() { 
    fmt.Println("###### In Test 2 ######")
    time.Sleep(SLEEP)
    fmt.Println("Setting up")
    time.Sleep(SLEEP)
    fmt.Println("Running test")
    time.Sleep(SLEEP)
    fmt.Println("Tearing down")
    time.Sleep(SLEEP)
    fmt.Println("Test Finished")
}

func (p *T3) run() { 
    fmt.Println("###### In Test 3 ######")
    time.Sleep(SLEEP)
    fmt.Println("Setting up")
    time.Sleep(SLEEP)
    fmt.Println("Running test")
    time.Sleep(SLEEP)
    fmt.Println("Tearing down")
    time.Sleep(SLEEP)
    fmt.Println("Test Finished")
}

//facade
type TestRunner struct {
    tests []Runnable
}

//constructor
func NewTestRunnable () *TestRunner {
    p := &TestRunner{tests:[]Runnable{&T1{}, &T2{}, &T3{}}}
    return p
}

func (p *TestRunner) runAll() {
    for _, t := range p.tests {
        t.run()
    }
}

func main() {
    testRunner := NewTestRunnable()
    testRunner.runAll()
}
