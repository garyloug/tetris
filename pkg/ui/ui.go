package ui

import (
	"fmt"

	"github.com/garyloug/tetris/pkg/tetris"
)

const (
	Running Status = iota
	Pause
	GameOver
	KeyUp KeyPress = iota
	KeyDown
	KeyLeft
	KeyRight
	KeyPause
	KeyStop
	Console UiType = iota
	ConsoleDev
	Mock
)

type (
	KeyPress int
	Status   int
	UiType   int
)

type UI interface {
	Init(boardHeight, boardWidth int) error
	GetBlockStyles() (o, i, s, z, l, j, t tetris.BlockStyle)
	Update(blocks []tetris.Block, queue []tetris.Tetro, score, level, linesCleared int, status Status)

	KeyPress() <-chan KeyPress

	Start()
	Stop()
}

func NewUI(uiType UiType) (UI, func(), error) {
	var ui UI
	var cleanup func()

	switch uiType {
	case Console:
		ui, cleanup = newTcellUI()
	case ConsoleDev:
		ui, cleanup = newTcellDevUI()
	case Mock:
		ui, cleanup = newMockUI()
	default:
		return nil, nil, fmt.Errorf("unsupported UI type: %d", uiType)
	}

	return ui, cleanup, nil
}

type ui struct {
	// board config
	boardHeight int
	boardWidth  int

	// block styles for each tetro type
	oStyles tetris.BlockStyle
	iStyles tetris.BlockStyle
	sStyles tetris.BlockStyle
	zStyles tetris.BlockStyle
	lStyles tetris.BlockStyle
	jStyles tetris.BlockStyle
	tStyles tetris.BlockStyle

	// event channel for user input
	eventChan chan KeyPress
}

func (u *ui) KeyPress() <-chan KeyPress {
	return u.eventChan
}
