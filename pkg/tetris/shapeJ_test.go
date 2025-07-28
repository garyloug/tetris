package tetris

import (
	"testing"
)

func TestNewJ(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 5,
		StyleJ: "J",
		StyleI: "I",
		StyleO: "O",
		StyleS: "S",
		StyleZ: "Z",
		StyleL: "L",
		StyleT: "T",
	})

	j := newJ()

	if j == nil {
		t.Fatal("newJ() returned nil")
	}

	blocks := j.Blocks()
	if len(blocks) != 4 {
		t.Errorf("newJ() should have 4 blocks, got %d", len(blocks))
	}

	expectedInitialCoords := []struct{ x, y int }{
		{5, 5}, {5, 4}, {5, 3}, {4, 5},
	}

	for i, expected := range expectedInitialCoords {
		x, y := blocks[i].Coordinates()
		if x != expected.x || y != expected.y {
			t.Errorf("Initial J shape: Block %d at (%d,%d), want (%d,%d)", i, x, y, expected.x, expected.y)
		}
	}

	expectedStyle := config.StyleJ
	for i, block := range blocks {
		if block.Style() != expectedStyle {
			t.Errorf("Block %d style = %v, want %v", i, block.Style(), expectedStyle)
		}
	}
}

func TestJ_Rotate(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 5,
		StyleJ: "J",
		StyleI: "I",
		StyleO: "O",
		StyleS: "S",
		StyleZ: "Z",
		StyleL: "L",
		StyleT: "T",
	})

	j := newJ()

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
			j = newJ()
			for r := 0; r < tt.rotation; r++ {
				j.Rotate()
			}

			blocks := j.Blocks()
			if len(blocks) != 4 {
				t.Errorf("J piece should have 4 blocks after %d rotations, got %d", tt.rotation, len(blocks))
			}

			expectedCoords := [][]struct{ x, y int }{
				{{5, 5}, {5, 4}, {5, 3}, {4, 5}}, // initial state
				{{5, 5}, {4, 5}, {3, 5}, {5, 6}}, // 1 rotation
				{{5, 5}, {5, 6}, {5, 7}, {6, 5}}, // 2 rotations
				{{5, 5}, {6, 5}, {7, 5}, {5, 4}}, // 3 rotations
				{{5, 5}, {5, 4}, {5, 3}, {4, 5}}, // 4 rotations (back to initial)
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

func TestJ_CanRotate(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 5,
		StyleJ: "J",
		StyleI: "I",
		StyleO: "O",
		StyleS: "S",
		StyleZ: "Z",
		StyleL: "L",
		StyleT: "T",
	})

	tests := []struct {
		name             string
		boardHeight      int
		boardWidth       int
		stationaryBlocks []Block
		expected         bool
		setupJ           func() Tetro
	}{
		{
			name:             "can rotate in empty board",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupJ:           func() Tetro { return newJ() },
		},
		{
			name:             "can rotate with space",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupJ: func() Tetro {
				j := newJ()
				j.MoveRight()
				return j
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
			setupJ:   func() Tetro { return newJ() },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := tt.setupJ()
			result := j.CanRotate(tt.boardHeight, tt.boardWidth, tt.stationaryBlocks)
			if result != tt.expected {
				t.Errorf("CanRotate() = %v, want %v", result, tt.expected)
			}
		})
	}
}
