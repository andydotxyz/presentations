//go:build tamago

package main

import (
	"fyne.io/fyne/v2/app"

	xEmbedded "fyne.io/x/fyne/driver/embedded"
)

func setup() xEmbedded.Driver {
	d := xEmbedded.NewUEFIDriver()
	app.SetDriverDetails(d.Details())
	return d
}
