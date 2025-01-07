package main

import "fmt"
import "factory/pizza"

func main() {

	//We con order pizza without knowing how exactly pizza is make
	//We just need to know the type of pizza we want

	cheesePizza := pizza.OrderPizza(pizza.CHEESE)
	fmt.Println(cheesePizza.Deliver())

	pepperoniPizza := pizza.OrderPizza(pizza.PEPPERONI)
	fmt.Println(pepperoniPizza.Deliver())
}
