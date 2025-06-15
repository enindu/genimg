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

import "errors"

var (
	errNoInstruction  error = errors.New("the instruction is not found, use \"genimg -h\" or \"genimg --help\" to see help message")
	errInvalidCommand error = errors.New("the command is invalid, use \"genimg -h\" or \"genimg --help\" to see help message")
)
