package main

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	amount, ok := units[unit]
	if ok {
		bill[item] += amount
	}
	return ok
}

// RemoveItem removes an item from customer bill.
func RemoveItem2(bill, units map[string]int, item, unit string) bool {
	quantity, okBill := bill[item]
	amount, okUnit := units[unit]
	result := okBill && okUnit && quantity >= amount
	if result {
		bill[item] -= amount
		if bill[item] == 0 {
			delete(bill, item)
		}
	}
	return result
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem2(bill map[string]int, item string) (int, bool) {
	quantity, ok := bill[item]
	return quantity, ok
}
