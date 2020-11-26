package views

import (
	"fmt"

	"github.com/rivo/tview"
)

const (
	EventStatusChangeText = "Status:ChangeText"
)

func (a *App) NewSecretStatus() *tview.TextView {
	statusBar := tview.NewTextView()
	err := a.On(EventStatusChangeText, func(data string) {
		a.onChangeStatus(statusBar, data)
	})
	if err != nil {
		panic("Can't handle event change status bar: " + err.Error())
	}
	return statusBar
}

func (a *App) onChangeStatus(t *tview.TextView, data interface{}) error {
	t.SetText(fmt.Sprintf("%v", data))
	return nil
}
