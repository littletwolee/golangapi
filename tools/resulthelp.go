package tools

import(
	"log"
	"errors"
)

type Resulthelp struct{}

func checkerr(err error) error {
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func newerr(err string) error {
	return errors.New(err)
}
