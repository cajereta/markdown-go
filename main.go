package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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

	win.SetContent(container.NewHSplit(edit, preview))

	win.Resize(fyne.Size{Width: 600, Height: 800})
	win.CenterOnScreen()
	win.ShowAndRun()

}

func (app *config) createUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	app.EditWidget = edit
	app.PreviewWidget = preview

	edit.OnChanged= preview.ParseMarkdown

	return edit, preview
}
