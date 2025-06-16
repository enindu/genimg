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

package source

import (
	"fmt"
	"os"
	"strings"
)

func Picsum(a []string) {
	if len(a) != 2 {
		help()
		return
	}

	width := a[0]
	height := a[1]

	path, err := saveFile(fmt.Sprintf("https://picsum.photos/%s/%s", width, height), "picsum.jpg")

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	fmt.Fprintf(os.Stdout, "Image saved in \"%s\"\n", path)
}
