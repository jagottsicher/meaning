// deepThough stellt die Antwort auf die Fragen nach dem Leben, dem Universum und allem bereit
package deepThought

import "fmt"

// question ist ein Typ mit unterliegendem struct, der drei Elemente vom Typ string beinhaltet
type question struct {
	thing1 string
	thing2 string
	thing3 string
}

// InLife ist die einzige Konstante im Leben
const InLife string = "change"

// Uqlue ist eine Variable, die die strings festhält, nach denen gefragt werden kann.
var Uqlue question

// Ask prüft, ob in der Frage enthaltene Elemente mit gespeicherten Elementen identisch sind
//
// Es wird auf die Frage nach Leben, Universum und allem 42 zurückgegeben, in allen anderen Fällen 0.
func Ask(life, universe, everything string) (int8, string) {

	// Uqlue ist eine Variable vom Typ question, die drei Elemente vom Typ String in einem Struct bereit hält
	Uqlue = question{
		thing1: "life",
		thing2: "universe",
		thing3: "everything",
	}

	fmt.Printf("What is the answer to the ultimate question of %v, the %v, and %v?\n", life, universe, everything)

	if life == Uqlue.thing1 && universe == Uqlue.thing2 && everything == Uqlue.thing3 {
		return 42, InLife
	}
	return 0, "nichts"
}
