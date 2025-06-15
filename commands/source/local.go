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
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Local(a []string) {
	if len(a) != 2 {
		Help(a)
		return
	}

	width, err := strconv.Atoi(a[0])

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	height, err := strconv.Atoi(a[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	graphic := image.NewRGBA(image.Rect(0, 0, width, height))
	fill := color.RGBA{R: uint8(rand.Int()), G: uint8(rand.Int()), B: uint8(rand.Int()), A: uint8(rand.Int())}

	for x := range width {
		for y := range height {
			graphic.SetRGBA(x, y, fill)
		}
	}

	home, err := os.UserHomeDir()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	path := filepath.Join(home, "graphic.png")
	file, err := os.Create(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	defer file.Close()

	err = png.Encode(file, graphic)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}
}
