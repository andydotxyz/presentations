package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Markdown")

	editor := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	editor.OnChanged = preview.ParseMarkdown

	w.SetContent(container.NewAdaptiveGrid(2, editor, preview))
	w.Resize(fyne.NewSize(280, 200))
	w.ShowAndRun()
}
