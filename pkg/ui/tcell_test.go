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
		oStyles: tetris.BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
		iStyles: tetris.BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
		sStyles: tetris.BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
		zStyles: tetris.BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
		lStyles: tetris.BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
		jStyles: tetris.BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
		tStyles: tetris.BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
	}

	o, i, s, z, l, j, _ := tc.GetBlockStyles()

	if o.Block0 != "O0" || o.Block1 != "O1" || o.Block2 != "O2" || o.Block3 != "O3" {
		t.Errorf("GetBlockStyles() O styles incorrect: got %+v", o)
	}
	if i.Block0 != "I0" || i.Block1 != "I1" || i.Block2 != "I2" || i.Block3 != "I3" {
		t.Errorf("GetBlockStyles() I styles incorrect: got %+v", i)
	}
	if s.Block0 != "S0" || s.Block1 != "S1" || s.Block2 != "S2" || s.Block3 != "S3" {
		t.Errorf("GetBlockStyles() S styles incorrect: got %+v", s)
	}
	if z.Block0 != "Z0" || z.Block1 != "Z1" || z.Block2 != "Z2" || z.Block3 != "Z3" {
		t.Errorf("GetBlockStyles() Z styles incorrect: got %+v", z)
	}
	if l.Block0 != "L0" || l.Block1 != "L1" || l.Block2 != "L2" || l.Block3 != "L3" {
		t.Errorf("GetBlockStyles() L styles incorrect: got %+v", l)
	}
	if j.Block0 != "J0" || j.Block1 != "J1" || j.Block2 != "J2" || j.Block3 != "J3" {
		t.Errorf("GetBlockStyles() J styles incorrect: got %+v", j)
	}
}

func TestTcell_SetDevStyles(t *testing.T) {
	tc := &tcell{}
	tc.ui = ui{
		oStyles: tetris.BlockStyles{
			Block0: TcellStyle{fill: full},
			Block1: TcellStyle{fill: shade},
			Block2: TcellStyle{fill: full},
			Block3: TcellStyle{fill: shade},
		},
		iStyles: tetris.BlockStyles{
			Block0: TcellStyle{fill: full},
			Block1: TcellStyle{fill: shade},
			Block2: TcellStyle{fill: full},
			Block3: TcellStyle{fill: shade},
		},
		sStyles: tetris.BlockStyles{
			Block0: TcellStyle{fill: full},
			Block1: TcellStyle{fill: shade},
			Block2: TcellStyle{fill: full},
			Block3: TcellStyle{fill: shade},
		},
		zStyles: tetris.BlockStyles{
			Block0: TcellStyle{fill: full},
			Block1: TcellStyle{fill: shade},
			Block2: TcellStyle{fill: full},
			Block3: TcellStyle{fill: shade},
		},
		lStyles: tetris.BlockStyles{
			Block0: TcellStyle{fill: full},
			Block1: TcellStyle{fill: shade},
			Block2: TcellStyle{fill: full},
			Block3: TcellStyle{fill: shade},
		},
		jStyles: tetris.BlockStyles{
			Block0: TcellStyle{fill: full},
			Block1: TcellStyle{fill: shade},
			Block2: TcellStyle{fill: full},
			Block3: TcellStyle{fill: shade},
		},
		tStyles: tetris.BlockStyles{
			Block0: TcellStyle{fill: full},
			Block1: TcellStyle{fill: shade},
			Block2: TcellStyle{fill: full},
			Block3: TcellStyle{fill: shade},
		},
	}

	tc.setDevStyles()

	checkDevStyle := func(styles tetris.BlockStyles, name string) {
		if styles.Block0.(TcellStyle).fill != '0' {
			t.Errorf("setDevStyles() %s Block0 fill = %c, want '0'", name, styles.Block0.(TcellStyle).fill)
		}
		if styles.Block1.(TcellStyle).fill != '1' {
			t.Errorf("setDevStyles() %s Block1 fill = %c, want '1'", name, styles.Block1.(TcellStyle).fill)
		}
		if styles.Block2.(TcellStyle).fill != '2' {
			t.Errorf("setDevStyles() %s Block2 fill = %c, want '2'", name, styles.Block2.(TcellStyle).fill)
		}
		if styles.Block3.(TcellStyle).fill != '3' {
			t.Errorf("setDevStyles() %s Block3 fill = %c, want '3'", name, styles.Block3.(TcellStyle).fill)
		}
	}

	checkDevStyle(tc.oStyles, "O")
	checkDevStyle(tc.iStyles, "I")
	checkDevStyle(tc.sStyles, "S")
	checkDevStyle(tc.zStyles, "Z")
	checkDevStyle(tc.lStyles, "L")
	checkDevStyle(tc.jStyles, "J")
	checkDevStyle(tc.tStyles, "T")
}
