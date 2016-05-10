package tools

import(
	"log"
	"errors"
)

type ResultHelp struct{}

func (r *ResultHelp)CheckErr(err error) error {
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *ResultHelp)NewErr(err string) error {
	return errors.New(err)
}
