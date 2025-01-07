package main

import "fmt"
import "builder/burger"

func main() {
	//You’re at a burger joint, and they let you customize your burger
	//Buy the joint wants to make sure that your configuration is valid
	//Instead of giving you access to burger directly, they ask you to fill a form
	//Why? Because there might be a dozens of ingredients with different restrictions and combinations

	beefBurger := burger.NewBurgerBuilder().
		AddPatty("beef").
		AddCheese().
		AddLettuce().
		AddSauce("ketchup").
		Build()
	fmt.Printf("Your burger: %+v\n", beefBurger)

}
