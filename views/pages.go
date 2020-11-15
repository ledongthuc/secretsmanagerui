package views

import (
	"github.com/ledongthuc/secretsmanagerui/actions"
	"github.com/rivo/tview"
)

type App struct {
	secretSort actions.SecretSortOption

	app     *tview.Application
	pages   *tview.Pages
	secrets *tview.Table
}

func (a *App) Init() error {
	a.secretSort = actions.SecretSortNameAsc

	a.pages = a.NewPages()
	a.app = tview.NewApplication().
		SetRoot(a.pages, true).
		EnableMouse(true)
	return nil
}

func (a *App) Run() error {
	return a.app.Run()
}

func (a *App) NewPages() *tview.Pages {
	pages := tview.NewPages()
	pages.AddPage("secret", a.NewSecretsScreen(), true, true)
	pages.AddPage("secret-filter", a.NewSecretFilterModal(), true, false)
	return pages
}
