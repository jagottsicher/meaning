package deepThought

import (
	"testing"
)

func TestAskTableTest(t *testing.T) {

	type testData struct {
		inputStrings []string
		result1      int8
		result2      string
	}

	tests := []testData{
		testData{[]string{"life", "universe", "everything"}, 42, "change"},
		testData{[]string{"universe", "life", "everything"}, 0, "nothing"},
		testData{[]string{"me", "myself", "I"}, 0, "nothing"},
	}

	for _, v := range tests {
		v1, v2 := Ask(v.inputStrings[0], v.inputStrings[1], v.inputStrings[2])
		if v1 != v.result1 || v2 != v.result2 {
			t.Error("Expected", v.result1, "and", v.result2, "got", v1, "and", v2)
		}
	}
}
