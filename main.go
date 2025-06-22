// genimg is a lightweight tool for generating random images at custom sizes.
// Copyright (C) 2025  Enindu Alahapperuma
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more
// details.
//
// You should have received a copy of the GNU General Public License along with
// this program.  If not, see <https://www.gnu.org/licenses/>.

// genimg is a lightweight tool for generating random images at custom sizes.
//
// Usage:
//
//	genimg <command>:<subcommand> [arguments]
//	genimg [flags]
//
// Available commands:
//
//	source
//
// Available flags:
//
//	-v, --version # View version message
//	-h, --help    # View help message
//
// Use "genimg <command>:help" to see more information about commands.
package main

import (
	"fmt"
	"os"

	"github.com/enindu/genimg/commands/source"
)

func main() {
	dispatchers := map[string]func([]string){
		"source:local":  source.Local,
		"source:picsum": source.Picsum,
		"source:pexels": source.Pexels,
		"source:help":   source.Help,
	}

	inputs := os.Args

	if len(inputs) < 2 {
		fmt.Fprintf(os.Stderr, "%s\n", errNoInstruction.Error())
		return
	}

	instruction := inputs[1]

	switch instruction {
	case "-v", "--version":
		version()
		return
	case "-h", "--help":
		help()
		return
	}

	execute, exists := dispatchers[instruction]
	if !exists {
		fmt.Fprintf(os.Stderr, "%s\n", errInvalidCommand.Error())
		return
	}

	arguments := inputs[2:]

	execute(arguments)
}
