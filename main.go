package main

import (
	"github.com/arcadefx/go-oop-grocerybag/grocerybag"
	"github.com/arcadefx/go-oop-grocerybag/labels"
)

func main() {
	bag := grocerybag.NewBag([]grocerybag.Item{
		{
			Name: "Apple",
			Cost: 0.5,
		},
		{
			Name: "Banana",
			Cost: 0.25,
		},
		{
			Name: "Orange",
			Cost: 0.75,
		},
	}, &labels.LabelImpl{
		Label: "John Smith",
	})
	bag.Add("Milk", 2.99)
	bag.PrintList()

	bag2 := grocerybag.NewBag([]grocerybag.Item{
		{
			Name: "Apple",
			Cost: 0.5,
		},
	}, &labels.LabelImpl{
		Label: "Jane Smith",
	})
	bag2.PrintList()
}
