package game

import (
	"testing"
	"time"

	"github.com/garyloug/tetris/pkg/tetris"
	"github.com/garyloug/tetris/pkg/ui"
)

func TestNewGame(t *testing.T) {
	mockUI, cleanup, err := ui.NewUI(ui.Mock)
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
		StyleO: "O",
		StyleI: "I",
		StyleS: "S",
		StyleZ: "Z",
		StyleL: "L",
		StyleJ: "J",
		StyleT: "T",
	})

	mockUI, cleanup, err := ui.NewUI(ui.Mock)
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
