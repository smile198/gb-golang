package fibonacci

import "fmt"

func ExampleFibonacciFinder_GetNumber() {
	ff := New(20)

	fmt.Println(ff.GetNumber(10))
	fmt.Println(ff.GetNumber(20))

	// Output:
	//  6765
	//  55
}
