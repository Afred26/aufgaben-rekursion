package strings

import "fmt"

func ExampleLength() {

	fmt.Println(Length("abc"))
	fmt.Println(Length("Hallo"))
	fmt.Println(Length(""))
	fmt.Println(Length("Rekursion"))

	//Output:
	//3
	//5
	//0
	//9
}
