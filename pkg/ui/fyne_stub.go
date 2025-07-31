//go:build console

package ui

import "fmt"

func newFyneUI() (UI, func(), error) {
	return nil, func() {}, fmt.Errorf("Fyne UI is not available in this build")
}
