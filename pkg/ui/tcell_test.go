package ui

import (
	"testing"

	"github.com/garyloug/tetris/pkg/tetris"
)

func TestTcell_Init(t *testing.T) {
	tc := &tcell{}

	tests := []struct {
		name          string
		boardH        int
		boardW        int
		expectError   bool
		errorContains string
	}{
		{
			name:        "valid dimensions",
			boardH:      20,
			boardW:      10,
			expectError: false,
		},
		{
			name:          "zero height",
			boardH:        0,
			boardW:        10,
			expectError:   true,
			errorContains: "invalid board dimensions",
		},
		{
			name:          "negative height",
			boardH:        -1,
			boardW:        10,
			expectError:   true,
			errorContains: "invalid board dimensions",
		},
		{
			name:          "zero width",
			boardH:        20,
			boardW:        0,
			expectError:   true,
			errorContains: "invalid board dimensions",
		},
		{
			name:          "negative width",
			boardH:        20,
			boardW:        -1,
			expectError:   true,
			errorContains: "invalid board dimensions",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tc.Init(tt.boardH, tt.boardW)

			if tt.expectError {
				if err == nil {
					t.Errorf("Init(%d, %d) expected error, got nil", tt.boardH, tt.boardW)
				} else if tt.errorContains != "" && err.Error() != tt.errorContains {
					t.Errorf("Init(%d, %d) error = %q, want %q", tt.boardH, tt.boardW, err.Error(), tt.errorContains)
				}
			} else {
				if err != nil {
					t.Errorf("Init(%d, %d) unexpected error: %v", tt.boardH, tt.boardW, err)
				}
				if tc.boardHeight != tt.boardH {
					t.Errorf("Init(%d, %d) boardHeight = %d, want %d", tt.boardH, tt.boardW, tc.boardHeight, tt.boardH)
				}
				if tc.boardWidth != tt.boardW {
					t.Errorf("Init(%d, %d) boardWidth = %d, want %d", tt.boardH, tt.boardW, tc.boardWidth, tt.boardW)
				}
			}
		})
	}
}

func TestTcell_GetBlockStyles(t *testing.T) {
	tc := &tcell{}
	tc.ui = ui{
		oStyles: "O",
		iStyles: "I",
		sStyles: "S",
		zStyles: "Z",
		lStyles: "L",
		jStyles: "J",
		tStyles: "T",
	}

	o, i, s, z, l, j, _ := tc.GetBlockStyles()

	if o != "O" {
		t.Errorf("GetBlockStyles() O style incorrect: got %+v, want %+v", o, "O")
	}
	if i != "I" {
		t.Errorf("GetBlockStyles() I style incorrect: got %+v, want %+v", i, "I")
	}
	if s != "S" {
		t.Errorf("GetBlockStyles() S style incorrect: got %+v, want %+v", s, "S")
	}
	if z != "Z" {
		t.Errorf("GetBlockStyles() Z style incorrect: got %+v, want %+v", z, "Z")
	}
	if l != "L" {
		t.Errorf("GetBlockStyles() L style incorrect: got %+v, want %+v", l, "L")
	}
	if j != "J" {
		t.Errorf("GetBlockStyles() J style incorrect: got %+v, want %+v", j, "J")
	}
}

func TestTcell_SetDevStyles(t *testing.T) {
	tc := &tcell{}
	tc.ui = ui{
		oStyles: TcellStyle{fill: full},
		iStyles: TcellStyle{fill: full},
		sStyles: TcellStyle{fill: full},
		zStyles: TcellStyle{fill: full},
		lStyles: TcellStyle{fill: full},
		jStyles: TcellStyle{fill: full},
		tStyles: TcellStyle{fill: full},
	}

	tc.setDevStyles()

	checkDevStyle := func(style tetris.BlockStyle, name string, expectedChar rune) {
		tcellStyle, ok := style.(TcellStyle)
		if !ok {
			t.Errorf("setDevStyles() %s style is not TcellStyle type", name)
			return
		}
		if tcellStyle.fill != expectedChar {
			t.Errorf("setDevStyles() %s fill = %c, want %c", name, tcellStyle.fill, expectedChar)
		}
	}

	checkDevStyle(tc.oStyles, "O", 'O')
	checkDevStyle(tc.iStyles, "I", 'I')
	checkDevStyle(tc.sStyles, "S", 'S')
	checkDevStyle(tc.zStyles, "Z", 'Z')
	checkDevStyle(tc.lStyles, "L", 'L')
	checkDevStyle(tc.jStyles, "J", 'J')
	checkDevStyle(tc.tStyles, "T", 'T')
}
