package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Markdown Editor")

	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	edit.OnChanged = preview.ParseMarkdown

	w.SetContent(
		container.NewAdaptiveGrid(2, edit, preview))
	w.Resize(fyne.NewSize(280, 220))
	w.ShowAndRun()
}
