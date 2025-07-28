package tetris

import (
	"math/rand"
)

var (
	config         Config
	tetroFactories = []func() Tetro{newO, newI, newS, newZ, newL, newJ, newT}
)

type Tetro interface {
	Blocks() []Block
	MoveRight()
	MoveLeft()
	MoveDown()
	Rotate()
	CanMoveDown(boardHeight, boardWidth int, stationaryBlocks []Block) bool
	CanMoveRight(boardHeight, boardWidth int, stationaryBlocks []Block) bool
	CanMoveLeft(boardHeight, boardWidth int, stationaryBlocks []Block) bool
	CanRotate(boardHeight, boardWidth int, stationaryBlocks []Block) bool
	clone() Tetro
}

type tetro struct {
	x        int
	y        int
	rotation int
	block0   Block
	block1   Block
	block2   Block
	block3   Block
}

type BlockStyle any

type Config struct {
	SpawnX int
	SpawnY int
	StyleO BlockStyle
	StyleI BlockStyle
	StyleS BlockStyle
	StyleZ BlockStyle
	StyleL BlockStyle
	StyleJ BlockStyle
	StyleT BlockStyle
}

func Init(cfg Config) {
	config = cfg
}

func NewRandomTetro() Tetro {
	return tetroFactories[rand.Intn(len(tetroFactories))]()
}

func (t *tetro) Blocks() []Block {
	return []Block{
		t.block0,
		t.block1,
		t.block2,
		t.block3,
	}
}

func (t *tetro) updateRotation() {
	t.rotation++
	if t.rotation > 3 {
		t.rotation = 0
	}
}

func (t *tetro) MoveRight() {
	t.x++
	t.block0.moveRight()
	t.block1.moveRight()
	t.block2.moveRight()
	t.block3.moveRight()
}

func (t *tetro) MoveLeft() {
	t.x--
	t.block0.moveLeft()
	t.block1.moveLeft()
	t.block2.moveLeft()
	t.block3.moveLeft()
}

func (t *tetro) MoveDown() {
	t.y++
	t.block0.MoveDown()
	t.block1.MoveDown()
	t.block2.MoveDown()
	t.block3.MoveDown()
}

func (t *tetro) CanMoveDown(boardHeight, boardWidth int, stationaryBlocks []Block) bool {
	for _, block := range t.Blocks() {
		x, y := block.Coordinates()
		if y >= boardHeight-1 {
			return false
		}

		for _, fallenBlock := range stationaryBlocks {
			fallenX, fallenY := fallenBlock.Coordinates()
			if x == fallenX && y+1 == fallenY {
				return false
			}
		}

	}
	return true
}

func (t *tetro) CanMoveRight(boardHeight, boardWidth int, stationaryBlocks []Block) bool {
	for _, block := range t.Blocks() {
		x, y := block.Coordinates()
		if x >= boardWidth-1 {
			return false
		}

		for _, fallenBlock := range stationaryBlocks {
			fallenX, fallenY := fallenBlock.Coordinates()
			if x+1 == fallenX && y == fallenY {
				return false
			}
		}
	}
	return true
}

func (t *tetro) CanMoveLeft(boardHeight, boardWidth int, stationaryBlocks []Block) bool {
	for _, block := range t.Blocks() {
		x, y := block.Coordinates()
		if x <= 0 {
			return false
		}

		for _, fallenBlock := range stationaryBlocks {
			fallenX, fallenY := fallenBlock.Coordinates()
			if x-1 == fallenX && y == fallenY {
				return false
			}
		}
	}
	return true
}

// CanRotate and Clone methods are implemented in each shape file, as each shape has unique rotation logic.
// These canRotate and clone functions are defined here to avoid some code duplication in the shapes.
func canRotate(tetro Tetro, boardHeight, boardWidth int, stationaryBlocks []Block) bool {
	clone := tetro.clone()
	clone.Rotate()
	for _, block := range clone.Blocks() {
		x, y := block.Coordinates()
		if x < 0 || x >= boardWidth || y >= boardHeight {
			return false
		}

		for _, fallenBlock := range stationaryBlocks {
			fallenX, fallenY := fallenBlock.Coordinates()
			if x == fallenX && y == fallenY {
				return false
			}
		}
	}
	return true
}

// Note that this private clone function returns a tetro type, not a Tetro interface.
// The clone method defined in each shape returns an actual Tetro, which includes a Rotate method.
// The Rotate method is used in canRotate above
func clone(original tetro) tetro {
	return tetro{
		x:        original.x,
		y:        original.y,
		rotation: original.rotation,
		block0:   Block{x: original.block0.x, y: original.block0.y, style: original.block0.style},
		block1:   Block{x: original.block1.x, y: original.block1.y, style: original.block1.style},
		block2:   Block{x: original.block2.x, y: original.block2.y, style: original.block2.style},
		block3:   Block{x: original.block3.x, y: original.block3.y, style: original.block3.style},
	}
}
