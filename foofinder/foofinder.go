package foofinder

import "errors"

func IsItFoo(word string) (bool, error) {
	switch word {
	case "fooZ":
		return true, nil
	case "bar":
		err := errors.New("Oh noooooo, it's barZ!")
		return false, err
	}
	return false, nil
}
