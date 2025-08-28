package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	d := setup(a)

	winSize := fyne.NewSize(640, 480)
	if d != nil {
		winSize = d.ScreenSize()
	}

	w := a.NewWindow("GopherCon Embedded App Demo")
	w.Resize(winSize)
	th, _ := theme.FromJSON(string(resourceMatrixJson.Content()))
	a.Settings().SetTheme(th)

	glyph := setContent(w)
	w.Show()
	w.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
		switch e.Name {
		case fyne.KeyUp:
			glyph.Move(glyph.Position().SubtractXY(0, 20))
		case fyne.KeyDown:
			glyph.Move(glyph.Position().AddXY(0, 20))
		case fyne.KeyLeft:
			glyph.Move(glyph.Position().SubtractXY(20, 0))
		case fyne.KeyRight:
			glyph.Move(glyph.Position().AddXY(20, 0))
		}
	})

	if d != nil {
		d.Run()
	} else {
		a.Run()
	}
}

func setContent(w fyne.Window) fyne.CanvasObject {
	welcome := widget.NewLabel("Fyne embedded - Esc to quit")
	welcome.Alignment = fyne.TextAlignCenter
	img := canvas.NewImageFromResource(resourceQrPng)
	img.FillMode = canvas.ImageFillContain
	neo := canvas.NewImageFromResource(resourceNeoPng)

	size := w.Canvas().Size()
	neo.Resize(fyne.NewSquareSize(size.Height * .8))
	neo.Move(fyne.NewPos(size.Width-neo.Size().Width-10, size.Height*.15))

	content := container.NewBorder(welcome, nil, nil, nil, container.NewStack(img,
		container.NewWithoutLayout(neo)))

	w.SetContent(content)
	return neo
}
