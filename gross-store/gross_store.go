package gross

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
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, unitMap map[string]int, billKey, unitkey string) bool {

	_, unitExists := unitMap[unitkey]
	if !unitExists {
		return false
	}

	billItemValue := bill[billKey]
	bill[billKey] = billItemValue + unitMap[unitkey]

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, unitMap map[string]int, billKey, unitKey string) bool {

	value, billKeyExists := bill[billKey]
	if !billKeyExists {
		return false
	}

	unit, unitKeyExists := unitMap[unitKey]
	if !unitKeyExists {
		return false
	}

	switch {
	case value-unit < 0:
		// Return false if the new quantity would be less than 0.
		return false
	case value-unit == 0:
		// If the new quantity is 0, completely remove the item from the bill then return true.
		delete(bill, billKey)
	default:
		// Otherwise, reduce the quantity of the item and return true.
		bill[billKey] = value - unit
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, billKey string) (int, bool) {
	val, exists := bill[billKey]

	return val, exists
}
