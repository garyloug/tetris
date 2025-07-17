package game

import (
	"testing"
	"time"

	"github.com/garyloug/tetris/pkg/tetris"
	"github.com/garyloug/tetris/pkg/ui"
)

func TestNewGame(t *testing.T) {
	mockUI, cleanup, err := ui.NewUI(ui.Mock, 20, 10)
	if err != nil {
		t.Fatalf("failed to create mock UI: %v", err)
	}
	defer cleanup()

	game := NewGame(mockUI)

	if game.ui == nil {
		t.Error("NewGame() ui is nil")
	}
	if game.boardHeight != boardHeight {
		t.Errorf("NewGame() boardHeight = %v, want %v", game.boardHeight, boardHeight)
	}
	if game.boardWidth != boardWidth {
		t.Errorf("NewGame() boardWidth = %v, want %v", game.boardWidth, boardWidth)
	}
	if game.score != 0 {
		t.Errorf("NewGame() score = %v, want %v", game.score, 0)
	}
	if game.level != 0 {
		t.Errorf("NewGame() level = %v, want %v", game.level, 0)
	}
	if game.cleared != 0 {
		t.Errorf("NewGame() cleared = %v, want %v", game.cleared, 0)
	}
	if game.pause != false {
		t.Errorf("NewGame() pause = %v, want %v", game.pause, false)
	}
	if game.over != false {
		t.Errorf("NewGame() over = %v, want %v", game.over, false)
	}
	if game.moveDelay != 500*time.Millisecond {
		t.Errorf("NewGame() moveDelay = %v, want %v", game.moveDelay, moveDelay)
	}
	if len(game.tetroQueue) != tetroQueueSize {
		t.Errorf("NewGame() tetroQueue length = %v, want %v", len(game.tetroQueue), tetroQueueSize)
	}
	if len(game.stationaryBlocks) != 0 {
		t.Error("NewGame() stationaryBlocks should be empty at the start")
	}
	if game.done == nil {
		t.Error("NewGame() done channel is nil")
	}
}

func TestGame_Start(t *testing.T) {
	tetris.Init(tetris.Config{
		SpawnX: tetroSpawnX,
		SpawnY: tetroSpawnY,
		StyleO: tetris.BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
		StyleI: tetris.BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
		StyleS: tetris.BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
		StyleZ: tetris.BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
		StyleL: tetris.BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
		StyleJ: tetris.BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
		StyleT: tetris.BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
	})

	mockUI, cleanup, err := ui.NewUI(ui.Mock, 20, 10)
	if err != nil {
		t.Fatalf("failed to create mock UI: %v", err)
	}
	defer cleanup()

	// assert for non-UI field access
	mock := mockUI.(*ui.MockUI)

	game := NewGame(mockUI)

	done := game.Start()
	if done == nil {
		t.Error("Start() returned nil channel")
	}

	if !mock.Started {
		t.Error("Start() did not start UI")
	}

	mock.SendKeyPress(ui.KeyStop)

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Error("Game did not receive stop signal within timeout")
	}
}
