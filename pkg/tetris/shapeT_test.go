package tetris

import (
	"testing"
)

func TestNewT(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 5,
		StyleT: "T",
		StyleI: "I",
		StyleO: "O",
		StyleS: "S",
		StyleZ: "Z",
		StyleL: "L",
		StyleJ: "J",
	})

	t_piece := newT()

	if t_piece == nil {
		t.Fatal("newT() returned nil")
	}

	blocks := t_piece.Blocks()
	if len(blocks) != 4 {
		t.Errorf("newT() should have 4 blocks, got %d", len(blocks))
	}

	expectedInitialCoords := []struct{ x, y int }{
		{5, 5}, // block0: spawn position
		{4, 5}, // block1: x - 1
		{6, 5}, // block2: x + 1
		{5, 4}, // block3: y - 1
	}

	for i, expected := range expectedInitialCoords {
		x, y := blocks[i].Coordinates()
		if x != expected.x || y != expected.y {
			t.Errorf("Initial T shape: Block %d at (%d,%d), want (%d,%d)", i, x, y, expected.x, expected.y)
		}
	}

	expectedStyle := config.StyleT
	for i, block := range blocks {
		if block.Style() != expectedStyle {
			t.Errorf("Block %d style = %v, want %v", i, block.Style(), expectedStyle)
		}
	}
}

func TestT_Rotate(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 5,
		StyleT: "T",
		StyleI: "I",
		StyleO: "O",
		StyleS: "S",
		StyleZ: "Z",
		StyleL: "L",
		StyleJ: "J",
	})

	tests := []struct {
		name     string
		rotation int
	}{
		{"rotation 0", 0},
		{"rotation 1", 1},
		{"rotation 2", 2},
		{"rotation 3", 3},
		{"rotation 4", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t_piece := newT()
			for r := 0; r < tt.rotation; r++ {
				t_piece.Rotate()
			}

			blocks := t_piece.Blocks()
			if len(blocks) != 4 {
				t.Errorf("T piece should have 4 blocks after %d rotations, got %d", tt.rotation, len(blocks))
			}

			expectedCoords := [][]struct{ x, y int }{
				{{5, 5}, {4, 5}, {6, 5}, {5, 4}}, // rotation 0
				{{5, 5}, {5, 6}, {5, 4}, {4, 5}}, // rotation 1
				{{5, 5}, {6, 5}, {4, 5}, {5, 6}}, // rotation 2
				{{5, 5}, {5, 4}, {5, 6}, {6, 5}}, // rotation 3
				{{5, 5}, {4, 5}, {6, 5}, {5, 4}}, // rotation 4 (back to original)
			}

			expected := expectedCoords[tt.rotation]
			for i, block := range blocks {
				x, y := block.Coordinates()
				if x != expected[i].x || y != expected[i].y {
					t.Errorf("Rotation %d, Block %d: got (%d,%d), want (%d,%d)", tt.rotation, i, x, y, expected[i].x, expected[i].y)
				}
			}
		})
	}
}

func TestT_CanRotate(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 5,
		StyleT: "T",
		StyleI: "I",
		StyleO: "O",
		StyleS: "S",
		StyleZ: "Z",
		StyleL: "L",
		StyleJ: "J",
	})

	tests := []struct {
		name             string
		boardHeight      int
		boardWidth       int
		stationaryBlocks []Block
		expected         bool
		setupT           func() Tetro
	}{
		{
			name:             "can rotate in empty board",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupT:           func() Tetro { return newT() },
		},
		{
			name:             "can rotate with space",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupT: func() Tetro {
				t_piece := newT()
				t_piece.MoveRight()
				return t_piece
			},
		},
		{
			name:        "can rotate with distant stationary blocks",
			boardHeight: 20,
			boardWidth:  10,
			stationaryBlocks: []Block{
				{x: 0, y: 0, style: "block"},
			},
			expected: true,
			setupT:   func() Tetro { return newT() },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t_piece := tt.setupT()
			result := t_piece.CanRotate(tt.boardHeight, tt.boardWidth, tt.stationaryBlocks)
			if result != tt.expected {
				t.Errorf("CanRotate() = %v, want %v", result, tt.expected)
			}
		})
	}
}
