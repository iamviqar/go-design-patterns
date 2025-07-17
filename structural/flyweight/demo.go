package main

import "fmt"

// TeaType represents flyweight interface
type TeaType interface {
	Serve(table int) string
}

// KarakTea - concrete flyweight
type KarakTea struct{}

func (k KarakTea) Serve(table int) string {
	return fmt.Sprintf("Serving Karak tea to table #%d", table)
}

// JasmineTea - concrete flyweight
type JasmineTea struct{}

func (j JasmineTea) Serve(table int) string {
	return fmt.Sprintf("Serving Jasmine tea to table #%d", table)
}

// TeaFactory - flyweight factory
type TeaFactory struct {
	teas map[string]TeaType
}

func NewTeaFactory() *TeaFactory {
	return &TeaFactory{
		teas: make(map[string]TeaType),
	}
}

func (tf *TeaFactory) GetTea(teaType string) TeaType {
	if tea, exists := tf.teas[teaType]; exists {
		return tea
	}
	
	var tea TeaType
	switch teaType {
	case "karak":
		tea = KarakTea{}
	case "jasmine":
		tea = JasmineTea{}
	default:
		tea = KarakTea{} // default
	}
	
	tf.teas[teaType] = tea
	fmt.Printf("Created new %s tea type\n", teaType)
	return tea
}

func (tf *TeaFactory) GetCreatedTeaTypesCount() int {
	return len(tf.teas)
}

// TeaShop - context
type TeaShop struct {
	orders  []TeaOrder
	factory *TeaFactory
}

type TeaOrder struct {
	table   int
	teaType string
}

func NewTeaShop() *TeaShop {
	return &TeaShop{
		orders:  make([]TeaOrder, 0),
		factory: NewTeaFactory(),
	}
}

func (ts *TeaShop) TakeOrder(table int, teaType string) {
	ts.orders = append(ts.orders, TeaOrder{table: table, teaType: teaType})
}

func (ts *TeaShop) Serve() {
	for _, order := range ts.orders {
		tea := ts.factory.GetTea(order.teaType)
		fmt.Println(tea.Serve(order.table))
	}
}

func (ts *TeaShop) GetReport() string {
	return fmt.Sprintf("Total orders: %d, Tea types created: %d", 
		len(ts.orders), ts.factory.GetCreatedTeaTypesCount())
}

func main() {
	fmt.Println("=== Flyweight Pattern Demo ===")
	
	shop := NewTeaShop()
	
	// Take multiple orders
	shop.TakeOrder(1, "karak")
	shop.TakeOrder(2, "karak")
	shop.TakeOrder(5, "jasmine")
	shop.TakeOrder(2, "karak")
	shop.TakeOrder(3, "jasmine")
	
	fmt.Println("\nServing orders:")
	shop.Serve()
	
	fmt.Printf("\n%s\n", shop.GetReport())
	
	fmt.Println("\nFlyweight pattern minimizes memory usage by sharing common data!")
}
