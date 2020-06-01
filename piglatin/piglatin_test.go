package piglatin

import "testing"

func TestEncode(t *testing.T) {
	testTable := map[string]string{
		"smile": "ilesmay",
		"string": "ingstray",
		"stupid": "upidstay",
		"glove": "oveglay",
		"trash": "ashtray",
		"floor": "oorflay",
		"store": "orestay",
		"explain": "explainyay",
		"I": "Iyay",
	}

	for input, expected := range testTable {
		if output := Encode(input); output != expected {
			t.Errorf("Expected %v to be %v", output, expected)
		}
	}
}