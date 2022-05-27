package deepThought

import "fmt"

func ExampleAsk() {
	v1, v2 := Ask("life", "universe", "everything")
	fmt.Println(v1, v2)
	// Output:
	// What is the answer to the ultimate question of life, the universe, and everything?
	// 42 change
}
