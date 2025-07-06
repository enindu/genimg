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
	"image/jpeg"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Local(a []string) {
	// Validate arguments and get parameters
	if len(a) != 3 {
		Help(nil)
		return
	}

	width, err := strconv.Atoi(a[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	if width < 1 {
		fmt.Fprintf(os.Stderr, "%s\n", errInvalidWidth.Error())
		return
	}

	height, err := strconv.Atoi(a[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	if height < 1 {
		fmt.Fprintf(os.Stderr, "%s\n", errInvalidHeight.Error())
		return
	}

	count, err := strconv.Atoi(a[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	if count < 1 {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(errInvalidCount.Error()))
		return
	}

	for i := range count {
		// Create graphic
		rectangle := image.Rect(0, 0, width, height)
		graphic := image.NewRGBA(rectangle)

		fill := color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: uint8(rand.Intn(256)),
		}

		for x := range width {
			for y := range height {
				graphic.SetRGBA(x, y, fill)
			}
		}

		// Create file
		directory, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
			return
		}

		name := fmt.Sprintf("local-%d.jpg", i+1)
		path := filepath.Join(directory, name)

		file, err := os.Create(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
			return
		}

		// Encode graphic into the file
		err = jpeg.Encode(file, graphic, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
			return
		}

		// Print message
		fmt.Fprintf(os.Stdout, "Image saved in \"%s\"\n", path)
	}
}
