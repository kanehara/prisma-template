// +build ignore

package main

import (
	"os"

	"github.com/99designs/gqlgen/cmd"
)

func main() {
	os.RemoveAll("./gqlgen/tmp")
	cmd.Execute()
}
