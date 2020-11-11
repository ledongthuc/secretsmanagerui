package views

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	tcell "github.com/gdamore/tcell/v2"
	"github.com/ledongthuc/secretsmanagerui/actions"
	"github.com/rivo/tview"
)

func (a *App) NewSecretsScreen() tview.Primitive {
	a.secrets = a.NewSecretsTable()
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(a.secrets, 0, 1, true).
		AddItem(a.NewSecretStatus(), 1, 0, false)
	return flex
}

func (a *App) NewSecretStatus() *tview.TextView {
	return tview.NewTextView().SetText("s: change sorting")
}

func (a *App) NewSecretsTable() *tview.Table {
	table := tview.NewTable().
		SetSeparator(tview.Borders.Vertical).
		// SetBorders(true).
		SetFixed(1, 0).
		SetSelectable(true, false).
		SetSelectedStyle(tcell.StyleDefault.Foreground(tcell.ColorLime))

	// table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	// 	if event.Key() == tcell.KeyRune && event.Rune() == 's' {
	// 		a.pages.ShowPage("secret-filter")
	// 	}
	// 	return event
	// })

	loadSecretTableHeaders(table)
	loadSecretTableData(table)
	return table
}

func loadSecretTableHeaders(table *tview.Table) {
	table.SetCell(0, 0, tview.NewTableCell("Name").SetTextColor(tcell.ColorGreen).SetExpansion(2).SetSelectable(false))
	table.SetCell(0, 1, tview.NewTableCell("Asscessed at").SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignCenter).SetSelectable(false))
	table.SetCell(0, 2, tview.NewTableCell("Created at").SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignCenter).SetSelectable(false))
	table.SetCell(0, 3, tview.NewTableCell("Updated at").SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignCenter).SetSelectable(false))
	table.SetCell(0, 4, tview.NewTableCell("Rotated").SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignCenter).SetSelectable(false))
	table.SetCell(0, 5, tview.NewTableCell("Rotated at").SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignCenter).SetSelectable(false))
}

func loadSecretTableData(table *tview.Table) {
	secrets, err := actions.GetListSecrets()
	if err != nil {
		panic(err)
	}

	for index, secret := range secrets {
		rotationEnabled := ""
		if aws.BoolValue(secret.RotationEnabled) {
			rotationEnabled = "✓"
		}
		table.SetCell(index+1, 0, tview.NewTableCell(aws.StringValue(secret.Name)))
		table.SetCell(index+1, 1, tview.NewTableCell(" "+aws.TimeValue(secret.LastAccessedDate).Format(time.RFC3339)+" "))
		table.SetCell(index+1, 2, tview.NewTableCell(" "+aws.TimeValue(secret.CreatedDate).Format(time.RFC3339)+" "))
		table.SetCell(index+1, 3, tview.NewTableCell(" "+aws.TimeValue(secret.LastChangedDate).Format(time.RFC3339)+" "))
		table.SetCell(index+1, 4, tview.NewTableCell(rotationEnabled).SetAlign(tview.AlignCenter))
		table.SetCell(index+1, 5, tview.NewTableCell(" "+aws.TimeValue(secret.LastRotatedDate).Format(time.RFC3339)+" "))
	}
}
