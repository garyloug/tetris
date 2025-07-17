package tetris

type Block struct {
	x, y  int
	style any
}

func (b Block) Coordinates() (x, y int) {
	return b.x, b.y
}

func (b Block) Style() any {
	return b.style
}

func (b *Block) MoveDown() {
	b.y++
}

func (b *Block) moveRight() {
	b.x++
}

func (b *Block) moveLeft() {
	b.x--
}

// nolint:unused
func (b *Block) moveUp() {
	b.y--
}
