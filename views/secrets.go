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
	return a.secrets
}

func (a *App) NewSecretsTable() *tview.Table {
	table := tview.NewTable().
		SetSeparator(tview.Borders.Vertical).
		SetFixed(1, 0).
		SetSelectable(true, false).
		SetSelectedStyle(tcell.StyleDefault.Foreground(tcell.ColorLime))

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune && event.Rune() == 's' {
			a.pages.ShowPage("secret-filter")
		}
		return event
	})

	loadSecretTableHeaders(table, a.secretSort)
	loadSecretTableData(table, a.secretSort)
	return table
}

func loadSecretTableHeaders(table *tview.Table, secretSort actions.SecretSortOption) {
	sortSymbol := " ▲"
	if secretSort.SortType == actions.SortTypeDesc {
		sortSymbol = " ▼"
	}
	name := "Name"
	if secretSort.Sort == actions.SecretSortName {
		name += sortSymbol
	}
	accessedAt := "Asscessed at"
	if secretSort.Sort == actions.SecretSortAccessedAt {
		accessedAt += sortSymbol
	}
	createdAt := "Created at"
	if secretSort.Sort == actions.SecretSortCreatedAt {
		createdAt += sortSymbol
	}
	updatedAt := "Updated at"
	if secretSort.Sort == actions.SecretSortUpdatedAt {
		updatedAt += sortSymbol
	}
	rotated := "Rotated"
	rotatedAt := "Rotated at"
	if secretSort.Sort == actions.SecretSortRotatedAt {
		rotatedAt += sortSymbol
	}

	table.SetCell(0, 0, tview.NewTableCell(name).SetTextColor(tcell.ColorGreen).SetExpansion(2).SetSelectable(false))
	table.SetCell(0, 1, tview.NewTableCell(accessedAt).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignCenter).SetSelectable(false))
	table.SetCell(0, 2, tview.NewTableCell(createdAt).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignCenter).SetSelectable(false))
	table.SetCell(0, 3, tview.NewTableCell(updatedAt).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignCenter).SetSelectable(false))
	table.SetCell(0, 4, tview.NewTableCell(rotated).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignCenter).SetSelectable(false))
	table.SetCell(0, 5, tview.NewTableCell(rotatedAt).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignCenter).SetSelectable(false))
}

func loadSecretTableData(table *tview.Table, secretSort actions.SecretSortOption) {
	secrets, err := actions.GetListSecrets(secretSort)
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
