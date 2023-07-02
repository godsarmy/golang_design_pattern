package main

import "fmt"

type PetShop struct {
    pet_factory Factory
}

func (p *PetShop) show_pet() {
    pet := p.pet_factory.get_pet()
    fmt.Println("This is a lovely", pet)
    fmt.Println("It says", pet.speak())
    fmt.Println("It eats", p.pet_factory.get_food())
}

type Factory interface {
    get_food () string
    get_pet () Pet
}

type DogFactory struct {}
type CatFactory struct {}

func (p *DogFactory) get_food() string {
    return "dog food"
}

func (p *DogFactory) get_pet() Pet {
    return new(Dog)
}

func (p *CatFactory) get_food() string {
    return "cat food"
}

func (p *CatFactory) get_pet() Pet {
    return new(Cat)
}

type Pet interface {
    speak() string
    String() string
}

type Dog struct { }
type Cat struct { }

func (p *Dog) speak() string {
    return "woaf"
}

func (p *Dog) String() string {
    return "Dog"
}

func (p *Cat) speak() string {
    return "meow"
}

func (p *Cat) String() string {
    return "Cat"
}

func main() {
    list := []Factory{new(DogFactory), new(CatFactory)}
    shop := PetShop{}

    for _, factory := range list {
        shop.pet_factory = factory
        shop.show_pet()
    }
}
