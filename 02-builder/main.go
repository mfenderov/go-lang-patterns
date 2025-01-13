package main

import "fmt"
import "builder/burger"

func main() {
	//You’re at a burger joint, and they let you customize your burger
	simpleBurger := burger.Burger{
		Patty:   "banana",
		Cheese:  false,
		Lettuce: false,
		Sauce:   "chocolate",
	}
	//What the hell did we just made?
	fmt.Printf("Your burger: %+v\n", simpleBurger)

	//The joint wants to make sure that your configuration is valid
	//Instead of giving you access to burger directly, they ask you to select already defined ingredients
	beefBurger := burger.NewBurgerBuilder().
		AddPatty("beef").
		AddCheese().
		AddLettuce().
		AddSauce("ketchup").
		Build()
	fmt.Printf("Your burger: %+v\n", beefBurger)

}
