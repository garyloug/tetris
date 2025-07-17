package tetris

type i struct{ tetro }

func newI() Tetro {
	x, y := config.SpawnX, config.SpawnY
	s0, s1, s2, s3 := config.StyleI.Block0, config.StyleI.Block1, config.StyleI.Block2, config.StyleI.Block3

	// 33001122
	return &i{
		tetro: tetro{
			x:      x,
			y:      y,
			block0: Block{x: x, y: y, style: s0},
			block1: Block{x: x + 1, y: y, style: s1},
			block2: Block{x: x + 2, y: y, style: s2},
			block3: Block{x: x - 1, y: y, style: s3},
		},
	}
}

func (i *i) Rotate() {
	switch i.rotation {
	case 0:
		i.block0.x, i.block0.y = i.x, i.y   // 22
		i.block1.x, i.block1.y = i.x, i.y-1 // 11
		i.block2.x, i.block2.y = i.x, i.y-2 // 00
		i.block3.x, i.block3.y = i.x, i.y+1 // 33
	case 1:
		i.block0.x, i.block0.y = i.x, i.y // 22110033
		i.block1.x, i.block1.y = i.x-1, i.y
		i.block2.x, i.block2.y = i.x-2, i.y
		i.block3.x, i.block3.y = i.x+1, i.y
	case 2:
		i.block0.x, i.block0.y = i.x, i.y   // 33
		i.block1.x, i.block1.y = i.x, i.y+1 // 00
		i.block2.x, i.block2.y = i.x, i.y+2 // 11
		i.block3.x, i.block3.y = i.x, i.y-1 // 22
	case 3:
		i.block0.x, i.block0.y = i.x, i.y // 33001122
		i.block1.x, i.block1.y = i.x+1, i.y
		i.block2.x, i.block2.y = i.x+2, i.y
		i.block3.x, i.block3.y = i.x-1, i.y
	default:
	}
	i.updateRotation()
}

func (original *i) clone() Tetro {
	clone := &i{
		tetro: clone(original.tetro), // implemented in tetris.go
	}
	return clone
}

func (i *i) CanRotate(boardHeight, boardWidth int, stationaryBlocks []Block) bool {
	return canRotate(i, boardHeight, boardWidth, stationaryBlocks) // implemented in tetris.go
}
