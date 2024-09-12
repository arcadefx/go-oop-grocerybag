package grocerybag

import (
	"fmt"

	"github.com/arcadefx/go-oop-grocerybag/labels"
)

type Bag interface {
	Add(name string, cost float64)
	PrintList()
}

type Item struct {
	Name string
	Cost float64
}

type GroceryBag struct {
	Items []Item
	Label labels.Label
}

func NewBag(
	items []Item,
	label labels.Label,
) Bag {
	return &GroceryBag{
		Items: items,
		Label: label,
	}
}

func (t *GroceryBag) Add(name string, cost float64) {
	t.Items = append(t.Items, Item{
		Name: name,
		Cost: cost,
	})
}

func (t *GroceryBag) PrintList() {
	fmt.Printf("Items in the bag for: %s\n", t.Label.GetLabel())
	for _, item := range t.Items {
		fmt.Printf("%s \t $%.2f\n", item.Name, item.Cost)
	}
	fmt.Printf("Total cost: $%.2f\n", t.totalCost())
	fmt.Println()
}

func (t *GroceryBag) totalCost() float64 {
	total := 0.0
	for _, item := range t.Items {
		total += item.Cost
	}
	return total
}
