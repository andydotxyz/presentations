package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	input := widget.NewMultiLineEntry()
	output := widget.NewRichTextFromMarkdown("")
	content := container.NewAdaptiveGrid(2, input, output)

	input.OnChanged = output.ParseMarkdown

	w.SetContent(content)
	w.ShowAndRun()
}
