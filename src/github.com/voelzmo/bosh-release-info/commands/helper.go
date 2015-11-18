package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func validateAndGetArgument(args cli.Args, length int) (string, error) {
	if len(args) != length {
		return "", fmt.Errorf("Wrong number of arguments. Expected %v, but found %v", length, len(args))
	}
	return args[0], nil
}
