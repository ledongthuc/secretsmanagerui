package secrettable

import (
	"fmt"

	"github.com/andlabs/ui"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type secretModelHandler struct {
	data []*secretsmanager.SecretListEntry
}

func newSecretModelHandler(data []*secretsmanager.SecretListEntry) *secretModelHandler {
	h := new(secretModelHandler)
	h.data = data
	return h
}

func (mh *secretModelHandler) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	return []ui.TableValue{
		ui.TableString(""),
	}
}

func (mh *secretModelHandler) NumRows(m *ui.TableModel) int {
	return len(mh.data)
}

func (mh *secretModelHandler) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	secret := mh.data[row]
	if secret == nil {
		return ui.TableString(fmt.Sprintf(""))
	}
	return ui.TableString(aws.StringValue(secret.Name))
}

func (mh *secretModelHandler) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {
}
