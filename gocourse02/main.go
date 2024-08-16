package main

func main() {
	zookeeper := Zookeeper{ID: 1, Name: "Jon Snow"}
	cageTiger := &Cage{ID: 1, State: stateClose}
	cageCoyote := &Cage{ID: 2, State: stateClose}
	cageBeaver := &Cage{ID: 2, State: stateClose}

	zookeeper.sleep(cageTiger)
	zookeeper.sleep(cageCoyote)
	zookeeper.sleep(cageBeaver)

	tiger := NewAnimal(1, "Mia", "Tiger", AnimalCondition{Status: "free", Mood: "happy", Satiety: 10})
	newTiger := tiger.multiply()
	newTiger.Name = "Born"

	coyote := NewAnimal(2, "Gena", "Coyote", AnimalCondition{Status: "free", Mood: "happy", Satiety: 0})
	beaver := NewAnimal(3, "Bobr", "Beaver", AnimalCondition{Status: "free", Mood: "happy", Satiety: 5})

	zookeeper.feed("meet", newTiger)
	zookeeper.feed("meet", tiger)
	zookeeper.feed("meet", coyote)
	zookeeper.feed("fish", beaver)

	zookeeper.checkAndCatch(newTiger, cageTiger)
	zookeeper.checkAndCatch(tiger, cageTiger)
	zookeeper.checkAndCatch(coyote, cageCoyote)
	zookeeper.checkAndCatch(beaver, cageBeaver)

	zookeeper.feed("meet", coyote)
	zookeeper.checkAndCatch(coyote, cageCoyote)
}
