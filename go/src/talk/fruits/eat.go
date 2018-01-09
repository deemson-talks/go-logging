package fruits

import (
	"fmt"
	"github.com/pkg/errors"
)

func Eat(fruit string) (result string, err error) {
	switch fruit {
	case "apple":
		return "good", nil
	case "banana":
		return "nice", nil
	case "pear":
		return "alright", nil
	case "coal":
		return "", errors.New("can't eat coal")
	default:
		panic(fmt.Sprintf(`unknown fruit "%s"`, fruit))
	}
}
