package utils

import (
	"dummy-app/models"
	"fmt"
)

type MyProduct models.Product //alias

func (p MyProduct) FormatProduct() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

// Behavior as a method (function with receiver)
func (p *MyProduct) ApplyDiscount(discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}
