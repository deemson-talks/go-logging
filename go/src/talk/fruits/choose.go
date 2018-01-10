package fruits

import (
	"github.com/jmcvetta/randutil"
)

func ChooseRandom() string {
	c, err := randutil.WeightedChoice([]randutil.Choice{
		{100, "apple"},
		{70, "banana"},
		{70, "pear"},
		{20, "coal"},
	})
	if err != nil {
		panic(err)
	}
	return c.Item.(string)
}
