package gosnack

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Game struct {
	Width  int
	Height int
	Score  int
	Frames int
	State  GameState
	Font   font.Face
	Board  *SnakeBoard
	UI     *ebiten.Image
}

func NewGame() *Game {
	g := &Game{
		Width:  480,
		Height: 480,
		Score:  0,
		Frames: 0,
		State:  RUNNING,
		Board:  NewSnakeBoard(),
		Font:   GetDefaultFont(),
		UI:     ebiten.NewImage(480, 480),
	}
	return g
}

func (g *Game) Update() error {
	if g.State == OVER {
		return nil
	}
	g.Frames += 1
	pressedKeys := inpututil.AppendPressedKeys(nil)
	pressedKey := ebiten.KeyF
	if len(pressedKeys) != 0 {
		pressedKey = pressedKeys[0]
	}

	err := g.Board.Update(&g.Frames, pressedKey)
	if g.Board.IsGameOver() {
		g.State = OVER
	}
	return err
}

func (g *Game) DrawGameText(ui *ebiten.Image) {
	if g.State == 1 {
		text.Draw(ui, "Snack Game!", g.Font, g.Width/3, 36, color.Black)
	}
	if g.State == 0 {
		text.Draw(ui, "Game Over!", g.Font, g.Width/3, 36, color.Black)
	}
}

func (g *Game) DrawFPS(ui *ebiten.Image) {
	fps := ebiten.ActualFPS()
	ebitenutil.DebugPrint(ui, fmt.Sprintf("FPS: %f", fps))
}

func (g *Game) DrawUI(screen *ebiten.Image) {
	g.UI.Clear()
	g.UI.Fill(color.White)
	g.DrawGameText(g.UI)
	g.DrawFPS(g.UI)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawUI(screen)
	g.Board.Draw(g.UI)
	screen.DrawImage(g.UI, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Width, g.Height
}

func RunGame() {
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
