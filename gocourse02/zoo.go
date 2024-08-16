package main

import "fmt"

type Animal struct {
	ID        int
	Name      string
	Species   string
	Condition AnimalCondition
}

type AnimalCondition struct {
	Status  string
	Mood    string
	Satiety int
}

func (an *Animal) eat(food string) {
	an.Condition.Status = "eating"
	an.Condition.Mood = "happy"
	an.Condition.Satiety += 5

	fmt.Printf("Animal %s is eating %s.\n", an.Name, food)
}

func NewAnimal(id int, name string, species string, condition AnimalCondition) *Animal {
	return &Animal{
		ID:        id,
		Name:      name,
		Species:   species,
		Condition: condition,
	}
}

type Zookeeper struct {
	ID   int
	Name string
}

func (zk *Zookeeper) feed(food string, an *Animal) {
	an.eat(food)
}

func (zk *Zookeeper) sleep(cage *Cage) {
	cage.open()
}

func (zk *Zookeeper) checkAndCatch(animal *Animal, cage *Cage) {
	if animal.Condition.Satiety >= 10 {
		animal.Condition.Status = "caught"
		cage.close(animal)

		fmt.Printf("Animal %s is caught.\n", animal.Name)
	} else {
		cage.open()
		fmt.Printf("Animal %s is dangerous.\n", animal.Name)
	}
}

const (
	StateOpen  = "open"
	StateClose = "close"
)

type Cage struct {
	ID     int
	State  string
	Animal *Animal
}

func (c *Cage) open() {
	c.State = StateOpen
	c.Animal = nil

	fmt.Printf("The cage is opened. There is no animal in the cage.\n")
}

func (c *Cage) close(animal *Animal) {
	c.State = StateClose
	c.Animal = animal

	fmt.Printf("The cage is closed. Animal %s is in the cage.\n", c.Animal.Name)
}
