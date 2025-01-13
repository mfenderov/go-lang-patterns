package burger

type Burger struct {
	Patty   string
	Cheese  bool
	Lettuce bool
	Sauce   string
}

type burger struct {
	patty   string
	cheese  bool
	lettuce bool
	sauce   string
}

type BurgerBuilder struct {
	patty   string
	cheese  bool
	lettuce bool
	sauce   string
}

func NewBurgerBuilder() *BurgerBuilder {
	return &BurgerBuilder{}
}

func (b *BurgerBuilder) AddPatty(patty string) *BurgerBuilder {
	b.patty = patty
	return b
}

func (b *BurgerBuilder) AddCheese() *BurgerBuilder {
	b.cheese = true
	return b
}

func (b *BurgerBuilder) AddLettuce() *BurgerBuilder {
	b.lettuce = true
	return b
}
func (b *BurgerBuilder) AddSauce(sauce string) *BurgerBuilder {
	b.sauce = sauce
	return b
}

func (b *BurgerBuilder) Build() burger {
	if b.patty == "" {
		panic("Patty is required")
	}
	return burger{
		patty:   b.patty,
		cheese:  b.cheese,
		lettuce: b.lettuce,
		sauce:   b.sauce,
	}
}
