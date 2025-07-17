package tetris

type t struct{ tetro }

func newT() Tetro {
	x, y := config.SpawnX, config.SpawnY
	s0, s1, s2, s3 := config.StyleT.Block0, config.StyleT.Block1, config.StyleT.Block2, config.StyleT.Block3

	//   33
	// 110022
	return &t{
		tetro: tetro{
			x:      x,
			y:      y,
			block0: Block{x: x, y: y, style: s0},
			block1: Block{x: x - 1, y: y, style: s1},
			block2: Block{x: x + 1, y: y, style: s2},
			block3: Block{x: x, y: y - 1, style: s3},
		},
	}
}

func (t *t) Rotate() {
	switch t.rotation {
	case 0:
		t.block0.x, t.block0.y = t.x, t.y   //   22
		t.block1.x, t.block1.y = t.x, t.y+1 // 3300
		t.block2.x, t.block2.y = t.x, t.y-1 //   11
		t.block3.x, t.block3.y = t.x-1, t.y
	case 1:
		t.block0.x, t.block0.y = t.x, t.y   // 220011
		t.block1.x, t.block1.y = t.x+1, t.y //   33
		t.block2.x, t.block2.y = t.x-1, t.y
		t.block3.x, t.block3.y = t.x, t.y+1
	case 2:
		t.block0.x, t.block0.y = t.x, t.y   // 11
		t.block1.x, t.block1.y = t.x, t.y-1 // 0033
		t.block2.x, t.block2.y = t.x, t.y+1 // 22
		t.block3.x, t.block3.y = t.x+1, t.y

	case 3:
		t.block0.x, t.block0.y = t.x, t.y   //   33
		t.block1.x, t.block1.y = t.x-1, t.y // 110022
		t.block2.x, t.block2.y = t.x+1, t.y
		t.block3.x, t.block3.y = t.x, t.y-1
	default:
	}
	t.updateRotation()
}

func (original *t) clone() Tetro {
	clone := &t{
		tetro: clone(original.tetro), // implemented in tetris.go
	}
	return clone
}

func (t *t) CanRotate(boardHeight, boardWidth int, stationaryBlocks []Block) bool {
	return canRotate(t, boardHeight, boardWidth, stationaryBlocks) // implemented in tetris.go
}
