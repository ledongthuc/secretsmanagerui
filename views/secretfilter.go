package views

import (
	"strconv"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/ledongthuc/secretsmanagerui/actions"
	"github.com/rivo/tview"
)

func (a *App) NewSecretFilterModal() tview.Primitive {
	list := tview.NewList()
	for index, sort := range actions.SecretSortPossibleValues {
		sort := sort
		list = list.AddItem(sort.NiceText, sort.Description, rune(strconv.Itoa(index)[0]), func() {
			a.secretSort = sort
			a.secrets.Clear()
			loadSecretTableHeaders(a.secrets, sort)
			loadSecretTableData(a.secrets, sort)
			a.pages.HidePage("secret-filter")
		})
	}

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			a.pages.HidePage("secret-filter")
			return nil
		}
		return event
	})

	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(list, 20, 1, true).
			AddItem(nil, 0, 1, false), 40, 1, true).
		AddItem(nil, 0, 1, false)
}
