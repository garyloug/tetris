package tetris

type z struct{ tetro }

func newZ() Tetro {
	x, y := config.SpawnX, config.SpawnY
	s0, s1, s2, s3 := config.StyleZ.Block0, config.StyleZ.Block1, config.StyleZ.Block2, config.StyleZ.Block3

	// 1100
	//   2233
	return &z{
		tetro: tetro{
			x:      x,
			y:      y,
			block0: Block{x: x, y: y, style: s0},
			block1: Block{x: x - 1, y: y, style: s1},
			block2: Block{x: x, y: y + 1, style: s2},
			block3: Block{x: x + 1, y: y + 1, style: s3},
		},
	}
}

func (z *z) Rotate() {
	switch z.rotation {
	case 0:
		z.block0.x, z.block0.y = z.x, z.y   //   33
		z.block1.x, z.block1.y = z.x, z.y+1 // 0022
		z.block2.x, z.block2.y = z.x+1, z.y // 11
		z.block3.x, z.block3.y = z.x+1, z.y-1
	case 1:
		z.block0.x, z.block0.y = z.x, z.y   // 3322
		z.block1.x, z.block1.y = z.x+1, z.y //   0011
		z.block2.x, z.block2.y = z.x, z.y-1
		z.block3.x, z.block3.y = z.x-1, z.y-1
	case 2:
		z.block0.x, z.block0.y = z.x, z.y   //   11
		z.block1.x, z.block1.y = z.x, z.y-1 // 2200
		z.block2.x, z.block2.y = z.x-1, z.y // 33
		z.block3.x, z.block3.y = z.x-1, z.y+1

	case 3:
		z.block0.x, z.block0.y = z.x, z.y   // 1100
		z.block1.x, z.block1.y = z.x-1, z.y //   2233
		z.block2.x, z.block2.y = z.x, z.y+1
		z.block3.x, z.block3.y = z.x+1, z.y+1
	default:
	}
	z.updateRotation()
}

func (original *z) clone() Tetro {
	clone := &z{
		tetro: clone(original.tetro), // implemented in tetris.go
	}
	return clone
}

func (z *z) CanRotate(boardHeight, boardWidth int, stationaryBlocks []Block) bool {
	return canRotate(z, boardHeight, boardWidth, stationaryBlocks) // implemented in tetris.go
}
