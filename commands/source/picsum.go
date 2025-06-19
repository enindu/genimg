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
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Picsum(a []string) {
	// Validate arguments and get parameters
	if len(a) != 2 {
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

	// Create request, send request, and get response
	url := fmt.Sprintf("https://picsum.photos/%d/%d", width, height)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	defer response.Body.Close()

	// Create file
	directory, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	path := filepath.Join(directory, "picsum.jpg")

	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	// Copy response body into the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	// Print message
	fmt.Fprintf(os.Stdout, "Image saved in \"%s\"\n", path)
}
