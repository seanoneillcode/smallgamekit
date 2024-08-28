package sgk

import (
	"bytes"
	"image"
	"log"
	"os"
	"path"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"

	_ "image/png" // evil, required for decoder to 'know' what a png is...
)

type ImageResources struct {
	images map[string]*ebiten.Image
}

// path to root folder containing images
func NewImageResources(folderPath string) *ImageResources {
	images := map[string]*ebiten.Image{}
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatalf("failed to read dir: %v", err)
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".png") {
			key, _ := strings.CutSuffix(e.Name(), ".png")
			images[key] = loadImage(path.Join(folderPath, e.Name()))
		}
	}
	return &ImageResources{
		images: images,
	}
}

func loadImage(imageFileName string) *ebiten.Image {
	b, err := os.ReadFile(imageFileName)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatalf("failed to decode file: %v", err)
	}
	return ebiten.NewImageFromImage(img)
}

func (r *ImageResources) GetImage(id string) *ebiten.Image {
	image, ok := r.images[id]
	if !ok {
		log.Fatalf("image not found: %s", id)
	}
	return image
}
