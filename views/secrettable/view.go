package secrettable

import (
	"github.com/andlabs/ui"
	"github.com/ledongthuc/secretsmanagerui/actions"
)

func New() *ui.Table {
	secrets, err := actions.GetListSecrets()
	if err != nil {
		panic(err)
	}

	model := ui.NewTableModel(newSecretModelHandler(secrets))
	table := ui.NewTable(&ui.TableParams{Model: model})
	table.AppendTextColumn("Name", 0, ui.TableModelColumnNeverEditable, nil)
	return table
}
