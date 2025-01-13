package main

import "fmt"
import "strategy/duck"

func main() {
	//Create different types of ducks

	// A majestic mallard duck
	mallard := duck.Duck{
		Name:          "Mallard",
		FlyBehavior:   duck.FlyWithWings{},
		QuackBehavior: duck.LoudQuack{},
	}

	// A rubber duck that squeaks but can't fly
	rubberDuck := duck.Duck{
		Name:          "Rubber Duck",
		FlyBehavior:   duck.FlyNoWay{},
		QuackBehavior: duck.Squeak{},
	}

	// A stealthy decoy duck that's silent and doesn't fly
	decoyDuck := duck.Duck{
		Name:          "Decoy Duck",
		FlyBehavior:   duck.FlyNoWay{},
		QuackBehavior: duck.SilentQuack{},
	}

	// Interacting with the ducks
	ducks := []duck.Duck{mallard, rubberDuck, decoyDuck}

	for _, duck := range ducks {
		duck.Display()
		duck.PerformFly()
		duck.PerformQuack()
		fmt.Println()
	}

	// Change behavior dynamically
	fmt.Println("Uh oh! Rubber Duck got wings!")
	rubberDuck.SetFlyBehavior(duck.FlyWithWings{})
	rubberDuck.PerformFly()
}
