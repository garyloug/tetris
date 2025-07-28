package tetris

import (
	"testing"
)

func TestNewL(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 5,
		StyleL: "L",
		StyleI: "I",
		StyleO: "O",
		StyleS: "S",
		StyleZ: "Z",
		StyleJ: "J",
		StyleT: "T",
	})

	l := newL()

	if l == nil {
		t.Fatal("newL() returned nil")
	}

	blocks := l.Blocks()
	if len(blocks) != 4 {
		t.Errorf("newL() should have 4 blocks, got %d", len(blocks))
	}

	expectedInitialCoords := []struct{ x, y int }{
		{5, 5}, // block0: spawn position
		{5, 4}, // block1: y - 1
		{5, 3}, // block2: y - 2
		{6, 5}, // block3: x + 1
	}

	for i, expected := range expectedInitialCoords {
		x, y := blocks[i].Coordinates()
		if x != expected.x || y != expected.y {
			t.Errorf("Initial L shape: Block %d at (%d,%d), want (%d,%d)", i, x, y, expected.x, expected.y)
		}
	}

	expectedStyle := config.StyleL
	for i, block := range blocks {
		if block.Style() != expectedStyle {
			t.Errorf("Block %d style = %v, want %v", i, block.Style(), expectedStyle)
		}
	}
}

func TestL_Rotate(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 5,
		StyleL: "L",
		StyleI: "I",
		StyleO: "O",
		StyleS: "S",
		StyleZ: "Z",
		StyleJ: "J",
		StyleT: "T",
	})

	l := newL()

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
			l = newL()
			for r := 0; r < tt.rotation; r++ {
				l.Rotate()
			}

			blocks := l.Blocks()
			if len(blocks) != 4 {
				t.Errorf("L piece should have 4 blocks after %d rotations, got %d", tt.rotation, len(blocks))
			}

			expectedCoords := [][]struct{ x, y int }{
				{{5, 5}, {5, 4}, {5, 3}, {6, 5}}, // rotation 0
				{{5, 5}, {4, 5}, {3, 5}, {5, 4}}, // rotation 1
				{{5, 5}, {5, 6}, {5, 7}, {4, 5}}, // rotation 2
				{{5, 5}, {6, 5}, {7, 5}, {5, 6}}, // rotation 3
				{{5, 5}, {5, 4}, {5, 3}, {6, 5}}, // rotation 4 (back to original)
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

func TestL_CanRotate(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 5,
		StyleL: "L",
		StyleI: "I",
		StyleO: "O",
		StyleS: "S",
		StyleZ: "Z",
		StyleJ: "J",
		StyleT: "T",
	})

	tests := []struct {
		name             string
		boardHeight      int
		boardWidth       int
		stationaryBlocks []Block
		expected         bool
		setupL           func() Tetro
	}{
		{
			name:             "can rotate in empty board",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupL:           func() Tetro { return newL() },
		},
		{
			name:             "can rotate with space",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupL: func() Tetro {
				l := newL()
				l.MoveRight()
				return l
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
			setupL:   func() Tetro { return newL() },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.setupL()
			result := l.CanRotate(tt.boardHeight, tt.boardWidth, tt.stationaryBlocks)
			if result != tt.expected {
				t.Errorf("CanRotate() = %v, want %v", result, tt.expected)
			}
		})
	}
}
