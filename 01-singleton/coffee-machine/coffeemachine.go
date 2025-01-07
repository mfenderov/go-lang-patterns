package coffeemachine

import (
	"fmt"
	"math/rand"
)

type CoffeeMachine struct {
	name string
}

func BuyNewMachine() *CoffeeMachine {
	return &CoffeeMachine{
		name: fmt.Sprintf("%d", rand.Int31()),
	}
}

func (c *CoffeeMachine) MakeCoffee() {
	fmt.Println("Hi, I am your coffee machine:", c.name, "and I am making coffee")
}
