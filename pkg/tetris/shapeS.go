package tetris

type s struct{ tetro }

func newS() Tetro {
	x, y := config.SpawnX, config.SpawnY
	s0, s1, s2, s3 := config.StyleS.Block0, config.StyleS.Block1, config.StyleS.Block2, config.StyleS.Block3

	//   0011
	// 3322
	return &s{
		tetro: tetro{
			x:      x,
			y:      y,
			block0: Block{x: x, y: y, style: s0},
			block1: Block{x: x + 1, y: y, style: s1},
			block2: Block{x: x, y: y + 1, style: s2},
			block3: Block{x: x - 1, y: y + 1, style: s3},
		},
	}
}

func (s *s) Rotate() {
	switch s.rotation {
	case 0:
		s.block0.x, s.block0.y = s.x, s.y   // 11
		s.block1.x, s.block1.y = s.x, s.y-1 // 0022
		s.block2.x, s.block2.y = s.x+1, s.y //   33
		s.block3.x, s.block3.y = s.x+1, s.y+1
	case 1:
		s.block0.x, s.block0.y = s.x, s.y   //   2233
		s.block1.x, s.block1.y = s.x-1, s.y // 1100
		s.block2.x, s.block2.y = s.x, s.y-1
		s.block3.x, s.block3.y = s.x+1, s.y-1
	case 2:
		s.block0.x, s.block0.y = s.x, s.y   // 33
		s.block1.x, s.block1.y = s.x, s.y+1 // 2200
		s.block2.x, s.block2.y = s.x-1, s.y //   11
		s.block3.x, s.block3.y = s.x-1, s.y-1
	case 3:
		s.block0.x, s.block0.y = s.x, s.y   //   0011
		s.block1.x, s.block1.y = s.x+1, s.y // 3322
		s.block2.x, s.block2.y = s.x, s.y+1
		s.block3.x, s.block3.y = s.x-1, s.y+1
	default:
	}
	s.updateRotation()
}

func (original *s) clone() Tetro {
	clone := &s{
		tetro: clone(original.tetro), // implemented in tetris.go
	}
	return clone
}

func (s *s) CanRotate(boardHeight, boardWidth int, stationaryBlocks []Block) bool {
	return canRotate(s, boardHeight, boardWidth, stationaryBlocks) // implemented in tetris.go
}
