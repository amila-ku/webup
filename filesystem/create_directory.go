package filesystem

import (
	"log"
	"os"
)

const (
	fileMode = 0755
)

// CreateDirectory is to create a new directory as per the input name
func CreateDirectory(dirname string) error {
	err := os.Mkdir(dirname, fileMode)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
		return err
	}

	err = os.WriteFile(dirname+"/index.html", []byte("Hello, Gophers!"), 0660)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
