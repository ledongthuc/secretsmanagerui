package views

import (
	"strings"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/ledongthuc/secretsmanagerui/actions"
	"github.com/rivo/tview"
)

func (a *App) NewSecretFilterModal() tview.Primitive {
	list := tview.NewList()
	for _, sort := range actions.SecretSortPossibleValues {
		list = list.AddItem(string(sort), "", rune(strings.ToLower(string(sort))[0]), nil)
	}

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			a.pages.HidePage("secret-filter")
			return nil
		}
		if event.Key() == tcell.KeyEnter {
			a.pages.HidePage("secret-filter")
			return nil
		}
		return event
	})

	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(list, 13, 1, true).
			AddItem(nil, 0, 1, false), 40, 1, true).
		AddItem(nil, 0, 1, false)
}
