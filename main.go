package main

import (
	"deepThought"
	"fmt"
)

// main hilft uns Konstanten im Leben zu finden und wichtigsten Fragen im Leben zu beantworten
func main() {
	answer, onlyConstant := deepThought.Ask("life", "universe", "everything")
	fmt.Printf("Press [Enter] to see the answer!")
	fmt.Scanln()
	fmt.Printf("The answer is %v.\n", answer)
	fmt.Println("BTW: The only constant in life is", onlyConstant)
}
