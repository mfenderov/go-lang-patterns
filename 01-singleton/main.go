package main

import (
	"singleton/coffee-machine"
	singletonmachine "singleton/coffee-machine-singleton"
)

func main() {
	//We're running a small office kitchen and, of course, we want to have a coffee machine
	coffeeMachine := coffeemachine.BuyNewMachine()
	coffeeMachine.MakeCoffee()

	//completely different person in the office wants a coffee
	coffeeMachine2 := coffeemachine.BuyNewMachine()
	coffeeMachine2.MakeCoffee()

	/*
		How we have 2 coffee machines, and that's not what we want

		Coffee machines are expensive, and it doesn't make sense to buy a new one each time someone wants a coffee
		It makes sense to have a single Coffe Machine per kitchen and make sure it's SHARED
		We can pass the same coffee machine around, sure, but we can guarantee that there's only one, and that someone else doesn't buy a new one
		This is where the singleton pattern comes in
		We can use the same coffee machine, and make sure that it's the same one, and that it's shared
	*/

	coffeeMachineSingleton := singletonmachine.BuyNewMachine()
	coffeeMachineSingleton.MakeCoffee()

	//totally different file
	coffeeMachineSingleton2 := singletonmachine.BuyNewMachine()
	coffeeMachineSingleton2.MakeCoffee()

	//no matter how many times someone will try to buy a new machine, it will always be the same one

	/*
		Practicality:
			Pooling Database Connections
			we don't want to open a new connection each time we want to query the database - it's slow and expensive, we have to send credentials, etc.
			instead we want to re-use connections when possible
			for that, we can use the singleton pattern to establish a connection pool once, and then re-use it throughout the application
	*/
}
