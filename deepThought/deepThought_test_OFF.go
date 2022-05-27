package deepThought

import "testing"

func TestAskPostive(t *testing.T) {
	argumentToTestPostive_1 := "life"
	argumentToTestPostive_2 := "universe"
	argumentToTestPostive_3 := "everything"
	v1, v2 := Ask(argumentToTestPostive_1, argumentToTestPostive_2, argumentToTestPostive_3)
	if v1 != 42 || v2 != "change" {
		t.Error("Expected 42 and change, got ", v1, v2)
	}
}

func TestAskNegative(t *testing.T) {
	argumentToTestPostive_1 := "me"
	argumentToTestPostive_2 := "myself"
	argumentToTestPostive_3 := "I"
	v1, v2 := Ask(argumentToTestPostive_1, argumentToTestPostive_2, argumentToTestPostive_3)
	if v1 != 0 || v2 != "nichts" {
		t.Error("Expected 0 and nichts got ", v1, v2)
	}
}
