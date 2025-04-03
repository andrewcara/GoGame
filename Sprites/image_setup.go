package Sprites

import (
	"log"
	"path/filepath"

	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func CreateImage(image_name string) *ebiten.Image {
	var err error
	file := filepath.Join("Assets", image_name)
	img, _, err := ebitenutil.NewImageFromFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
