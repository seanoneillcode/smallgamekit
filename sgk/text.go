package sgk

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	_ "embed"
)

//go:embed text-source.png
var textSource []byte

type TextRenderer struct {
	textImage           *ebiten.Image
	textCharacterImages map[rune]*ebiten.Image
}

func NewTextRenderer() *TextRenderer {
	img, _, err := image.Decode(bytes.NewReader(textSource))
	if err != nil {
		log.Fatalf("failed to decode file: %v", err)
	}
	return &TextRenderer{
		textImage:           ebiten.NewImageFromImage(img),
		textCharacterImages: map[rune]*ebiten.Image{},
	}
}

func (tr *TextRenderer) DrawText(screen *ebiten.Image, str string, ox, oy int, scale float64) {
	op := &ebiten.DrawImageOptions{}
	x := 0
	y := 0
	const (
		cw = 10
		ch = 12
	)
	for _, c := range str {
		if c == '\n' {
			x = 0
			y += ch
			continue
		}
		s, ok := tr.textCharacterImages[c]
		if !ok {
			cval := int(c)
			index := -1
			if cval > 96 && cval < 123 {
				index = int(c) - 97
			}
			if cval > 47 && cval < 59 {
				index = int(c) - 48 + 26 // the width of the preceding letters
			}
			if c == ',' {
				index = 36
			}
			if c == '.' {
				index = 37
			}
			if c == '!' {
				index = 38
			}
			if c == '?' {
				index = 39
			}
			if c == ' ' {
				x += cw - 5
			}
			if index != -1 {
				sx := index * cw
				rect := image.Rect(sx, 0, sx+cw-1, ch-1)
				s = tr.textImage.SubImage(rect).(*ebiten.Image)
				tr.textCharacterImages[c] = s
			}
		}
		if s != nil {
			op.GeoM.Reset()
			op.GeoM.Translate(float64(ox), float64(oy))
			op.GeoM.Translate(float64(x), float64(y))
			op.GeoM.Scale(scale, scale)
			screen.DrawImage(s, op)
			x += cw - 4
		}
	}
}
