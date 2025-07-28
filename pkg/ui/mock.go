package ui

import "github.com/garyloug/tetris/pkg/tetris"

type MockUI struct {
	BoardHeight  int
	BoardWidth   int
	keyPressChan chan KeyPress
	blockStyle   tetris.BlockStyle
	Started      bool
	Stopped      bool
}

func newMockUI() (UI, func()) {
	return &MockUI{
		keyPressChan: make(chan KeyPress, 10),
		blockStyle:   "mockStyle",
	}, func() {}
}

func (m *MockUI) Init(boardHeight, boardWidth int) error {
	m.BoardHeight = boardHeight
	m.BoardWidth = boardWidth
	return nil
}

func (m *MockUI) GetBlockStyles() (o, i, s, z, l, j, t tetris.BlockStyle) {
	return m.blockStyle, m.blockStyle, m.blockStyle, m.blockStyle, m.blockStyle, m.blockStyle, m.blockStyle
}

func (m *MockUI) Update(blocks []tetris.Block, queue []tetris.Tetro, score, level, linesCleared int, status Status) {
}

func (m *MockUI) KeyPress() <-chan KeyPress {
	return m.keyPressChan
}

func (m *MockUI) Start() {
	m.Started = true
}

func (m *MockUI) Stop() {
	if !m.Stopped {
		m.Stopped = true
		close(m.keyPressChan)
	}
}

func (m *MockUI) SendKeyPress(key KeyPress) {
	if !m.Stopped {
		select {
		case m.keyPressChan <- key:
		default:
		}
	}
}
