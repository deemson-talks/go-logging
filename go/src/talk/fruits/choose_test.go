package fruits

import (
	"testing"
)

func TestChooseRandom(t *testing.T) {
	choices := map[string]int{}
	for i := 1; i <= 100000; i++ {
		fruit := ChooseRandom()
		num, ok := choices[fruit]
		if ok {
			choices[fruit] = num + 1
		} else {
			choices[fruit] = 1
		}
	}
	t.Log(choices)
}
