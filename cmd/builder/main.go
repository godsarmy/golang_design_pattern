package main

import (
	"fmt"
)

type Director struct {
	builder Builder
}

func (p *Director) construct_building() {
	p.builder.new_building()
	p.builder.build_floor()
	p.builder.build_size()
}

func (p *Director) get_building() *Building {
	return p.builder.get_building()
}

// builder interface
type Builder interface {
	new_building()
	build_floor()
	build_size()
	get_building() *Building
}

// Concrete builder
type BuilderHouse struct {
	building *Building
}

func (p *BuilderHouse) new_building() {
	p.building = new(Building)
}

func (p *BuilderHouse) build_floor() {
	p.building.floor = "One"
}

func (p *BuilderHouse) build_size() {
	p.building.size = "Big"
}

func (p *BuilderHouse) get_building() *Building {
	return p.building
}

type BuilderFlat struct {
	building *Building
}

func (p *BuilderFlat) new_building() {
	p.building = new(Building)
}

func (p *BuilderFlat) build_floor() {
	p.building.floor = "More than One"
}

func (p *BuilderFlat) build_size() {
	p.building.size = "Small"
}

func (p *BuilderFlat) get_building() *Building {
	return p.building
}

// Product
type Building struct {
	floor, size string
}

func (p *Building) String() string {
	return fmt.Sprintf("Floor: %s | Size: %s", p.floor, p.size)
}

func main() {
	var building *Building
	director := Director{}
	director.builder = new(BuilderHouse)
	director.construct_building()
	building = director.get_building()
	fmt.Println(building)
	director.builder = new(BuilderFlat)
	director.construct_building()
	building = director.get_building()
	fmt.Println(building)
}
