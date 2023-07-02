package main

import "fmt"

//State interface
type State interface {
    scan()
    toggle_amfm()
}

type AbstractState struct {
    pos      int
    radio    *Radio
    name     string
    stations []string
}

func (p *AbstractState) scan() {
    p.pos++
    if p.pos == len(p.stations) {
        p.pos = 0
    }
    fmt.Println("Scanning... Station is", p.stations[p.pos], p.name)
}

//implementor 1
type AmState struct{ AbstractState }

func (p *AmState) toggle_amfm() {
    fmt.Println("Switching to FM")
    p.radio.state = p.radio.fmstate
}

//implementor 2
type FmState struct{ AbstractState }

func (p *FmState) toggle_amfm() {
    fmt.Println("Switching to AM")
    p.radio.state = p.radio.amstate
}

type Radio struct {
    amstate *AmState
    fmstate *FmState
    state   State
}

func (p *Radio) toggle_amfm() {
    p.state.toggle_amfm()
}

func (p *Radio) scan() {
    p.state.scan()
}

func new_radio() *Radio {
    radio := new(Radio)
    amstate := &AmState{AbstractState{0, radio, "AM", []string{"1250", "1380", "1510"}}}
    fmstate := &FmState{AbstractState{0, radio, "FM", []string{"81.3", "89.1", "103.9"}}}

    radio.amstate = amstate
    radio.fmstate = fmstate
    radio.state = amstate

    return radio
}

func main() {
    radio := new_radio()

    //am
    radio.scan()

    //fm
    radio.toggle_amfm()
    radio.scan()

    //am
    radio.toggle_amfm()
    radio.scan()
    radio.toggle_amfm()
    radio.toggle_amfm()
    radio.scan()
}
