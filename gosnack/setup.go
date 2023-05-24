package gosnack

import "github.com/hajimehoshi/ebiten/v2"

func SetupWindow() {
	ebiten.SetWindowSize(480, 480)
	ebiten.SetWindowTitle("Snack")
}
