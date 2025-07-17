package tetris

type o struct{ tetro }

func newO() Tetro {
	x, y := config.SpawnX, config.SpawnY
	s0, s1, s2, s3 := config.StyleO.Block0, config.StyleO.Block1, config.StyleO.Block2, config.StyleO.Block3

	// 0033
	// 1122
	return &o{
		tetro: tetro{
			x:      x,
			y:      y,
			block0: Block{x: x, y: y, style: s0},
			block1: Block{x: x, y: y + 1, style: s1},
			block2: Block{x: x + 1, y: y + 1, style: s2},
			block3: Block{x: x + 1, y: y, style: s3},
		},
	}
}

func (o *o) Rotate() {
	tlx, tly := o.x, o.y
	blx, bly := o.x, o.y+1
	brx, bry := o.x+1, o.y+1
	trx, try := o.x+1, o.y

	switch o.rotation {
	case 0:
		o.block0.x, o.block0.y = blx, bly // 3322
		o.block1.x, o.block1.y = brx, bry // 0011
		o.block2.x, o.block2.y = trx, try
		o.block3.x, o.block3.y = tlx, tly
	case 1:
		o.block0.x, o.block0.y = brx, bry // 2211
		o.block1.x, o.block1.y = trx, try // 3300
		o.block2.x, o.block2.y = tlx, tly
		o.block3.x, o.block3.y = blx, bly
	case 2:
		o.block0.x, o.block0.y = trx, try // 1100
		o.block1.x, o.block1.y = tlx, tly // 2233
		o.block2.x, o.block2.y = blx, bly
		o.block3.x, o.block3.y = brx, bry
	case 3:
		o.block0.x, o.block0.y = tlx, tly // 0033
		o.block1.x, o.block1.y = blx, bly // 1122
		o.block2.x, o.block2.y = brx, bry
		o.block3.x, o.block3.y = trx, try
	default:
	}
	o.updateRotation()
}

func (original *o) clone() Tetro {
	clone := &o{
		tetro: clone(original.tetro), // implemented in tetris.go
	}
	return clone
}

func (o *o) CanRotate(boardHeight, boardWidth int, stationaryBlocks []Block) bool {
	return canRotate(o, boardHeight, boardWidth, stationaryBlocks) // implemented in tetris.go
}
