package main

import (
	"fmt"
	"log"
	"smallgamekit/sgk"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	textRenderer sgk.TextRenderer
}

func (r *Game) Update() error {
	return nil
}

func (r *Game) Draw(screen *ebiten.Image) {
	r.textRenderer.DrawText(screen, "hello", 10, 12, 4)
}

func (r *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 200, 200
}

func main() {
	game := &Game{
		textRenderer: *sgk.NewTextRenderer(),
	}
	ebiten.SetWindowSize(200, 200)
	ebiten.SetWindowTitle("Total")
	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}

	ir := sgk.NewImageResources("assets/images/")

	im := ir.GetImage("example")
	fmt.Printf("image color: %v", im)
}
