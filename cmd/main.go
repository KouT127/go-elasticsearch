package main

import (
	"go-elasticsearch/cmd/commands"
)

func main() {
	cmd := commands.NewRootCmd()
	cmd.Execute()
}
