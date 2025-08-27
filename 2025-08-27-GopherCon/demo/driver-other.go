//go:build !tamago && !tinygo && !noos

package main

import (
	xEmbedded "fyne.io/x/fyne/driver/embedded"
)

func setup() xEmbedded.Driver {
	return nil
}
