package duck

import "fmt"

// FlyBehavior defines the fly behavior of a duck
type FlyBehavior interface {
	Fly()
}

// QuackBehavior defines the quack behavior of a duck
type QuackBehavior interface {
	Quack()
}

// Implement concrete behaviors
// Flying behaviors

// FlyWithWings is a duck that can fly
type FlyWithWings struct{}

func (f FlyWithWings) Fly() {
	fmt.Println("Flapping wings and flying high! 🦅")
}

// FlyNoWay is a duck that can't fly
type FlyNoWay struct{}

func (f FlyNoWay) Fly() {
	fmt.Println("I can't fly... 🐥")
}

// Quacking behaviors

// LoudQuack is a duck that quacks loudly
type LoudQuack struct{}

func (q LoudQuack) Quack() {
	fmt.Println("QUACK QUACK!! 🔊")
}

// Squeak is a duck that squeaks
type Squeak struct{}

func (q Squeak) Quack() {
	fmt.Println("Squeak squeak! 🐤")
}

// SilentQuack is a duck that doesn't quack
type SilentQuack struct{}

func (q SilentQuack) Quack() {
	fmt.Println("...")
}

// Define the Duck type with behavior interfaces

type Duck struct {
	FlyBehavior   FlyBehavior
	QuackBehavior QuackBehavior
	Name          string
}

func (d Duck) SetFlyBehavior(fb FlyBehavior) {
	d.FlyBehavior = fb
}

func (d Duck) SetQuackBehavior(qb QuackBehavior) {
	d.QuackBehavior = qb
}

// Duck performs its behaviors

func (d Duck) PerformFly() {
	fmt.Printf("%s: ", d.Name)
	d.FlyBehavior.Fly()
}

func (d Duck) PerformQuack() {
	fmt.Printf("%s: ", d.Name)
	d.QuackBehavior.Quack()
}

func (d Duck) Display() {
	fmt.Printf("Hi, I'm %s the duck! 🦆\n", d.Name)
}
