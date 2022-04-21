package file

import (
	"mime/multipart"
	"strings"

	"github.com/disintegration/imaging"
)

func Resize(dir string, filename string, file *multipart.FileHeader, resizeRatio float64) (string, error) {
	// Get path
	dir = strings.TrimRight(dir, "/") + "/"
	path := dir + filename

	// Open image
	src, err := imaging.Open(path, imaging.AutoOrientation(true))
	if err != nil {
		return "", err
	}

	// Resize ratio
	if resizeRatio <= 0 {
		resizeRatio = getResizeRatio(file)
	}
	width := float64(src.Bounds().Size().X) * resizeRatio

	// Resize
	src = imaging.Resize(src, int(width), 0, imaging.Lanczos)
	err = imaging.Save(src, path)
	if err != nil {
		return "", err
	}

	return path, nil
}

func getResizeRatio(file *multipart.FileHeader) float64 {
	// < 100k
	if file.Size < 1024*100 {
		return 1
	}

	// 100k - 300k
	if file.Size <= 1024*300 {
		return 0.8
	}

	// 300k - 500k
	if file.Size <= 1024*500 {
		return 0.6
	}

	// 500k - 1M
	if file.Size <= 1024*1024 {
		return 0.5
	}

	// 1M - 5M
	if file.Size <= 1024*1024*5 {
		return 0.3
	}

	// 5M - 10M
	if file.Size <= 1024*1024*10 {
		return 0.2
	}

	// > 5M
	return 0.1
}
