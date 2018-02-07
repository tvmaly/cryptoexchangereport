package parsing

import (
	"io/ioutil"
)

func NewFileSource(filename string) ([]byte, error) {

	rawbytes, err := ioutil.ReadFile(filename)

	return rawbytes, err

}
