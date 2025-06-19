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
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/image/draw"
)

func Pexels(a []string) {
	// Validate arguments and get parameters
	if len(a) != 4 {
		Help(nil)
		return
	}

	key := a[0]

	width, err := strconv.Atoi(a[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	if width < 1 {
		fmt.Fprintf(os.Stderr, "%s\n", errInvalidWidth.Error())
		return
	}

	height, err := strconv.Atoi(a[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	if height < 1 {
		fmt.Fprintf(os.Stderr, "%s\n", errInvalidHeight.Error())
		return
	}

	keyword := a[3]

	// Create request, send request, and get response of search
	searchURL := fmt.Sprintf("https://api.pexels.com/v1/search?query=%s&per_page=80", keyword)

	searchRequest, err := http.NewRequest(http.MethodGet, searchURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	searchRequest.Header.Set("Authorization", key)

	searchResponse, err := http.DefaultClient.Do(searchRequest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	defer searchResponse.Body.Close()

	// Convert response body into bytes
	bytes, err := io.ReadAll(searchResponse.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	// Parse JSON data
	type Photo struct {
		ID              int               `json:"id"`
		Width           int               `json:"width"`
		Height          int               `json:"height"`
		URL             string            `json:"url"`
		Photographer    string            `json:"photographer"`
		PhotographerURL string            `json:"photographer_url"`
		PhotographerID  int               `json:"photographer_id"`
		AvgColor        string            `json:"avg_color"`
		Src             map[string]string `json:"src"`
		Liked           bool              `json:"liked"`
		Alt             string            `json:"alt"`
	}

	type Result struct {
		TotalResults int     `json:"total_results"`
		Page         int     `json:"page"`
		PerPage      int     `json:"per_page"`
		Photos       []Photo `json:"photos"`
		NextPage     string  `json:"next_page"`
	}

	var result Result

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	// Create request, send request, and get response of photo
	random := rand.Intn(81)
	photoURL := result.Photos[random].Src["original"]

	photoRequest, err := http.NewRequest(http.MethodGet, photoURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	photoResponse, err := http.DefaultClient.Do(photoRequest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	defer photoResponse.Body.Close()

	// Decode response body into input graphic
	inputGraphic, _, err := image.Decode(photoResponse.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	// Get input graphic properties
	inputGraphicWidth := inputGraphic.Bounds().Dx()
	inputGraphicHeight := inputGraphic.Bounds().Dy()
	inputGraphicAspectRatio := float64(inputGraphicWidth) / float64(inputGraphicHeight)

	// Get resize graphic properties
	resizeGraphicAspectRatio := float64(width) / float64(height)

	// Get crop graphic properties
	cropGraphicWidth := inputGraphicWidth
	cropGraphicHeight := inputGraphicHeight

	if inputGraphicAspectRatio > resizeGraphicAspectRatio {
		cropGraphicWidth = int(float64(cropGraphicHeight) * resizeGraphicAspectRatio)
	} else {
		cropGraphicHeight = int(float64(cropGraphicWidth) / resizeGraphicAspectRatio)
	}

	// Create crop graphic
	left := (inputGraphicWidth - cropGraphicWidth) / 2
	top := (inputGraphicHeight - cropGraphicHeight) / 2
	right := left + cropGraphicWidth
	bottom := top + cropGraphicHeight
	cropRectangle := image.Rect(left, top, right, bottom)
	cropGraphic := image.NewRGBA(cropRectangle)

	draw.Draw(cropGraphic, cropGraphic.Bounds(), inputGraphic, cropRectangle.Min, draw.Src)

	// Create resize graphic
	resizeRectangle := image.Rect(0, 0, width, height)
	resizeGraphic := image.NewRGBA(resizeRectangle)

	draw.CatmullRom.Scale(resizeGraphic, resizeGraphic.Bounds(), cropGraphic, cropGraphic.Bounds(), draw.Over, nil)

	// Create file
	directory, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	path := filepath.Join(directory, "pexels.jpg")

	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	// Encode resize graphic into file
	err = jpeg.Encode(file, resizeGraphic, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	// Print message
	fmt.Fprintf(os.Stdout, "Image saved in \"%s\"\n", path)
}
