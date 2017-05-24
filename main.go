package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/smallfish/simpleyaml"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: yaml-hook filename.yaml")
	} else {
		filename := os.Args[1]
		source, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
		}
		_, err = simpleyaml.NewYaml(source)
		if err != nil {
			fmt.Printf("Invalid yaml detected in %s:\n %s\n", filename, err)
			os.Exit(2)
		}
	}
}
