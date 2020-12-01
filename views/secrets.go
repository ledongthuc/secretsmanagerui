package views

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	tcell "github.com/gdamore/tcell/v2"
	"github.com/ledongthuc/secretsmanagerui/actions"
	"github.com/rivo/tview"
)

const (
	EventSecretsLoaded = "Secrets:Loaded"
)

func (a *App) NewSecretsScreen() tview.Primitive {
	a.secrets = a.NewSecretsTable()
	return a.secrets
}

func (a *App) ShowSecretsScreen() {
	a.pages.ShowPage("secret")
	a.Emit(EventStatusChangeText, "Type S to sort")
	if secrets, err := actions.GetListSecrets(a.secretSort); err != nil {
		panic(err)
	} else {
		a.Emit(EventSecretsLoaded, secrets)
	}
}

func (a *App) NewSecretsTable() *tview.Table {
	table := tview.NewTable().
		SetSeparator(tview.Borders.Vertical).
		SetFixed(1, 0).
		SetSelectable(true, false).
		SetSelectedStyle(tcell.StyleDefault.Foreground(tcell.ColorLime))
	loadSecretTableHeaders(table, a.secretSort)

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune && event.Rune() == 's' {
			a.ShowSecretsFilterModal()
		}
		return event
	})

	err := a.On(EventSecretsLoaded, func(ms []*secretsmanager.SecretListEntry) {
		a.onLoadedSecrets(table, ms)
	})
	if err != nil {
		panic("Can't handle event load data table: " + err.Error())
	}
	return table
}

func (a *App) onLoadedSecrets(table *tview.Table, ms []*secretsmanager.SecretListEntry) error {
	for index, secret := range ms {
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
	return nil
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
