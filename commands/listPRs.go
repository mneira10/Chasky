package commands

import (
	"log"
)

type listPRs struct {
}

func (env listPRs) Execute() {
	log.Println("executing azdevops")
}

func (env listPRs) GetOutput() (string, bool) {
	return "some azOuput", false
}
