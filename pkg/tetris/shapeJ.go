package tetris

type j struct{ tetro }

func newJ() Tetro {
	x, y := config.SpawnX, config.SpawnY
	style := config.StyleJ

	//   22
	//   11
	// 3300
	return &j{
		tetro: tetro{
			x:      x,
			y:      y,
			block0: Block{x: x, y: y, style: style},
			block1: Block{x: x, y: y - 1, style: style},
			block2: Block{x: x, y: y - 2, style: style},
			block3: Block{x: x - 1, y: y, style: style},
		},
	}
}

func (j *j) Rotate() {
	switch j.rotation {
	case 0:
		j.block0.x, j.block0.y = j.x, j.y   // 221100
		j.block1.x, j.block1.y = j.x-1, j.y //     33
		j.block2.x, j.block2.y = j.x-2, j.y
		j.block3.x, j.block3.y = j.x, j.y+1
	case 1:
		j.block0.x, j.block0.y = j.x, j.y   // 0033
		j.block1.x, j.block1.y = j.x, j.y+1 // 11
		j.block2.x, j.block2.y = j.x, j.y+2 // 22
		j.block3.x, j.block3.y = j.x+1, j.y
	case 2:
		j.block0.x, j.block0.y = j.x, j.y   // 33
		j.block1.x, j.block1.y = j.x+1, j.y // 001122
		j.block2.x, j.block2.y = j.x+2, j.y
		j.block3.x, j.block3.y = j.x, j.y-1
	case 3:
		j.block0.x, j.block0.y = j.x, j.y   //   22
		j.block1.x, j.block1.y = j.x, j.y-1 //   11
		j.block2.x, j.block2.y = j.x, j.y-2 // 3300
		j.block3.x, j.block3.y = j.x-1, j.y
	default:
	}
	j.updateRotation()
}

func (original *j) clone() Tetro {
	clone := &j{
		tetro: clone(original.tetro), // implemented in tetris.go
	}
	return clone
}

func (j *j) CanRotate(boardHeight, boardWidth int, stationaryBlocks []Block) bool {
	return canRotate(j, boardHeight, boardWidth, stationaryBlocks) // implemented in tetris.go
}
