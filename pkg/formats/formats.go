package formats

import (
	"fmt"
	"github.com/Shackelford-Arden/BookBrowser/pkg/booklist"
	"image"
	"path"
	"strings"
)

var supportedFormats = map[string]func(filename string) (BookInfo, error){}

type BookInfo interface {
	Book() *booklist.Book
	HasCover() bool
	GetCover() (image.Image, error)
}

func Register(ext string, load func(filename string) (BookInfo, error)) {
	ext = strings.ToLower(ext)
	if _, ok := supportedFormats[ext]; ok {
		panic("attempted to register existing format " + ext)
	}
	supportedFormats[ext] = load
}

func Load(filename string) (BookInfo, error) {
	ext := strings.Replace(path.Ext(filename), ".", "", 1)
	load, ok := supportedFormats[strings.ToLower(ext)]
	if !ok {
		return nil, fmt.Errorf("could not load format %s", ext)
	}
	return load(filename)
}

func GetExts() []string {
	var exts []string
	for ext := range supportedFormats {
		exts = append(exts, ext)
	}
	return exts
}
