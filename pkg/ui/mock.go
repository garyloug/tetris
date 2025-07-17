package ui

import "github.com/garyloug/tetris/pkg/tetris"

type MockUI struct {
	BoardHeight  int
	BoardWidth   int
	keyPressChan chan KeyPress
	blockStyles  tetris.BlockStyles
	Started      bool
	Stopped      bool
}

func newMockUI() (UI, func()) {
	return &MockUI{
		keyPressChan: make(chan KeyPress, 10),
		blockStyles: tetris.BlockStyles{
			Block0: "style0",
			Block1: "style1",
			Block2: "style2",
			Block3: "style3",
		},
	}, func() {}
}

func (m *MockUI) Init(boardHeight, boardWidth int) error {
	m.BoardHeight = boardHeight
	m.BoardWidth = boardWidth
	return nil
}

func (m *MockUI) GetBlockStyles() (o, i, s, z, l, j, t tetris.BlockStyles) {
	return m.blockStyles, m.blockStyles, m.blockStyles, m.blockStyles, m.blockStyles, m.blockStyles, m.blockStyles
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
