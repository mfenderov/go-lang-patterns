package pizza

type PizzaType int

const (
	CHEESE PizzaType = iota
	PEPPERONI
)

type Pizza interface {
	Deliver() string
}

type CheesePizza struct {
	Doe      string
	Size     string
	Toppings []string
}

type GenericPizza struct {
	Doe      string
	Size     string
	Toppings []string
}

func (p GenericPizza) Deliver() string { return "Delivering a generic pizza!" }

func NewPizza(doe string, size string, toppings ...string) Pizza {
	return GenericPizza{
		Doe:      doe,
		Size:     size,
		Toppings: toppings,
	}
}

func (p CheesePizza) Deliver() string { return "Delivering a cheesy pizza!" }

type PepperoniPizza struct {
	Doe      string
	Size     string
	Toppings []string
}

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

func NewCheesePizza() Pizza {
	return CheesePizza{
		Doe:      "thin",
		Size:     "large",
		Toppings: []string{"cheese", "extra cheese"},
	}
}

func NewPepperoniPizza() Pizza {
	return PepperoniPizza{
		Doe:      "thin",
		Size:     "large",
		Toppings: []string{"pepperoni", "extra cheese"},
	}
}
