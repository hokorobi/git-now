package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/mattn/go-pipeline"
)

func main() {
	cmd := exec.Command("git", "add", "-u")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	out, err := pipeline.Output(
		[]string{"git", "diff", "--cached"},
		[]string{"git", "commit", "-F", "-"},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
