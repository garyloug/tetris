package ui

import (
	"testing"
)

func TestMockUI_Implementation(t *testing.T) {
	mockUI, cleanup := newMockUI()
	defer cleanup()

	// check mockUI implements UI interface
	var _ UI = mockUI

	// assert for non-UI field access
	mock := mockUI.(*MockUI)

	err := mock.Init(20, 10)
	if err != nil {
		t.Errorf("mockUI.Init() error = %v, want nil", err)
	}
	if mock.BoardHeight != 20 {
		t.Errorf("mockUI.boardHeight = %v, want %v", mock.BoardHeight, 20)
	}
	if mock.BoardWidth != 10 {
		t.Errorf("mockUI.boardWidth = %v, want %v", mock.BoardWidth, 10)
	}

	o, i, s, z, l, j, tStyle := mock.GetBlockStyles()
	if o.Block0 != "style0" {
		t.Errorf("GetBlockStyles() o.Block0 = %v, want %v", o.Block0, "style0")
	}
	if i.Block1 != "style1" {
		t.Errorf("GetBlockStyles() i.Block1 = %v, want %v", i.Block1, "style1")
	}
	if s.Block2 != "style2" {
		t.Errorf("GetBlockStyles() s.Block2 = %v, want %v", s.Block2, "style2")
	}
	if z.Block3 != "style3" {
		t.Errorf("GetBlockStyles() z.Block3 = %v, want %v", z.Block3, "style3")
	}
	if l.Block0 != "style0" {
		t.Errorf("GetBlockStyles() l.Block0 = %v, want %v", l.Block0, "style0")
	}
	if j.Block1 != "style1" {
		t.Errorf("GetBlockStyles() j.Block1 = %v, want %v", j.Block1, "style1")
	}
	if tStyle.Block2 != "style2" {
		t.Errorf("GetBlockStyles() tStyle.Block2 = %v, want %v", tStyle.Block2, "style2")
	}

	keyChannel := mock.KeyPress()
	if keyChannel == nil {
		t.Error("mockUI.KeyPress() returned nil channel")
	}

	mock.Start()
	if !mock.Started {
		t.Error("mockUI.Start() did not set started flag")
	}

	mock.SendKeyPress(KeyDown)
	select {
	case key := <-keyChannel:
		if key != KeyDown {
			t.Errorf("Received key = %v, want %v", key, KeyDown)
		}
	default:
		t.Error("No key received on channel")
	}

	mock.Stop()
	if !mock.Stopped {
		t.Error("mockUI.Stop() did not set stopped flag")
	}
}
