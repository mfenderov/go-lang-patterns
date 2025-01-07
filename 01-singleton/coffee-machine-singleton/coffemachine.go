package coffee_machine_singleton

import (
	"fmt"
	"math/rand"
	"sync"
)

var once sync.Once
var coffeeMachine *CoffeeMachine

type CoffeeMachine struct {
	name string
}

func BuyNewMachine() *CoffeeMachine {
	once.Do(func() {
		coffeeMachine = &CoffeeMachine{
			name: fmt.Sprintf("%d", rand.Int31()),
		}
	})
	return coffeeMachine
}

func (c *CoffeeMachine) MakeCoffee() {
	fmt.Println("Hi, I am your coffee machine:", c.name, "and I am making coffee")
}
