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
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func saveFile(u string, n string) (string, error) {
	request, err := http.NewRequest(http.MethodGet, u, nil)

	if err != nil {
		return "", err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	file, path, err := createFile(n)

	if err != nil {
		return "", err
	}

	_, err = io.Copy(file, response.Body)

	if err != nil {
		return "", err
	}

	return path, nil
}

func createFile(n string) (*os.File, string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return nil, "", err
	}

	path := filepath.Join(home, n)

	file, err := os.Create(path)

	if err != nil {
		return nil, "", err
	}

	return file, path, nil
}
