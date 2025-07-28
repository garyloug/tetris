package tetris

import (
	"testing"
)

func TestNewO(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 0,
		StyleO: "O",
		StyleI: "I",
		StyleS: "S",
		StyleZ: "Z",
		StyleL: "L",
		StyleJ: "J",
		StyleT: "T",
	})

	o := newO()

	if o == nil {
		t.Fatal("newO() returned nil")
	}

	blocks := o.Blocks()
	if len(blocks) != 4 {
		t.Errorf("newO() should have 4 blocks, got %d", len(blocks))
	}

	expectedInitialCoords := []struct{ x, y int }{
		{5, 0}, // block0: spawn position
		{5, 1}, // block1: y + 1
		{6, 1}, // block2: x + 1, y + 1
		{6, 0}, // block3: x + 1
	}

	for i, expected := range expectedInitialCoords {
		x, y := blocks[i].Coordinates()
		if x != expected.x || y != expected.y {
			t.Errorf("Initial O shape: Block %d at (%d,%d), want (%d,%d)", i, x, y, expected.x, expected.y)
		}
	}

	expectedStyle := config.StyleO
	for i, block := range blocks {
		if block.Style() != expectedStyle {
			t.Errorf("Block %d style = %v, want %v", i, block.Style(), expectedStyle)
		}
	}
}

func TestO_Rotate(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 0,
		StyleO: "O",
		StyleI: "I",
		StyleS: "S",
		StyleZ: "Z",
		StyleL: "L",
		StyleJ: "J",
		StyleT: "T",
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
			o := newO()
			for r := 0; r < tt.rotation; r++ {
				o.Rotate()
			}

			blocks := o.Blocks()
			if len(blocks) != 4 {
				t.Errorf("O piece should have 4 blocks after %d rotations, got %d", tt.rotation, len(blocks))
			}

			expectedCoords := [][]struct{ x, y int }{
				{{5, 0}, {5, 1}, {6, 1}, {6, 0}}, // initial state
				{{5, 1}, {6, 1}, {6, 0}, {5, 0}}, // 1 rotation
				{{6, 1}, {6, 0}, {5, 0}, {5, 1}}, // 2 rotations
				{{6, 0}, {5, 0}, {5, 1}, {6, 1}}, // 3 rotations
				{{5, 0}, {5, 1}, {6, 1}, {6, 0}}, // 4 rotations (back to original)
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

func TestO_CanRotate(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 5,
		StyleO: "O",
		StyleI: "I",
		StyleS: "S",
		StyleZ: "Z",
		StyleL: "L",
		StyleJ: "J",
		StyleT: "T",
	})

	tests := []struct {
		name             string
		boardHeight      int
		boardWidth       int
		stationaryBlocks []Block
		expected         bool
		setupO           func() Tetro
	}{
		{
			name:             "can rotate in empty board",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupO:           func() Tetro { return newO() },
		},
		{
			name:             "can rotate at edges (O shape fits in 2x2)",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupO: func() Tetro {
				Init(Config{
					SpawnX: 8,
					SpawnY: 18,
					StyleO: "O",
					StyleI: "I",
					StyleS: "S",
					StyleZ: "Z",
					StyleL: "L",
					StyleJ: "J",
					StyleT: "T",
				})
				return newO()
			},
		},
		{
			name:        "can rotate with stationary blocks that don't interfere",
			boardHeight: 20,
			boardWidth:  10,
			stationaryBlocks: []Block{
				{x: 0, y: 0, style: "block"},
			},
			expected: true,
			setupO:   func() Tetro { return newO() },
		},
		{
			name:        "can always rotate - O piece stays in same 2x2 grid",
			boardHeight: 20,
			boardWidth:  10,
			stationaryBlocks: []Block{
				{x: 3, y: 5, style: "block"},
				{x: 7, y: 5, style: "block"},
				{x: 5, y: 3, style: "block"},
				{x: 5, y: 7, style: "block"},
			},
			expected: true,
			setupO:   func() Tetro { return newO() },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := tt.setupO()
			result := o.CanRotate(tt.boardHeight, tt.boardWidth, tt.stationaryBlocks)
			if result != tt.expected {
				t.Errorf("CanRotate() = %v, want %v", result, tt.expected)
			}
		})
	}
}
