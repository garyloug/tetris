package tetris

type l struct{ tetro }

func newL() Tetro {
	x, y := config.SpawnX, config.SpawnY
	style := config.StyleL

	// 22
	// 11
	// 0033
	return &l{
		tetro: tetro{
			x:      x,
			y:      y,
			block0: Block{x: x, y: y, style: style},
			block1: Block{x: x, y: y - 1, style: style},
			block2: Block{x: x, y: y - 2, style: style},
			block3: Block{x: x + 1, y: y, style: style},
		},
	}
}

func (l *l) Rotate() {
	switch l.rotation {
	case 0:
		l.block0.x, l.block0.y = l.x, l.y   //     33
		l.block1.x, l.block1.y = l.x-1, l.y // 221100
		l.block2.x, l.block2.y = l.x-2, l.y
		l.block3.x, l.block3.y = l.x, l.y-1
	case 1:
		l.block0.x, l.block0.y = l.x, l.y   // 3300
		l.block1.x, l.block1.y = l.x, l.y+1 //   11
		l.block2.x, l.block2.y = l.x, l.y+2 //   22
		l.block3.x, l.block3.y = l.x-1, l.y
	case 2:
		l.block0.x, l.block0.y = l.x, l.y   // 001122
		l.block1.x, l.block1.y = l.x+1, l.y // 33
		l.block2.x, l.block2.y = l.x+2, l.y
		l.block3.x, l.block3.y = l.x, l.y+1
	case 3:
		l.block0.x, l.block0.y = l.x, l.y   // 22
		l.block1.x, l.block1.y = l.x, l.y-1 // 11
		l.block2.x, l.block2.y = l.x, l.y-2 // 0033
		l.block3.x, l.block3.y = l.x+1, l.y
	default:
	}
	l.updateRotation()
}

func (original *l) clone() Tetro {
	clone := &l{
		tetro: clone(original.tetro), // implemented in tetris.go
	}
	return clone
}

func (l *l) CanRotate(boardHeight, boardWidth int, stationaryBlocks []Block) bool {
	return canRotate(l, boardHeight, boardWidth, stationaryBlocks) // implemented in tetris.go
}
