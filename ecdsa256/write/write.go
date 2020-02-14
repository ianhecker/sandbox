package write

import (
	"io/ioutil"
)

const Dir = "./out/"

func WriteFile(bytes []byte, filename string) error {
	return ioutil.WriteFile(Dir+filename, bytes, 0644)
}
