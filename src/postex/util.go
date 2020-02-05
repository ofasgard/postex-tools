package postex

import "os"

func Exists(filepath string) error {
	_,err := os.Stat(filepath)
	return err
}
