package views

import (
	"github.com/asaskevich/EventBus"
	"github.com/rivo/tview"

	"github.com/ledongthuc/secretsmanagerui/actions"
)

type App struct {
	secretSort actions.SecretSortOption

	app     *tview.Application
	pages   *tview.Pages
	secrets *tview.Table
	bus     EventBus.Bus
}

func (a *App) Init() error {
	a.bus = EventBus.New()
	a.secretSort = actions.SecretSortNameAsc
	a.app = tview.NewApplication().
		SetRoot(a.NewTemplate(), true).
		EnableMouse(true)
	return nil
}

func (a *App) Run() error {
	return a.app.Run()
}

func (a *App) NewTemplate() tview.Primitive {
	a.pages = a.NewPages()
	status := a.NewSecretStatus()
	a.Emit(EventStatusChangeText, "Type S to sort")
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(a.pages, 0, 1, true).
		AddItem(status, 1, 0, false)
	return flex
}

func (a *App) NewPages() *tview.Pages {
	pages := tview.NewPages()
	pages.AddPage("secret", a.NewSecretsScreen(), true, true)
	pages.AddPage("secret-filter", a.NewSecretFilterModal(), true, false)
	return pages
}

func (a *App) On(eventName string, fn interface{}) error {
	return a.bus.Subscribe(eventName, fn)
}

func (a *App) Emit(eventName string, data interface{}) {
	a.bus.Publish(eventName, data)
}
