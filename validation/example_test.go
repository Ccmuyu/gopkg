package validation

import "fmt"

func ExampleIsUUID() {
	fmt.Println(IsUUID("550e8400-e29b-41d4-a716-446655440000"))
	fmt.Println(IsUUID("not-a-uuid"))
	// Output:
	// true
	// false
}

func ExampleIsCreditCard() {
	fmt.Println(IsCreditCard("4111111111111111"))
	fmt.Println(IsCreditCard("1234567890123456"))
	// Output:
	// true
	// false
}

func ExampleIsJSON() {
	fmt.Println(IsJSON(`{"key":"value"}`))
	fmt.Println(IsJSON(`{invalid}`))
	// Output:
	// true
	// false
}
