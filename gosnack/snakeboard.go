package gosnack

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type SnakeBoard struct {
	SnakeDirection *Direction
	SnakeSpeed     int
	Snake          []*Point
	Fruit          *Point
	Board          *ebiten.Image
	SnakeBoard     *ebiten.Image
	BoxMargin      int
	BoxSize        int
	BoardSize      int
}

func NewSnakeBoard() *SnakeBoard {
	return &SnakeBoard{
		SnakeDirection: UpDirection,
		SnakeSpeed:     30, // 1s
		Snake:          []*Point{NewPoint(5, 5)},
		Fruit:          NewPoint(5, 4),
		BoxMargin:      1,
		BoxSize:        20,
		BoardSize:      400,
	}
}

func (sb *SnakeBoard) updateSnakeDirection(pressedKey ebiten.Key) {
	switch pressedKey {
	case ebiten.KeyArrowLeft:
		if !RightDirection.IsDirection(sb.SnakeDirection) {
			sb.SnakeDirection = LeftDirection
		}
	case ebiten.KeyArrowUp:
		if !DownDirection.IsDirection(sb.SnakeDirection) {
			sb.SnakeDirection = UpDirection
		}
	case ebiten.KeyArrowRight:
		if !LeftDirection.IsDirection(sb.SnakeDirection) {
			sb.SnakeDirection = RightDirection
		}
	case ebiten.KeyArrowDown:
		if !UpDirection.IsDirection(sb.SnakeDirection) {
			sb.SnakeDirection = DownDirection
		}
	}
}

func (sb *SnakeBoard) Update(totalFrames *int, pressedKey ebiten.Key) (bool, error) {
	sb.updateSnakeDirection(pressedKey)

	if *totalFrames < sb.SnakeSpeed {
		return false, nil
	}

	*totalFrames = 1
	snake := sb.Snake
	last := len(snake) - 1
	snakeHead := snake[last]
	fruit := sb.Fruit
	snakeDic := sb.SnakeDirection
	nx := snakeHead.X + snakeDic.X
	ny := snakeHead.Y + snakeDic.Y
	maxBox := sb.BoardSize / sb.BoxSize

	if nx == fruit.X && ny == fruit.Y {
		sb.Snake = append(snake, NewPoint(nx, ny))
		sb.SnakeSpeed -= 3
		sb.SnakeSpeed = int(math.Max(float64(sb.SnakeSpeed), 5))
		fx := rand.Intn(maxBox - 1)
		fy := rand.Intn(maxBox - 1)
		sb.Fruit = NewPoint(fx, fy)
		return false, nil
	}

	sb.Snake = append(snake, NewPoint(nx, ny))
	isGameOver := sb.IsGameOver()
	if !isGameOver {
		sb.Snake = sb.Snake[1:]
	} else {
		sb.Snake = sb.Snake[:last+1]
	}
	return isGameOver, nil
}

func (sb *SnakeBoard) DrawBoard(ui *ebiten.Image) *ebiten.DrawImageOptions {
	boardOpt := &ebiten.DrawImageOptions{}
	boardOpt.GeoM.Translate(40, 60)

	if sb.Board != nil {
		ui.DrawImage(sb.Board, boardOpt)
		return boardOpt
	}
	margin := sb.BoxMargin
	marginF := float64(margin)
	boxSize := sb.BoxSize
	boardSize := sb.BoardSize
	maxBox := boardSize / boxSize

	sb.Board = ebiten.NewImage(boardSize, boardSize-margin*(maxBox-1))
	sb.Board.Fill(color.RGBA{222, 220, 220, 0xff})

	sb.SnakeBoard = ebiten.NewImage(boardSize, boardSize-margin*(maxBox-1))

	board := sb.Board
	// draw row
	for i := 0; i < maxBox; i++ {
		boxOpt := &ebiten.DrawImageOptions{}
		boxOpt.GeoM.Translate(marginF, float64(margin-i+boxSize*i))
		box := ebiten.NewImage(boardSize-margin*2, boxSize-margin*2)
		box.Fill(color.White)
		board.DrawImage(box, boxOpt)
	}

	// draw col
	for i := 1; i < maxBox; i++ {
		x0 := float32(boxSize*i + margin)
		y0 := float32(margin)
		x1 := x0
		y1 := float32(boardSize - margin*10)
		vector.StrokeLine(board, x0, y0, x1, y1, 1, color.RGBA{222, 220, 220, 0xff}, false)
	}

	ui.DrawImage(sb.Board, boardOpt)
	return boardOpt
}

func (sb *SnakeBoard) CreateBox(x, y int) (*ebiten.Image, *ebiten.DrawImageOptions) {
	boxOpt := &ebiten.DrawImageOptions{}
	fx := x
	fy := y
	box := ebiten.NewImage(sb.BoxSize-1, sb.BoxSize-2)
	boxOpt.GeoM.Translate(
		float64(sb.BoxSize*fx)+float64(sb.BoxMargin),
		float64(sb.BoxSize*fy-fy)+float64(sb.BoxMargin),
	)
	return box, boxOpt
}

func (sb *SnakeBoard) DrawFruit(ui *ebiten.Image) {
	fruit, fruitOpt := sb.CreateBox(sb.Fruit.X, sb.Fruit.Y)
	fruit.Fill(color.RGBA{90, 175, 204, 0xff})
	ui.DrawImage(fruit, fruitOpt)
}

func (sb *SnakeBoard) DrawSnake(ui *ebiten.Image) {
	last := len(sb.Snake) - 1
	for _, snk := range sb.Snake[:last] {
		snake, snakeOpt := sb.CreateBox(snk.X, snk.Y)
		snake.Fill(color.RGBA{250, 147, 147, 0xff})
		ui.DrawImage(snake, snakeOpt)
	}
	snk := sb.Snake[last]
	snake, snakeOpt := sb.CreateBox(snk.X, snk.Y)
	snake.Fill(color.RGBA{156, 149, 149, 0xff})
	ui.DrawImage(snake, snakeOpt)
}

func (sb *SnakeBoard) Draw(ui *ebiten.Image) {
	boardOpt := sb.DrawBoard(ui)
	sb.SnakeBoard.Clear()
	sb.DrawSnake(sb.SnakeBoard)
	sb.DrawFruit(sb.SnakeBoard)
	ui.DrawImage(sb.SnakeBoard, boardOpt)
}

func (sb *SnakeBoard) IsGameOver() bool {
	last := len(sb.Snake) - 1
	snakeHead := sb.Snake[last]
	maxBox := sb.BoardSize / sb.BoxSize
	if snakeHead.X < 0 || snakeHead.Y < 0 || snakeHead.X >= maxBox || snakeHead.Y >= maxBox {
		return true
	}

	for _, snk := range sb.Snake[:last] {
		if snakeHead.IsEqual((snk)) {
			return true
		}
	}

	return false
}
