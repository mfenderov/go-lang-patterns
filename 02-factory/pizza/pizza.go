package pizza

type PizzaType int

const (
	CHEESE PizzaType = iota
	PEPPERONI
)

type Pizza interface {
	Deliver() string
}

type CheesePizza struct{}

func (p CheesePizza) Deliver() string { return "Delivering a cheesy pizza!" }

type PepperoniPizza struct{}

func (p PepperoniPizza) Deliver() string { return "Delivering a pepperoni pizza!" }

func OrderPizza(pizzaType PizzaType) Pizza {
	switch pizzaType {
	case CHEESE:
		return CheesePizza{}
	case PEPPERONI:
		return PepperoniPizza{}
	default:
		panic("Unknown pizza type!")
	}
}
