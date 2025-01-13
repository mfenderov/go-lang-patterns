package main

import "fmt"
import "factory/pizza"

func main() {

	//we want to order a pizza from a store
	//how can we do it?
	//simplest way to oder a pizza
	simpleCheesePizza := pizza.CheesePizza{
		Doe:      "thin",
		Size:     "large",
		Toppings: []string{"cheese", "extra cheese"},
	}
	//What state is the pizza in???
	fmt.Println(simpleCheesePizza.Deliver())

	simplePepperoniPizza := pizza.PepperoniPizza{
		Doe:      "thin",
		Size:     "large",
		Toppings: []string{"pepperoni", "extra cheese"},
	}
	fmt.Println(simplePepperoniPizza.Deliver())

	//We can create a pizza with a more conventional way
	simpleCheesePizza2 := pizza.NewCheesePizza()
	fmt.Println(simpleCheesePizza2.Deliver())

	simplePepperoniPizza2 := pizza.NewPepperoniPizza()
	fmt.Println(simplePepperoniPizza2.Deliver())

	//We can also create a more generic pizza
	cheesePizza := pizza.NewPizza("thin", "large", "cheese", "extra cheese")
	fmt.Println(cheesePizza.Deliver())
	pepperoniPizza := pizza.NewPizza("thin", "large", "pepperoni", "extra cheese")
	fmt.Println(pepperoniPizza.Deliver())

	//We can order pizza without knowing how exactly pizza is make
	//We just need to know the type of pizza we want
	cheesePizza = pizza.OrderPizza(pizza.CHEESE)
	fmt.Println(cheesePizza.Deliver())

	pepperoniPizza = pizza.OrderPizza(pizza.PEPPERONI)
	fmt.Println(pepperoniPizza.Deliver())
}
