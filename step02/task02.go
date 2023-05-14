package step02

import (
	"log"
	"os"
	"os/exec"
)

func cleanup() {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalln("unable to restore cooked mode: ", err)
	}
}
