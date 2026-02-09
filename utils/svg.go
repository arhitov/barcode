package utils

import (
	"fmt"
	"github.com/arhitov/barcode"
)

// MakeSVG Формирует SVG
//
//	dm - данные штрихкода
//	cellSize - размер модуля в пикселях (SVG units per module)
func MakeSVG(dm barcode.Barcode, cellSize int) (string, error) {
	// Точный размер в модулях
	w := dm.Bounds().Dx()
	h := dm.Bounds().Dy()

	// Масштаб 1:size — 1 модуль = size пиксель
	img, err := barcode.Scale(dm, w, h)
	if err != nil {
		return "", err
	}

	svgString := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`+"\n",
		w*cellSize, h*cellSize, w*cellSize, h*cellSize)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			// RGBA() возвращает значения в диапазоне [0, 0xffff]
			if a != 0 && r == 0 && g == 0 && b == 0 { // чёрный модуль
				svgString += fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" fill="black"/>`+"\n",
					x*cellSize, y*cellSize, cellSize, cellSize)
			}
		}
	}

	svgString += `</svg>`

	return svgString, nil
}
