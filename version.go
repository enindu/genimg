// This file is part of genimg.
//
// genimg is free software: you can redistribute it and/or modify it under the
// terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// genimg is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
// A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with
// genimg. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"
)

func version() {
	message := `genimg v1.1.0
Copyright (C) 2025 Enindu Alahapperuma
Licensed under GNU GPL 3.0`

	fmt.Fprintf(os.Stdout, "%s\n", message)
}
