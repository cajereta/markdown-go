package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var cfg config

func main() {
	application := app.New()

	win := application.NewWindow("Markdown")

	edit, preview := cfg.createUI()
	cfg.createMenuItems(win)

	win.SetContent(container.NewHSplit(edit, preview))

	win.Resize(fyne.Size{Width: 650, Height: 800})
	win.CenterOnScreen()
	win.ShowAndRun()
}

func (app *config) createUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	app.EditWidget = edit
	app.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

func (app *config) createMenuItems(win fyne.Window) {
	openMenuOption := fyne.NewMenuItem("Open...", func() {})
	saveMenuOption := fyne.NewMenuItem("Save", func() {})
	app.SaveMenuItem = saveMenuOption
	app.SaveMenuItem.Disabled = true
	saveAsMenuOption := fyne.NewMenuItem("Save as...", app.saveAsFunc(win))

	fileMenu := fyne.NewMenu("File", openMenuOption, saveMenuOption, saveAsMenuOption)
	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}

func (app *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if write == nil {
				return
			}

			write.Write([]byte(app.EditWidget.Text))
			app.CurrentFile = write.URI()
			defer write.Close()

			win.SetTitle(win.Title() + " - " + write.URI().Name())

			app.SaveMenuItem.Disabled = false
		}, win)
		saveDialog.Show()
	}
}
